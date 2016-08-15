// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/maps"
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
	statement string            // <!--#include file="/footer.html" params="home"-->test<!--#endinclude-->
	directive string            // include
	attr      map[string]string // file="/footer.html" params="home"
	body      string            // test
	begin     string            // <!--#include
	mid       string            // -->
	end       string            // <!--#endinclude-->
}

/*
type directiveInfo struct {
	statement string // {%if eq .params `home`%}cur{%end%}
	directive string // if
	attr      string // eq .params `home`
	body      string // cur
	begin     string // {%if
	mid       string // %}
	end       string // {%end%}
}

type directiveInfo struct {
	statement  string // <!--#settings project="sys" module="user" title="User"-->test<!--#endsettings-->
	directive  string // settings
	attr       string // project="sys" module="user" title="User"
	body       string // test
	begin      string // <!--#settings
	mid        string // -->
	end        string // <!--#endsettings-->
}
*/

type tagInfo struct {
	statement string   // <link rel="shortcut icon" href="/favicon.ico" go:href="{%assets%}/favicon.ico">
	newstmt   string   // <link rel="shortcut icon" href="/static/favicon.ico">
	tag       string   // link
	goattr    string   // href
	attr      []string // rel="shortcut icon" href="/favicon.ico" go:href="{%assets%}/favicon.ico"
	begin     int      // postion:<
	end       int      // postion: href="
}

type tagTextInfo struct {
	statement string   // <title go:title="/title.html">login</title>
	newstmt   string   // <title>login-appendTitle</title>
	tag       string   // title
	goattr    string   // title
	attr      []string // go:title="/title.html"
	body      string   // login
	begin     int      // postion:<
	mid       int      // postion:>
	end       int      // postion:</
}

var hsm = &htmlServeMux{
	templates: make(map[string]*templateInfo),
}

var htmlMutex sync.Mutex

var ver string = "ver=1"

