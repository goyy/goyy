// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
)

type htmlServeMux struct {
	isCompiled bool
	templates  map[string]*templateInfo
}

type templateInfo struct {
	id              string
	content         string
	includes        []string
	lastModified    int64
	maxLastModified int64 // include file:lastModified
}

type directiveInfo struct {
	statement  string // <!--#include file="/footer.html" param="home"-->content<!--#endinclude"-->
	directive  string // include
	argKey     string // file
	argValue   string // /footer.html
	paramKey   string // param
	paramValue string // home
	content    string // content
	begin      int    // postion:<!--#include file="
	center     int    // postion:"-->
	end        int    // postion:<!--#endinclude"-->
}

/*
type directiveInfo struct {
	statement string // {{if param `home`}}cur{{end}}
	directive string // if
	argKey    string // param
	argValue  string // home
	content   string // cur
	begin     int    // postion:{{if
	center    int    // postion:}}
	end       int    // postion:{{end}}
}
*/

type tagInfo struct {
	statement string // <link rel="shortcut icon" href="/favicon.ico" go:href="{{assets}}/favicon.ico">
	newstmt   string // <link rel="shortcut icon" href="/static/favicon.ico">
	attr      string // href
	tagBegin  int    // postion:<
	tagEnd    int    // postion:>
	srcBegin  int    // postion: href="
	srcEnd    int    // postion:"
	dstBegin  int    // postion: go:href="
	dstEnd    int    // postion:"
	dstVal    string // {{assets}}/favicon.ico
}

type tagTextInfo struct {
	statement string // <title go:title="/title.html">login</title>
	newstmt   string // <title>login-appendTitle</title>
	attr      string // title
	tagBegin  int    // postion:<
	tagEnd    int    // postion:</
	srcBegin  int    // postion:>
	srcEnd    int    // postion:</
	srcVal    string // login
	dstBegin  int    // postion: go:title="
	dstEnd    int    // postion:"
	dstVal    string // /title.html
}

var hsm = &htmlServeMux{
	templates: make(map[string]*templateInfo),
}

var htmlMutex sync.Mutex

func (me *htmlServeMux) compile(options *htmlOptions) error {
	filepath.Walk(options.Dir, func(path string, info os.FileInfo, err error) error {
		r, err := filepath.Rel(options.Dir, path)
		if err != nil {
			logger.Error(err.Error())
			return err
		} else {
			r = strings.Replace(r, "\\", "/", -1)
		}

		ext := strings.ToLower(files.Extension(r))

		for _, extension := range options.Extensions {
			if ext == extension {
				if c, err := files.Read(path); err == nil {
					lm := times.NowUnix()
					if modTime, err := files.ModTimeUnix(path); err == nil {
						lm = modTime
					} else {
						logger.Error(err.Error())
					}
					id := "/" + filepath.ToSlash(r)
					c, i, m := me.parseContent(c)
					mls := m
					if lm > mls {
						mls = lm
					}
					ti := &templateInfo{
						id:              id,
						content:         c,
						includes:        i,
						lastModified:    lm,
						maxLastModified: mls,
					}
					if _, ok := me.templates[id]; !ok {
						me.templates[id] = ti
					}
				} else {
					logger.Error(err.Error())
					return err
				}
				break
			}
		}
		return nil
	})
	return nil
}

func (me *htmlServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	if me.isHtml(r.URL.Path) {
		filename := Conf.Html.Dir + r.URL.Path
		if files.IsExist(filename) {
			if me.isCompiled == false {
				htmlMutex.Lock()
				if me.isCompiled == false {
					if err := me.compile(Conf.Html); err != nil {
						htmlMutex.Unlock()
						logger.Error(err.Error())
						return true
					}
					me.isCompiled = true
				}
				htmlMutex.Unlock()
			}
			if Conf.Html.Reloaded {
				if me.isUseBrowserCache(w, r, filename) {
					return true
				}
				if c, err := files.Read(filename); err != nil {
					w.Write([]byte(err.Error()))
				} else {
					w.Write([]byte(me.parseFile(w, r, c)))
				}
			} else {
				if ti, ok := me.templates[r.URL.Path]; ok {
					if me.hasUseBrowserCache(w, r, ti.maxLastModified) {
						return true
					}
					w.Write([]byte(ti.content))
				}
			}
			return true
		}
	}
	return false
}

