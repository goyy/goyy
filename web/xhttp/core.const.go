// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"math"
)

const (
	defaultStatus  = 200
	default404Body = "404 page not found"
	default403Body = "403 Forbidden"
	default405Body = "405 method not allowed"
	default500Body = "500 Internal Server Error"

	directiveArgEnd                  = `"-->`
	directiveIncludeBegin            = `<!--#include file="`
	directiveIncludeParamBegin       = `" param="`
	directiveIncludeEnd              = `<!--#endinclude-->`
	directiveIfBegin                 = `<!--#if expr="`
	directiveIfEnd                   = `<!--#endif-->`
	directiveSecLoginBegin           = `<!--#sec login="`
	directiveSecUserBegin            = `<!--#sec user="`
	directiveSecIsPermissionBegin    = `<!--#sec isPermission="`
	directiveSecIsAnyPermissionBegin = `<!--#sec isAnyPermission="`
	directiveSecEnd                  = `<!--#endsec-->`

	tplBegin  = "{{if param `"
	tplArgEnd = "`}}"
	tplEnd    = "{{end}}"

	tagBeginPre         = `<`
	tagEndPre           = `>`
	tagTextEndPre       = `</`
	tagTextTitle        = `title`
	tagAttrHref         = `href`
	tagAttrSrc          = `src`
	tagAttrAction       = `action`
	tagAttrPre          = ` go:`
	tagAttrPost         = `="`
	tagAttrEnd          = `"`
	tagStaticApis       = `{{apis}}`
	tagStaticAssets     = `{{assets}}`
	tagStaticConsumers  = `{{consumers}}`
	tagStaticOperations = `{{operations}}`

	noWritten = -1

	AbortIndex int8 = math.MaxInt8 / 2

	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)
