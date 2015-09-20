// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"math"
)

const (
	defaultEntityPrefix = ""
	defaultStatus       = 200
	default404Body      = "404 page not found"
	default405Body      = "405 method not allowed"

	directiveArgEnd             = `"-->`
	directiveIncludeBegin       = `<!--#include file="`
	directiveIncludeEnd         = `<!--#endinclude-->`
	directiveIfBegin            = `<!--#if expr="`
	directiveIfEnd              = `<!--#endif-->`
	directiveSecLoginBegin      = `<!--#sec login="`
	directiveSecUserBegin       = `<!--#sec user="`
	directiveSecHasRoleBegin    = `<!--#sec hasRole="`
	directiveSecHasAnyRoleBegin = `<!--#sec hasAnyRole="`
	directiveSecEnd             = `<!--#endsec-->`

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

	principalId          = "_gsessionid_principal_id_"
	principalName        = "_gsessionid_principal_name_"
	principalLoginName   = "_gsessionid_principal_login_name_"
	principalLoginTime   = "_gsessionid_principal_login_time_"
	principalPermissions = "_gsessionid_principal_permissions_"
)
