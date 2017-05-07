// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreIncludeHead = `<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">

<link rel="shortcut icon" go:href="{%statics%}/favicon.ico">

<!-- Bootstrap core CSS -->
<link go:href="{%developers%}/ui/comm/css/comm.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/comm/css/navbar.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/bootstrap/css/bootstrap.min.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/bootstrap/css/bootstrap-theme.min.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/bootstrap-fileinput/css/fileinput.min.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/jquery-ui/jquery-ui.min.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/jquery-ui-timepicker/jquery-ui-timepicker-addon.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/jquery-ztree/css/zTreeStyle/zTreeStyle.css?{%ver%}" rel="stylesheet">
<link go:href="{%developers%}/ui/jquery-qtip/jquery.qtip.min.css?{%ver%}" rel="stylesheet">

<!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
<!--[if lt IE 9]>
<script go:src="{%developers%}/ui/html5shiv/html5shiv.min.js?{%ver%}"></script>
<script go:src="{%developers%}/ui/respond.js/respond.min.js?{%ver%}"></script>
<![endif]-->

<script type="text/javascript" go:src="{%developers%}/ui/jquery/jquery.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery/jquery.ajaxSetup.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/bootstrap/js/bootstrap.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/bootstrap-fileinput/js/fileinput.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/bootstrap-fileinput/js/fileinput_locale_zh.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-ui/jquery-ui.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-ui/datepicker-zh-CN.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-ui-timepicker/jquery-ui-timepicker-addon.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-ui-timepicker/jquery-ui-timepicker-zh-CN.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-ztree/js/jquery.ztree.all.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-validation/jquery.validate.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-validation/jquery.validate.defaults.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-validation/jquery.validate.methods.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-validation/messages_zh.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-form/jquery.form.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-qtip/jquery.qtip.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-dateFormat/jquery-dateFormat.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/jquery-cookie/jquery.cookie.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/js-base64/base64.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/md5/md5.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/purl/purl.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/ckeditor/ckeditor.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/handlebars/handlebars.min.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/handlebars/handlebars.helpers.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/comm/js/jquery.form.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/comm/js/jquery.page.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/comm/js/jquery.util.js?{%ver%}"></script>
<script type="text/javascript" go:src="{%developers%}/ui/comm/js/jquery.tree.js?{%ver%}"></script>
<script go:type="text/javascript">
	var profile = "{%profile%}";
	var apis = "{%apis%}";
	var assets = "{%assets%}";
	var statics = "{%statics%}";
	var developers = "{%developers%}";
	var operations = "{%operations%}";
	var uploads = "{%uploads%}";
	var state="list";
</script>
`
