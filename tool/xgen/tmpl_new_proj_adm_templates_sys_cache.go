// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesSysCache = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="sys" module="cache" title="<%message "tmpl.sys.cache.title"%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/core/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/js/{%project%}/{%module%}/{%module%}.js?{%ver%}"></script>
</head>
<body>
<!--#include file="/core/include/header.html" param="{%project%}"--><!--#endinclude-->
<div class="container-fluid">
<div class="content" class="row-fluid">
	<div class="tabbable">
		<ul id="{%.prefix%}NavTabs" class="nav nav-tabs dn" >
			<li go:data-permissions="{%project%}:{%module%}:view" class="active">
				<a href="#eListContent" go:data-state="list" onclick="{%.prefix%}Action(this,'')">{%title%}<%message "tmpl.sys.btn.show"%></a>
			</li>
		</ul>
		<script type="text/javascript">
			$(function(){
				$("#eNavTabs").permission();
			});
		</script>
		<div class="tab-content">
			<div class="tab-pane<%"<"%>%if eqindex .State%<%">"%> active<%"<"%>%end%<%">"%>">
				<br/>
				<div class="well well-sm">
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/show">
						<div class="form-group">
							<label class="control-label">KEY </label>
							<input class="form-control" id="key" name="key" value="">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary"><%message "tmpl.sys.btn.query"%></button>
					</form>
				</div>
				<table class="table table-bordered table-condensed">
					<textarea id="data" class="form-control" rows="15"></textarea>
				</table>
			</div>
		</div>
	</div>
</div>
<!--#include file="/core/include/footer.html"--><!--#endinclude-->
</div>
</body>
</html>
`
