// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplHTMLMain = `<%range $i, $e := .Entities%><!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="<%$e.Project%>" module="<%$e.Module%>" title="<%$e.Comment%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/layout/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/{%project%}/js/{%module%}.js?{%ver%}"></script>
</head>
<body>
<!--#include file="/layout/include/header.html" param="{%project%}"--><!--#endinclude-->
<div class="container-fluid">
<div class="content" class="row-fluid">
	<!--#include file="/ui/include/page.html" param="{%.prefix%}"--><!--#endinclude-->
	<!--#include file="/ui/include/action.html" param="{%.prefix%}"--><!--#endinclude-->
	<!--#include file="/ui/include/disable.html" param="{%.prefix%}"--><!--#endinclude-->
	<div class="tabbable">
		<!--#include file="/ui/include/navtabs.html" param="{%.prefix%}"--><!--#endinclude-->
		<div class="tab-content">
			<div id="{%.prefix%}ListContent" class="tab-pane active">
				<br/>
				<div class="well well-sm">
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/index">
						<input type="hidden" id="sDeletionU0" name="sDeletionU0" value="0" />
						<input type="hidden" id="sPageNoTR" name="sPageNoTR" value="1" />
						<input type="hidden" id="sPageSizeTR" name="sPageSizeTR" value="" />
						<div class="form-group">
							<label class="control-label"><%message "html.list.sid"%></label>
							<input class="form-control" id="sIdEQ" name="sIdEQ">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary"><%message "html.list.sbtn"%></button>
					</form>
				</div>
				<!--#include file="/ui/include/alert.html" param="{%.prefix%}List"--><!--#endinclude-->
				<table class="table table-bordered table-condensed">
					<thead>
						<tr>
							<th class="f-w50"><%message "html.list.sid"%></th><%range $f := $e.Fields%>
							<th><%$f.Comment%></th><%end%>
							<th class="f-w150"><%message "html.list.created"%></th>
							<th><%message "html.list.opr"%></th>
						</tr>
					</thead>
					<tbody id="{%.prefix%}ListData"></tbody>
					<script id="{%.prefix%}ListTemplate" type="text/x-handlebars-template">
					{{#each slice}}
						<tr{{#if (divisible @index 2)}} class="active"{{/if}}>
							<td title="{{id}}">{{abbr id 5}}</td><%range $f := $e.Fields%>
							<td>{{<%$f.Name%>}}</td><%end%>
							<td>{{uyymdhms created}}</td>
							<td>
								<a href="#eFormContent" go:data-state="edit" go:data-template="{%.prefix%}FormTemplate" go:data-url="{%apis%}/{%project%}/{%module%}/edit" onclick="{%.prefix%}Action(this,'{{id}}',{%.prefix%}PostLoadForm)" go:data-permissions="{%project%}:{%module%}:edit"><%message "html.list.edit"%></a>
								<a href="#disable" go:data-state="disable" go:data-url="{%apis%}/{%project%}/{%module%}/disable" onclick="disable(this,'{{id}}',{%.prefix%}Page)" go:data-permissions="{%project%}:{%module%}:disable"><%message "html.list.del"%></a>
							</td>
						</tr>
					{{/each}}
					</script>
				</table>
				<ul id="{%.prefix%}ListPagination" class="pagination"></ul>
			</div>
			<div id="{%.prefix%}FormContent" class="tab-pane"></div>
			<script id="{%.prefix%}FormTemplate" type="text/x-handlebars-template">
				<form id="{%.prefix%}Form" method="post" go:action="{%apis%}/{%project%}/{%module%}/save">
					<br/>
					<!--#include file="/ui/include/alert.html" param="{%.prefix%}Form"--><!--#endinclude-->
					<input type="hidden" id="id" name="id" value="{{id}}" /><%range $f := $e.Fields%>
					<div class="form-group">
						<label class="control-label"><%$f.Comment%></label>
						<div>
							<input class="form-control required" id="<%$f.Name%>" name="<%$f.Name%>" value="{{<%$f.Name%>}}"/>
						</div>
					</div><%end%>
					{{#if (ne state 'show')}}
					<button type="submit" class="btn btn-primary" go:data-permissions="{%project%}:{%module%}:edit,{%project%}:{%module%}:add"><%message "html.form.save"%></button>
					{{/if}}
				</form>
			</script>
		</div>
	</div>
</div>
<!--#include file="/layout/include/footer.html"--><!--#endinclude-->
</div>
</body>
</html><%end%>
`