func (me *htmlServeMux) replaceAssets(content string) string {
	content = strings.Replace(content, tagProfile, Conf.Profile, -1)
	content = strings.Replace(content, tagApis, Conf.Api.URL, -1)
	content = strings.Replace(content, tagAssets, Conf.Asset.URL, -1)
	content = strings.Replace(content, tagAssetsStatics, Conf.Static.URL, -1)
	content = strings.Replace(content, tagAssetsDevelopers, Conf.Developer.URL, -1)
	content = strings.Replace(content, tagAssetsOperations, Conf.Operation.URL, -1)
	return strings.Replace(content, tagAssetsUploads, Conf.Upload.URL, -1)
}

func (me *htmlServeMux) isHtml(path string) bool {
	for _, extension := range Conf.Html.Extensions {
		if strings.HasSuffix(path, "."+extension) {
			return true
		}
	}
	return false
}

func (me *htmlServeMux) isInclude(content string) bool {
	if strings.Index(content, directiveIncludeEnd) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) hasUseBrowserCache(w http.ResponseWriter, r *http.Request, fileModTime int64) bool {
	// Browser save file last modified time
	browserModTime := r.Header.Get(ifModifiedSince)
	if strings.IsNotBlank(browserModTime) {
		if v, err := times.ParseUnixGMT(browserModTime); err == nil {
			if fileModTime > v {
				maxLastModified := times.FormatUnixGMT(fileModTime)
				w.Header().Set(lastModified, maxLastModified)
			} else {
				// Tell the browser to use the cache
				w.WriteHeader(304)
				return true
			}
		} else {
			logger.Error(err.Error())
		}
	} else {
		maxLastModified := times.FormatUnixGMT(fileModTime)
		w.Header().Set(lastModified, maxLastModified)
	}
	return false
}

func (me *htmlServeMux) isUseBrowserCache(w http.ResponseWriter, r *http.Request, filename string) bool {
	if fileModTimeUnix, err := files.ModTimeUnix(filename); err == nil {
		var browserModTimeUnix int64
		// Browser save file last modified time
		browserModTime := r.Header.Get(ifModifiedSince)
		if strings.IsNotBlank(browserModTime) {
			if v, err := times.ParseUnixGMT(browserModTime); err == nil {
				browserModTimeUnix = v
			} else {
				logger.Error(err.Error())
			}
		}
		if browserModTimeUnix < fileModTimeUnix {
			// Actual file last modified time
			fileModTime := times.FormatUnixGMT(fileModTimeUnix)
			// Tell the browser not to use cache
			w.Header().Set(lastModified, fileModTime)
			return false
		} else {
			var content string
			if c, err := files.Read(filename); err == nil {
				content = c
			}
			if me.isInclude(content) {
				var includeFileModTimeUnix int64
				directives := make([]directiveInfo, 0)
				directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveArgEnd, directiveIncludeEnd, directives)
				for _, v := range directives {
					if val, err := files.ModTimeUnix(v.argValue); err == nil {
						if includeFileModTimeUnix < val {
							includeFileModTimeUnix = val
						}
					}
				}
				if browserModTimeUnix < includeFileModTimeUnix {
					// The actual last modification time of the include file
					includeFileModTime := times.FormatUnixGMT(includeFileModTimeUnix)
					// Tell the browser not to use cache
					w.Header().Set(lastModified, includeFileModTime)
					return false
				}
			}
			// Tell the browser to use the cache
			w.WriteHeader(304)
			return true
		}
	}
	return false
}

func (me *htmlServeMux) parseContent(content string) (string, []string, int64) {
	content, i, m := me.parseIncludeFile(content)
	//content = me.replaceAssets(content)
	content = me.parseIfFile(content)
	content = me.parseProfileFile(content)
	content = me.parseTagAttrFile(content, tagAttrClass)
	content = me.parseTagAttrFile(content, tagAttrHref)
	content = me.parseTagAttrFile(content, tagAttrSrc)
	content = me.parseTagAttrFile(content, tagAttrAction)
	content = me.parseTagAttrFile(content, tagAttrOnerror)
	content = me.parseTagAttrFile(content, tagAttrOnclick)
	content = me.parseTagTextFile(content, tagTextTitle)
	content = me.parseTagTextFile(content, tagTextType)
	return content, i, m
}

