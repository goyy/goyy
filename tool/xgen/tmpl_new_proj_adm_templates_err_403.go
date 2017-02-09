// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesErr403 = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<link rel="shortcut icon" href="{%statics%}/favicon.ico">
<title go:title="/title.html"><%message "tmpl.err.403.title"%></title>
</head>
<body>
<h1><%message "tmpl.err.403.h1"%></h1>
<a href="javascript:void(0);" onclick="history.go(-1);"><%message "tmpl.err.goback"%></a>
</body>
</html>
`