func (me *htmlServeMux) compile() error {
	options := Conf.Html
	fver := options.Dir + "/version.html"
	if files.IsExist(fver) {
		if c, err := files.Read(fver); err == nil {
			ver = c
		}
	}
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
					if err := me.compile(); err != nil {
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
	content = strings.Replace(content, tagAssetsVer, ver, -1)
	return strings.Replace(content, tagAssetsUploads, Conf.Upload.URL, -1)
}

func (me *htmlServeMux) replaceSettings(content string, settings map[string]string) string {
	if settings != nil {
		content = strings.Replace(content, tplBegin+attrProject+tplEnd, settings[attrProject], -1)
		content = strings.Replace(content, tplBegin+attrModule+tplEnd, settings[attrModule], -1)
		content = strings.Replace(content, tplBegin+attrTitle+tplEnd, settings[attrTitle], -1)
	}
	return content
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
	if strings.Index(content, drtBegin+drtInclude) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isSettings(content string) bool {
	if strings.Index(content, drtBegin+drtSettings) > 0 {
		return true
	}
	return false
}

func (me *htmlServeMux) isTplIf(content string) bool {
	if strings.Index(content, tplBegin+drtIf) > 0 {
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
				directives = me.buildDirectiveInfo(content, drtInclude, directives)
				for _, v := range directives {
					if val, err := files.ModTimeUnix(v.attr[attrFile]); err == nil {
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
	content, settings := me.parseSettingsFile(content)
	content, i, m := me.parseIncludeFile(content, settings)
	//content = me.replaceAssets(content)
	content = me.parseIfFile(content)
	content = me.parseProfileFile(content)
	content = me.parseTagDataAttrFile(content)
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

func (me *htmlServeMux) parseSettingsFile(content string) (string, map[string]string) {
	if me.isSettings(content) {
		directives := make([]directiveInfo, 0)
		directives = me.buildDirectiveInfo(content, drtSettings, directives)
		if len(directives) > 0 {
			settings := map[string]string{
				attrProject: directives[0].attr[attrProject],
				attrModule:  directives[0].attr[attrModule],
				attrTitle:   directives[0].attr[attrTitle],
			}
			content = me.replaceSettings(content, settings)
			return content, settings
		}
	}
	return content, nil
}

func (me *htmlServeMux) parseIncludeFile(content string, settings map[string]string) (string, []string, int64) {
	includes := []string{}
	lastModified := int64(0)
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, drtInclude, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		filename := directives[i].attr[attrFile]
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

		v = me.replaceSettings(v, settings)

		params := directives[i].attr[attrParams]
		paramsMap := maps.ParseURLQuery(params)
		if len(paramsMap) > 0 {
			for key, value := range paramsMap {
				s := tplBegin + "." + attrParams + "." + key + tplEnd
				v = strings.Replace(v, s, value, -1)
			}
		} else {
			v = strings.Replace(v, tagParams, params, -1)
			v = strings.Replace(v, tagParams2, params, -1)
		}

		if me.isTplIf(v) {
			ifparams := make([]directiveInfo, 0)
			ifparams = me.buildTplDirectiveInfo(v, drtIf, ifparams)
			for p := len(ifparams) - 1; p >= 0; p-- {
				if len(paramsMap) > 0 {
					for key, value := range paramsMap {
						pKey := ifparams[p].attr["0"] + " " + ifparams[p].attr["1"]
						pVal := "`" + value + "`"
						if pKey == tplEqParams+"."+key && pVal == ifparams[p].attr["2"] {
							v = strings.Replace(v, ifparams[p].statement, ifparams[p].body, -1)
						} else {
							v = strings.Replace(v, ifparams[p].statement, "", -1)
						}
					}
				} else {
					pKey := ifparams[p].attr["0"] + " " + ifparams[p].attr["1"]
					pVal := "`" + directives[i].attr[attrParams] + "`"
					if pKey == tplEqParams && pVal == ifparams[p].attr["2"] {
						v = strings.Replace(v, ifparams[p].statement, ifparams[p].body, -1)
					} else {
						v = strings.Replace(v, ifparams[p].statement, "", -1)
					}
				}
			}
		}

		content = strings.Replace(content, directives[i].statement, v, -1)
	}
	return content, includes, lastModified
}

func (me *htmlServeMux) parseIfFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, drtIf, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if directives[i].attr[attrExpr] == "false" {
			content = strings.Replace(content, directives[i].statement, "", -1)
		}
	}
	return content
}

func (me *htmlServeMux) parseProfileFile(content string) string {
	directives := make([]directiveInfo, 0)
	directives = me.buildDirectiveInfo(content, drtProfile, directives)
	for i := len(directives) - 1; i >= 0; i-- {
		if !profile.Accepts(directives[i].attr[attrAccepts]) {
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

func (me *htmlServeMux) parseTagDataAttrFile(content string) string {
	tags := make([]tagInfo, 0)
	tags = me.buildTagDataAttrInfo(content, tags)
	for i := len(tags) - 1; i >= 0; i-- {
		content = me.parseTagAttrFile(content, tags[i].goattr)
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

func (me *htmlServeMux) buildDirectiveInfo(content, directive string, directives []directiveInfo) []directiveInfo {
	dBegin := drtBegin + directive                    // <!--#include
	dEnd := drtBegin + drtEndKey + directive + drtEnd // <!--#endinclude-->
	outs := strings.Betweens(content, dBegin, dEnd)
	for _, out := range outs {
		statement := dBegin + out + dEnd
		di := directiveInfo{
			statement: statement,
			directive: directive,
			attr:      make(map[string]string, 0),
			begin:     dBegin,
			mid:       drtEnd,
			end:       dEnd,
		}
		di.body = strings.Between(statement, di.mid, di.end)
		switch directive {
		case drtIf:
			di.attr[attrExpr] = me.getAttrVal(statement, attrExpr)
		case drtProfile:
			di.attr[attrAccepts] = me.getAttrVal(statement, attrAccepts)
		case drtInclude:
			di.attr[attrFile] = me.getAttrVal(statement, attrFile)
			di.attr[attrParams] = me.getAttrVal(statement, attrParams)
			if strings.IsBlank(di.attr[attrFile]) {
				continue
			}
			di.attr[attrFile] = Conf.Html.Dir + di.attr[attrFile]
			if !files.IsExist(di.attr[attrFile]) {
				continue
			}
		case drtSettings:
			di.attr[attrProject] = me.getAttrVal(statement, attrProject)
			di.attr[attrModule] = me.getAttrVal(statement, attrModule)
			di.attr[attrTitle] = me.getAttrVal(statement, attrTitle)
		}
		directives = append(directives, di)
	}
	return directives
}

func (me *htmlServeMux) getAttrVal(statement, attr string) string {
	return strings.Between(statement, attr+`="`, `"`)
}

func (me *htmlServeMux) buildTplDirectiveInfo(content, directive string, directives []directiveInfo) []directiveInfo {
	dBegin := tplBegin + directive        // {%if
	dEnd := tplBegin + drtEndKey + tplEnd // {%end%}
	outs := strings.Betweens(content, dBegin, dEnd)
	for _, out := range outs {
		statement := dBegin + out + dEnd
		di := directiveInfo{
			statement: statement,
			directive: directive,
			attr:      make(map[string]string, 0),
			begin:     dBegin,
			mid:       tplEnd,
			end:       dEnd,
		}
		di.body = strings.Between(statement, di.mid, di.end)
		switch directive {
		case drtIf:
			attr := strings.Before(out, tplEnd)
			attrs := strings.FieldsSpace(attr)
			i := 0
			for _, v := range attrs {
				if strings.IsNotBlank(v) {
					di.attr[strconv.Itoa(i)] = v
					i++
				}
			}
		}
		directives = append(directives, di)
	}
	return directives
}

func (me *htmlServeMux) buildTagDataAttrInfo(content string, tags []tagInfo) []tagInfo {
	dataBeginPre := tagAttrPre + tagAttrData // go:data-
	outs := strings.Betweens(content, dataBeginPre, tagAttrPost)
	for _, out := range outs {
		attr := tagAttrData + out // data-permissions
		tags = me.buildTagAttrInfo(content, attr, tags)
	}
	return tags
}

func (me *htmlServeMux) buildTagAttrInfo(content, attr string, tags []tagInfo) []tagInfo {
	pos := 0
	goAttr := tagAttrPre + attr + tagAttrPost // go:href=
	for {
		goAttrBegin := strings.IndexStart(content, goAttr, pos)
		if goAttrBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		begin := strings.IndexForward(content, tagBeginPre, goAttrBegin)
		if begin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		end := strings.IndexStart(content, tagEndPre, goAttrBegin)
		if end == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		pos = end
		statement := strings.Slice(content, begin, end+len(tagEndPre))
		newstmt := statement
		tag := ""
		attrs := []string{}
		statements := strings.FieldsSpace(statement)
		for i, v := range statements {
			if i == 0 {
				tag = v
			}
			if strings.Contains(v, tagAttrPost) {
				attrs = append(attrs, v)
			}
			if strings.HasPrefix(v, attr+tagAttrPost) {
				newstmt = strings.Replace(newstmt, v, "", 1)
			}
		}
		for _, v := range statements {
			if strings.HasPrefix(v, goAttr) {
				val := strings.After(v, tagAttrPre)
				val = me.replaceAssets(val)
				newstmt = strings.Replace(newstmt, v, val, 1)
			}
		}
		ti := tagInfo{
			statement: statement,
			newstmt:   newstmt,
			tag:       tag,
			goattr:    attr,
			attr:      attrs,
			begin:     begin,
			end:       end,
		}
		tags = append(tags, ti)
	}
	return tags
}

func (me *htmlServeMux) buildTagTextInfo(content, attr string, tags []tagTextInfo) []tagTextInfo {
	pos := 0
	goAttr := tagAttrPre + attr + tagAttrPost // go:text=
	for {
		goAttrBegin := strings.IndexStart(content, goAttr, pos)
		if goAttrBegin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		begin := strings.IndexForward(content, tagBeginPre, goAttrBegin)
		if begin == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		mid := strings.IndexStart(content, tagEndPre, goAttrBegin)
		if mid == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		end := strings.IndexStart(content, tagTextEndPre, mid)
		if end == -1 {
			if pos == 0 {
				return tags
			}
			break
		}
		pos = end
		statement := strings.Slice(content, begin, end+len(tagTextEndPre))
		newstmt := statement
		body := strings.Slice(content, mid+len(tagEndPre), end)
		tag := ""
		attrVal := ""
		attrs := []string{}
		as := strings.Slice(content, begin, mid+len(tagEndPre))
		statements := strings.FieldsSpace(as)
		for i, v := range statements {
			if i == 0 {
				tag = v
			}
			if strings.Contains(v, tagAttrPost) {
				attrs = append(attrs, v)
			}
			if strings.HasPrefix(v, attr+tagAttrPost) {
				newstmt = strings.Replace(newstmt, v, "", 1)
			}
		}
		for _, v := range statements {
			if strings.HasPrefix(v, goAttr) {
				attrVal = strings.Slice(v, len(goAttr), len(v)-1)
				val := strings.After(v, tagAttrPre)
				newstmt = strings.Replace(newstmt, v, val, 1)
			}
		}
		if strings.IsNotBlank(attrVal) {
			switch attr {
			case tagTextTitle:
				filename := Conf.Html.Dir + attrVal
				if files.IsExist(filename) {
					if c, err := files.Read(filename); err == nil {
						newstmt = strings.Replace(newstmt, body, body+c, 1)
					}
				}
			}
		}

		newstmt = me.replaceAssets(newstmt)
		tti := tagTextInfo{
			statement: statement,
			newstmt:   newstmt,
			tag:       tag,
			goattr:    attr,
			attr:      attrs,
			body:      body,
			begin:     begin,
			mid:       mid,
			end:       end,
		}
		tags = append(tags, tti)
	}
	return tags
}