func (me *htmlServeMux) parseFile(w http.ResponseWriter, r *http.Request, content string) string {
	content, _, _ = me.parseContent(content)
	return content
}

func (me *htmlServeMux) parseIncludeFile(content string) (string, []string, int64) {
	includes := []string{}
	lastModified := int64(0)
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveArgEnd, directiveIncludeEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		filename := directives[i].argValue
		v, err := files.Read(filename)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if modTime, err := files.ModTimeUnix(filename); err == nil {
			if modTime > lastModified {
				lastModified = modTime
			}
		} else {
			logger.Error(err.Error())
		}
		f := strings.After(filename, Conf.Html.Dir)
		includes = append(includes, f)

		ifparams := make([]directiveInfo, 0)
		ifparams = me.buildDirectiveInfo(v, tplBegin, tplArgEnd, tplEnd, ifparams)
		for p := len(ifparams) - 1; p >= 0; p-- {
			if directives[i].paramValue == ifparams[p].argValue {
				v = strings.Replace(v, ifparams[p].statement, ifparams[p].content, -1)
			} else {
				v = strings.Replace(v, ifparams[p].statement, "", -1)
			}
		}

		content = strings.Replace(content, directives[i].statement, v, -1)
	}
	return content, includes, lastModified
}

func (me *htmlServeMux) parseIfFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveIfBegin, directiveArgEnd, directiveIfEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if directives[i].argValue == "false" {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseProfileFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveProfileBegin, directiveArgEnd, directiveProfileEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if !profile.Accepts(directives[i].argValue) {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseTagAttrFile(content, attr string) string {
	tags := make([]tagInfo, 0)
	tags = me.buildTagAttrInfo(content, attr, tags)
	for i := len(tags) - 1; i >= 0; i-- {
		content = strings.Replace(content, tags[i].statement, tags[i].newstmt, -1)
	}
	return content
}

func (me *htmlServeMux) parseTagTextFile(content, attr string) string {
	tags := make([]tagTextInfo, 0)
	tags = me.buildTagTextInfo(content, attr, tags)
	for i := len(tags) - 1; i >= 0; i-- {
		content = strings.Replace(content, tags[i].statement, tags[i].newstmt, -1)
	}
	return content
}

func (me *htmlServeMux) buildDirectiveInfo(content, directiveBegin, argEnd, directiveEnd string, directives []directiveInfo) []directiveInfo {
	pos := 0
	for {
		begin := strings.IndexStart(content, directiveBegin, pos)
		if begin == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		center := strings.IndexStart(content, argEnd, begin)
		if center == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		end := strings.IndexStart(content, directiveEnd, center)
		if end == -1 {
			if pos == 0 {
				return directives
			}
			break
		}
		contents := strings.Slice(content, center+len(argEnd), end)
		pos = end
		paramKey := "param"
		paramValue := ""
		argValue := strings.Slice(content, begin+len(directiveBegin), center)
		if directiveBegin == directiveIncludeBegin {
			if strings.Contains(argValue, directiveIncludeParamBegin) {
				paramValue = strings.After(argValue, directiveIncludeParamBegin)
				argValue = strings.Before(argValue, directiveIncludeParamBegin)
			}
			argValue = Conf.Html.Dir + argValue
			if !files.IsExist(argValue) {
				continue
			}
		}
		statement := strings.Slice(content, begin, end+len(directiveEnd))
		directive := "include"
		argKey := "file"
		switch directiveBegin {
		case directiveIfBegin:
			directive = "if"
			argKey = "expr"
		}
		ii := directiveInfo{
			statement:  statement,
			directive:  directive,
			argKey:     argKey,
			argValue:   argValue,
			content:    contents,
			paramKey:   paramKey,
			paramValue: paramValue,
			begin:      begin,
			center:     center,
			end:        end,
		}
		directives = append(directives, ii)
	}
	return directives
}

func (me *htmlServeMux) buildTagAttrInfo(content, attr string, tags []tagInfo) []tagInfo {
	pos := 0
	srcBeginPre := " " + attr + tagAttrPost
	dstBeginPre := tagAttrPre + attr + tagAttrPost
	for {
		dstBegin := strings.IndexStart(content, dstBeginPre, pos)
		if dstBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		dstEnd := strings.IndexStart(content, tagAttrEnd, dstBegin+len(dstBeginPre))
		if dstEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagBegin := strings.IndexForward(content, tagBeginPre, dstBegin)
		if tagBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagEnd := strings.IndexStart(content, tagEndPre, dstEnd)
		if tagEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		pos = tagEnd
		statement := strings.Slice(content, tagBegin, tagEnd+len(tagEndPre))
		newstmt := statement
		dstVal := strings.Slice(content, dstBegin+len(dstBeginPre), dstEnd)
		srcBegin := strings.Index(newstmt, srcBeginPre)
		var srcEnd int
		if srcBegin > 0 {
			srcEnd = strings.IndexStart(newstmt, tagAttrEnd, srcBegin+len(srcBeginPre))
			newstmt = strings.Overlay(newstmt, "", srcBegin, srcEnd+len(tagAttrEnd))
		}
		newstmt = strings.Replace(newstmt, dstBeginPre, srcBeginPre, -1)
		newstmt = me.replaceAssets(newstmt)
		ti := tagInfo{
			statement: statement,
			newstmt:   newstmt,
			attr:      attr,
			tagBegin:  tagBegin,
			tagEnd:    tagEnd,
			srcBegin:  srcBegin,
			srcEnd:    srcEnd,
			dstBegin:  dstBegin,
			dstEnd:    dstEnd,
			dstVal:    dstVal,
		}
		tags = append(tags, ti)
	}
	return tags
}

func (me *htmlServeMux) buildTagTextInfo(content, attr string, tags []tagTextInfo) []tagTextInfo {
	pos := 0
	dstBeginPre := tagAttrPre + attr + tagAttrPost
	for {
		dstBegin := strings.IndexStart(content, dstBeginPre, pos)
		if dstBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		dstEnd := strings.IndexStart(content, tagAttrEnd, dstBegin+len(dstBeginPre))
		if dstEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagBegin := strings.IndexForward(content, tagBeginPre, dstBegin)
		if tagBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		tagEnd := strings.IndexStart(content, tagTextEndPre, dstEnd)
		if tagEnd == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		srcBegin := strings.IndexStart(content, tagEndPre, dstEnd)
		if srcBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		srcEnd := tagEnd
		srcVal := strings.Slice(content, srcBegin+len(tagEndPre), srcEnd)
		pos = tagEnd
		statement := strings.Slice(content, tagBegin, tagEnd+len(tagTextEndPre))
		newstmt := statement
		dstVal := strings.Slice(content, dstBegin+len(dstBeginPre), dstEnd)
		if strings.IsNotBlank(dstVal) {
			switch attr {
			case tagTextTitle:
				filename := Conf.Html.Dir + dstVal
				if files.IsExist(filename) {
					if c, err := files.Read(filename); err == nil {
						title := strings.Slice(content, srcBegin+len(tagBeginPre), srcEnd)
						newstmt = strings.Replace(newstmt, title, title+c, 1)
						dst := dstBeginPre + dstVal + tagAttrEnd
						newstmt = strings.Replace(newstmt, dst, "", 1)
					}
				}
			case tagTextType:
				dst := tagAttrPre + tagTextType
				newstmt = strings.Replace(newstmt, dst, " "+tagTextType, 1)
			}
		}

		newstmt = me.replaceAssets(newstmt)
		tti := tagTextInfo{
			statement: statement,
			newstmt:   newstmt,
			attr:      attr,
			tagBegin:  tagBegin,
			tagEnd:    tagEnd,
			srcBegin:  srcBegin,
			srcEnd:    srcEnd,
			srcVal:    srcVal,
			dstBegin:  dstBegin,
			dstEnd:    dstEnd,
			dstVal:    dstVal,
		}
		tags = append(tags, tti)
	}
	return tags
}
