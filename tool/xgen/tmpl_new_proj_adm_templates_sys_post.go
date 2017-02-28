// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesSysPost = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="sys" module="post" title="<%message "tmpl.sys.post.title"%>"--><!--#endsettings-->
<!--#with prefix="e" --><!--#endwith-->
<title go:title="/title.html">{%title%}</title>
<!--#include file="/core/include/head.html"--><!--#endinclude-->

<script type="text/javascript" go:src="{%statics%}/{%project%}/js/{%module%}.js?{%ver%}"></script>
</head>
<body>
<!--#include file="/core/include/header.html" param="{%project%}"--><!--#endinclude-->
<div class="container-fluid">
<div class="content" class="row-fluid">
<!--#include file="/core/comm/page.html" param="{%.prefix%}"--><!--#endinclude-->
<!--#include file="/core/comm/action.html" param="{%.prefix%}"--><!--#endinclude-->
<!--#include file="/core/comm/disable.html" param="{%.prefix%}"--><!--#endinclude-->
	<div class="tabbable">
		<ul id="{%.prefix%}NavTabs" class="nav nav-tabs dn" >
			<li go:data-permissions="{%project%}:{%module%}:view" class="active">
				<a href="#eListContent" go:data-state="list" onclick="{%.prefix%}Action(this,'')">{%title%}<%message "tmpl.sys.list.title"%></a>
			</li>
			<li go:data-permissions="{%project%}:{%module%}:add">
				<a href="#eFormContent" go:data-state="add" data-template="{%.prefix%}FormTemplate" go:data-url="{%apis%}/{%project%}/{%module%}/add"
					onclick="{%.prefix%}Action(this,'root',{%.prefix%}PostLoadForm)" >{%title%}<span id="eTabTitle"><%message "tmpl.sys.add.title"%></span></a>
			</li>
		</ul>
		<script type="text/javascript">
			$(function(){
				$("#eNavTabs").permission();
			});
		</script>
		<div class="tab-content">
			<div id="{%.prefix%}ListContent" class="tab-pane active">
				<br/>
				<div class="well well-sm">
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/index">
						<input type="hidden" id="sParentId" name="sParentId" value="root" />
						<input type="hidden" id="sOrdinalOA" name="sOrdinalOA" value="10" />
						<input type="hidden" id="sDeletionU0" name="sDeletionU0" value="0" />
						<input type="hidden" id="sPageNoTR" name="sPageNoTR" value="1" />
						<input type="hidden" id="sPageSizeTR" name="sPageSizeTR" value="" />
						<div class="form-group">
							<label class="control-label"><%message "tmpl.sys.col.name"%></label>
							<input class="form-control" id="sNameLK" name="sNameLK" value="">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary"><%message "tmpl.sys.btn.query"%></button>
					</form>
				</div>
				<div id="{%.prefix%}ListAlert" class="alert alert-warning fade in hidden">
					<button id="msgBtnClose" type="button" class="close" >&times;</button>
					<strong id="message"></strong>
				</div>
				<!-- 树导航栏 -->
				<!--#include file="/core/comm/breadcrumb.html" param="{%.prefix%}"--><!--#endinclude-->
				<table class="table table-bordered table-condensed">
					<thead>
						<tr>
							<th class="f-w50"><%message "tmpl.sys.col.id"%></th>
							<th><%message "tmpl.sys.col.name"%></th>
							<th><%message "tmpl.sys.col.fullname"%></th>
							<th><%message "tmpl.sys.col.code"%></th>
							<th class="f-w50"><%message "tmpl.sys.col.ordinal"%></th>
							<th><%message "tmpl.sys.col.memo"%></th>
							<th class="f-w150"><%message "tmpl.sys.col.created"%></th>
							<th><%message "tmpl.sys.col.opt"%></th>
						</tr>
					</thead>
					<tbody id="{%.prefix%}ListData">
					
					</tbody>
					<script id="{%.prefix%}ListTemplate" type="text/x-handlebars-template">
					{{#each slice}}
						<tr{{#if (divisible @index 2)}} class="active"{{/if}}>
							<td>
								<a href="#eFormContent" go:data-state="show" go:data-template="{%.prefix%}FormTemplate" go:data-url="{%apis%}/{%project%}/{%module%}/show" 
									onclick="{%.prefix%}Action(this,'{{id}}',{%.prefix%}PostLoadForm)" title="{{id}}" >{{abbr id 5}}</a>
							</td>
							<td>
								<a href="#eListContent" onclick="{%.prefix%}BreadcrumbPage('{{id}}')">{{name}}</a>{{!配合树导航栏查询使用}}
							</td>
							<td>{{fullname}}</td>
							<td>{{code}}</td>
							<td>{{ordinal}}</td>
							<td>{{memo}}</td>
							<td>{{uyymdhms created}}</td>
							<td>								
								<a href="#eFormContent" go:data-state="edit" go:data-template="{%.prefix%}FormTemplate" go:data-url="{%apis%}/{%project%}/{%module%}/edit" onclick="{%.prefix%}Action(this,'{{id}}',{%.prefix%}PostLoadForm)" go:data-permissions="{%project%}:{%module%}:edit"><%message "tmpl.sys.btn.edit"%></a>
								<a href="#disable" go:data-state="disable" go:data-url="{%apis%}/{%project%}/{%module%}/disable" onclick="disable(this,'{{id}}',{%.prefix%}Page)" go:data-permissions="{%project%}:{%module%}:disable"><%message "tmpl.sys.btn.disable"%></a>
							</td>
						</tr>
					{{/each}}
					</script>
				</table>
                <ul id="{%.prefix%}ListPagination" class="pagination"></ul>
			</div>
			<div id="{%.prefix%}FormContent" class="tab-pane">
				
			</div>
			<script id="{%.prefix%}FormTemplate" type="text/x-handlebars-template">
				<form id="{%.prefix%}Form" method="post" go:action="{%apis%}/{%project%}/{%module%}/save">
					<br/>
					<!--#include file="/core/comm/alert.html" param="{%.prefix%}Form"--><!--#endinclude-->
					<input type="hidden" id="id" name="id" value="{{id}}" />
					<input type="hidden" id="parentId" name="parentId" value="{{parentId}}" />
					<div class="form-inline">
						<div class="form-group">
							<label class="control-label">*<%message "tmpl.sys.col.name"%></label>
							<div>
								<input class="form-control required" id="name" name="name" value="{{name}}"/>
							</div>
						</div>
						<div class="form-group">
							<label class="control-label">*<%message "tmpl.sys.col.genre"%></label>
							<div>
								<select class="form-control" id="genre" name="genre" value="{{genre}}"></select>
							</div>
						</div>
						<div class="form-group">
							<label class="control-label">*<%message "tmpl.sys.col.ordinal"%></label>
							<div>
								<input class="form-control required" id="ordinal" name="ordinal" value="{{ordinal}}"/>
							</div>
						</div>
						<div class="form-group">
							<label class="control-label">*<%message "tmpl.sys.col.memo"%></label>
							<div>
								<input class="form-control required" id="memo" name="memo" value="{{memo}}"/>
							</div>
						</div>
					</div>
					<br/>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.menu"%></label>
						<div>
							<input type="hidden" id="menuIds" name="menuIds" value="{{menuIds}}" />
							<ul id="treeMenu" class="ztree"></ul>
						</div>
					</div>
					{{#if (ne state 'show')}}
					<button type="submit" class="btn btn-primary" go:data-permissions="{%project%}:{%module%}:edit,{%project%}:{%module%}:add"><%message "tmpl.sys.btn.save"%></button>
					{{/if}}
				</form>
			</script>
		</div>
	</div>
</div>
<!--#include file="/core/include/footer.html"--><!--#endinclude-->
</div>
</body>
</html>
`
