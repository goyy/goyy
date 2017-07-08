// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesHome = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="sys" module="user" title="<%message "tmpl.home.title"%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/core/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/home/js/home.js?{%ver%}"></script>
</head>
<body>
<!--#include file="/core/include/header.html" param="{%project%}"--><!--#endinclude-->
<div class="container-fluid">
<div class="content" class="row-fluid">
	<div class="tabbable">
		<ul id="{%.prefix%}NavTabs" class="nav nav-tabs dn" >
			<li class="active">
				<a href="#e-info" data-toggle="tab"><%message "tmpl.home.tab.info"%></a>
			</li>
			<li>
				<a href="#e-passwd" data-toggle="tab"><%message "tmpl.home.tab.passwd"%></a>
			</li>
		</ul>
		<div class="tab-content">
			<!--#include file="/ui/include/alert.html" param="{%.prefix%}-info"--><!--#endinclude-->
			<div id="{%.prefix%}-info" class="tab-pane active">
			</div>
			<script id="{%.prefix%}-info-template" type="text/x-handlebars-template">
				<br/>
				<div class="form-group">
					<label class="control-label"><%message "tmpl.home.form.login"%></label>
					<div>
						<input class="form-control" readonly="true" value="{{loginName}}">
					</div>
				</div>
				<div class="form-group">
					<label class="control-label"><%message "tmpl.home.form.name"%></label>
					<div>
						<input class="form-control" readonly="true" value="{{name}}">
					</div>
				</div>
				<div class="form-group">
					<label class="control-label"><%message "tmpl.home.form.time"%></label>
					<div>
						<input class="form-control" readonly="true" value="{{uyymdhms loginTime}}">
					</div>
				</div>
			</script>
			<div id="{%.prefix%}-passwd" class="tab-pane">
				<form id="eFormPasswd" role="form" method="post" go:action="{%apis%}/sys/user/repwd">
					<br/>
					<!--#include file="/ui/include/alert.html" param="{%.prefix%}Form"--><!--#endinclude-->
					<div class="form-group">
						<label class="control-label"><%message "tmpl.home.form.oldpwd"%></label>
						<div>
							<input type="password" class="form-control" id="oldPasswd" name="oldPasswd" autocomplete="off">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.home.form.newpwd"%></label>
						<div>
							<input type="password" class="form-control" id="newPasswd" name="newPasswd">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.home.form.okpwd"%></label>
						<div>
							<input type="password" class="form-control" id="okPasswd" name="okPasswd">
						</div>
					</div>
					<button type="submit" class="btn btn-primary"><%message "tmpl.home.form.save"%></button>
					<button type="button" class="btn btn-default" onclick="history.go(-1)"><%message "tmpl.home.form.goback"%></button>
				</form>
			</div>
		</div>
	</div>
</div>
<!--#include file="/core/include/footer.html"--><!--#endinclude-->
</div>
</body>
</html>
`
