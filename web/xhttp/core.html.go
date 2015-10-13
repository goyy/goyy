// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
	"net/http"
	"strconv"
)

type htmlServeMux struct{}

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

var hsm = &htmlServeMux{}

func (me *htmlServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	if me.isHtml(r.URL.Path) {
		filename := Conf.Template.Dir + r.URL.Path
		if files.IsExist(filename) {
			if me.isUseBrowserCache(w, r, filename) {
				return true
			}
			if c, err := files.Read(filename); err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(me.parseFile(w, r, c)))
			}
			return true
		}
	}
	return false
}

func (me *htmlServeMux) replaceAssets(content string) string {
	content = strings.Replace(content, tagStaticApis, Conf.Api.URL, -1)
	content = strings.Replace(content, tagStaticAssets, Conf.Asset.URL, -1)
	content = strings.Replace(content, tagStaticDevelopers, Conf.Developer.URL, -1)
	content = strings.Replace(content, tagStaticOperations, Conf.Operation.URL, -1)
	return strings.Replace(content, tagStaticUploads, Conf.Upload.URL, -1)
}

func (me *htmlServeMux) isHtml(path string) bool {
	if strings.HasSuffix(path, ".html") {
		return true
	}
	return false
}

func (me *htmlServeMux) isInclude(content string) bool {
	if strings.Index(content, directiveIncludeEnd) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isSec(content string) bool {
	if strings.Index(content, directiveSecEnd) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isUseBrowserCache(w http.ResponseWriter, r *http.Request, filename string) bool {
	if fileModTimeUnix, err := files.ModTime(filename); err == nil {
		var browserModTimeUnix int64
		// Browser save file last modified time
		browserModTime := r.Header.Get("If-Modified-Since")
		if strings.IsNotBlank(browserModTime) {
			if v, err := times.ParseUnixGMT(browserModTime); err == nil {
				browserModTimeUnix = v
			}
		}
		if browserModTimeUnix < fileModTimeUnix {
			// Actual file last modified time
			fileModTime := times.FormatUnixGMT(fileModTimeUnix)
			// Tell the browser not to use cache
			w.Header().Set("last-modified", fileModTime)
			return false
		} else {
			var content string
			if c, err := files.Read(filename); err == nil {
				content = c
			}
			if me.isSec(content) {
				var lastLoginTimeUnix int64
				s := newSession4Redis(w, r)
				if p, err := s.Principal(); err == nil {
					if strings.IsNotBlank(p.LoginTime) {
						if i, err := strconv.Atoi(p.LoginTime); err == nil {
							lastLoginTimeUnix = int64(i)
						}
					}
				}
				if browserModTimeUnix < lastLoginTimeUnix {
					// Actual last login time
					lastLoginTime := times.FormatUnixGMT(lastLoginTimeUnix)
					// Tell the browser not to use cache
					w.Header().Set("Last-Modified", lastLoginTime)
					return false
				}
			}
			if me.isInclude(content) {
				var includeFileModTimeUnix int64
				directives := make([]directiveInfo, 0)
				directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveArgEnd, directiveIncludeEnd, directives)
				for _, v := range directives {
					if val, err := files.ModTime(v.argValue); err == nil {
						if includeFileModTimeUnix < val {
							includeFileModTimeUnix = val
						}
					}
				}
				if browserModTimeUnix < includeFileModTimeUnix {
					// The actual last modification time of the include file
					includeFileModTime := times.FormatUnixGMT(includeFileModTimeUnix)
					// Tell the browser not to use cache
					w.Header().Set("last-modified", includeFileModTime)
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

func (me *htmlServeMux) parseFile(w http.ResponseWriter, r *http.Request, content string) string {
	content = me.parseIncludeFile(content)
	content = me.replaceAssets(content)
	content = me.parseIfFile(content)
	content = me.parseTagAttrFile(content, tagAttrHref)
	content = me.parseTagAttrFile(content, tagAttrSrc)
	content = me.parseTagAttrFile(content, tagAttrAction)
	content = me.parseTagTextFile(content, tagTextTitle)
	if me.isSec(content) {
		content = me.parseSecUserFile(w, r, content)
		content = me.parseSecLoginFile(w, r, content)
		content = me.parseSecIsPermissionFile(w, r, content)
		content = me.parseSecIsAnyPermissionFile(w, r, content)
	}
	return content
}

func (me *htmlServeMux) parseIncludeFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveIncludeBegin, directiveArgEnd, directiveIncludeEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		v, err := files.Read(directives[i].argValue)
		if err != nil {
			logger.Error(err.Error())
			continue
		}

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
	return content
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

func (me *htmlServeMux) parseTagAttrFile(content, attr string) string {
	tags := make([]tagInfo, 0)
	tags = me.buildTagInfo(content, attr, tags)
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

func (me *htmlServeMux) parseSecLoginFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecLoginBegin, directiveArgEnd, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		isLogin := false
		s := newSession4Redis(w, r)
		if p, err := s.Principal(); err == nil {
			if strings.IsNotBlank(p.Id) {
				isLogin = true
			}
		}
		if directives[i].argValue == "true" && !isLogin {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
		if directives[i].argValue == "false" && isLogin {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseSecUserFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecUserBegin, directiveArgEnd, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if directives[i].argValue == "name" {
			s := newSession4Redis(w, r)
			p, err := s.Principal()
			if err != nil {
				logger.Error(err.Error())
				break
			}
			content = strings.Replace(content, directives[i].statement, p.LoginName, -1)
			break
		}
	}
	return content
}

func (me *htmlServeMux) parseSecIsPermissionFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecIsPermissionBegin, directiveArgEnd, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if strings.IsNotBlank(directives[i].argValue) {
			isLogin := false
			isContains := false
			s := newSession4Redis(w, r)
			if p, err := s.Principal(); err == nil {
				if strings.IsNotBlank(p.Id) {
					isLogin = true
				}
				if strings.Contains(p.Permissions, directives[i].argValue) {
					isContains = true
				}
			}
			if !isLogin || !isContains {
				content = strings.Replace(content, directives[i].statement, "", -1)
			}
			break
		}
	}
	return content
}

func (me *htmlServeMux) parseSecIsAnyPermissionFile(w http.ResponseWriter, r *http.Request, content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, directiveSecIsAnyPermissionBegin, directiveArgEnd, directiveSecEnd, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if strings.IsNotBlank(directives[i].argValue) {
			isLogin := false
			isContains := false
			s := newSession4Redis(w, r)
			if p, err := s.Principal(); err == nil {
				if strings.IsNotBlank(p.Id) {
					isLogin = true
				}
				avs := strings.Split(directives[i].argValue, ",")
				if strings.ContainsSliceAny(p.Permissions, avs) {
					isContains = true
				}
			}
			if !isLogin || !isContains {
				content = strings.Replace(content, directives[i].statement, "", -1)
			}
			break
		}
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
			argValue = Conf.Template.Dir + argValue
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
		case directiveSecUserBegin:
			directive = "sec"
			argKey = "user"
		case directiveSecIsPermissionBegin:
			directive = "sec"
			argKey = "isPermission"
		case directiveSecIsAnyPermissionBegin:
			directive = "sec"
			argKey = "isAnyPermission"
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

func (me *htmlServeMux) buildTagInfo(content, attr string, tags []tagInfo) []tagInfo {
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
			filename := Conf.Template.Dir + dstVal
			if files.IsExist(filename) {
				if c, err := files.Read(filename); err == nil {
					title := strings.Slice(content, srcBegin+len(tagBeginPre), srcEnd)
					newstmt = strings.Replace(newstmt, title, title+c, 1)
					dst := dstBeginPre + dstVal + tagAttrEnd
					newstmt = strings.Replace(newstmt, dst, "", 1)
				}
			}
		}

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
