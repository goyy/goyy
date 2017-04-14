// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesLogin = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="login" module="login" title="<%message "tmpl.login.title"%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/core/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/login/js/login.js?{%ver%}"></script>
</head>
<body>
<div class="container-fluid">
<div class="content" class="row-fluid">
	<br/><br/><br/><br/><br/><br/>
	<div class="col-md-offset-4 col-md-4 col-md-offset-4">
		<div class="panel panel-default">
			<div class="panel-heading"><%.NewProjTitle%></div>
			<div class="panel-body">
				<form id="{%.prefix%}Form" role="form" action="/signin" method="post">
					<input type="hidden" id="redirect" name="redirect" value=""/>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.login.form.name"%></label>
						<div class="form-group">
							<input type="text" id="loginName" name="loginName" class="form-control required" value="">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.login.form.pwd"%></label>
						<div class="form-group">
							<input type="password" id="passwd" name="passwd" class="form-control required">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.login.form.captcha"%></label>
						<div class="form-group form-inline">
							<input type="text" id="captcha" name="captcha" class="form-control required input-sm" />
							<img id="captchaImg" go:src="/captcha/build" onclick="resetCaptcha();" width="90" height="30" />
							<span>&nbsp;<%message "tmpl.login.form.look"%>&nbsp;&nbsp;<a href="javascript:resetCaptcha();"><%message "tmpl.login.form.change"%></a></span>
						</div>
					</div>
					<input class="btn btn-large btn-primary" type="submit" value="<%message "tmpl.login.form.submit"%>" />
				</form>
			</div>
			<div class="panel-footer">
				Copyright &copy; <%year%> <a href="http://www.<%.NewProjHost%>" target="_blank"><%.NewProjTitle%></a>
			</div>
		</div>
	</div>
</div>
</div> <!-- /container -->
</body>
</html>
`
