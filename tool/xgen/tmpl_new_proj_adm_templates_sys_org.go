// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesSysOrg = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="sys" module="org" title="<%message "tmpl.sys.org.title"%>"--><!--#endsettings-->
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
		<!--#include file="/core/comm/navtabs.html" param="{%.prefix%}"--><!--#endinclude-->
		<div class="tab-content">
			<div id="{%.prefix%}ListContent" class="tab-pane active">
				<br/>
				<div class="well well-sm">
					<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="sAreaId" url="/sys/area/tree"--><!--#endinclude-->
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/index">
						<input type="hidden" id="sParentId" name="sParentId" value="root" />
						<input type="hidden" id="sOrdinalOA" name="sOrdinalOA" value="10" />
						<input type="hidden" id="sDeletionU0" name="sDeletionU0" value="0" />
						<input type="hidden" id="sPageNoTR" name="sPageNoTR" value="1" />
						<input type="hidden" id="sPageSizeTR" name="sPageSizeTR" value="" />
						<label class="control-label"><%message "tmpl.sys.area.title"%></label>
						<div id="{%.prefix%}sAreaIdFormTree" class="form-group" data-val="">
						</div>
						<div class="form-group">
							<label class="control-label text-right"><%message "tmpl.sys.col.name"%></label>
							<input class="form-control" id="sNameLK" name="sNameLK" value="">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary"><%message "tmpl.sys.btn.query"%></button>
					</form>
				</div>
				<!--#include file="/core/comm/alert.html" param="{%.prefix%}List"--><!--#endinclude-->
				<!-- 树导航栏 -->
				<!--#include file="/core/comm/breadcrumb.html" param="{%.prefix%}"--><!--#endinclude-->
				<table class="table table-bordered table-condensed">
					<thead>
						<tr>
							<th class="f-w50"><%message "tmpl.sys.col.id"%></th>
							<th><%message "tmpl.sys.col.name"%></th>
							<th><%message "tmpl.sys.area.title"%></th>
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
							<td>{{areafullname areaId}}</td>
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
			<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="areaId" url="/sys/area/tree"--><!--#endinclude-->
			<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="parentId" url="/sys/org/tree"--><!--#endinclude-->
			<script id="{%.prefix%}FormTemplate" type="text/x-handlebars-template">
				<form id="{%.prefix%}Form" method="post" go:action="{%apis%}/{%project%}/{%module%}/save">
					<br/>
					<!--#include file="/core/comm/alert.html" param="{%.prefix%}Form"--><!--#endinclude-->
					<input type="hidden" id="id" name="id" value="{{id}}" />
					<input type="hidden" id="fullname" name="fullname" value="{{fullname}}" />
					{{#if (eq state 'add')}}
					<input type="hidden" id="parentId" name="parentId" value="{{parentId}}" />
					{{/if}}
					<input type="hidden" id="parentIds" name="parentIds" value="{{parentIds}}" />
					<input type="hidden" id="parentNames" name="parentNames" value="{{parentNames}}" />
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.area"%></label>
						<div id="{%.prefix%}areaIdFormTree" class="form-inline" data-val="{{areaId}}">
						</div>
					</div>
					{{#if (eq state 'edit')}}
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.parent"%></label>
						<div id="{%.prefix%}parentIdFormTree" class="form-inline" data-val="{{parentId}}">
						</div>
					</div>
					{{/if}}
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.name"%></label>
						<div>
							<input class="form-control required" id="name" name="name" value="{{name}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.code"%></label>
						<div>
							<input class="form-control required" id="code" name="code" value="{{code}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.ordinal"%></label>
						<div>
							<input class="form-control required" id="ordinal" name="ordinal" value="{{ordinal}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.sys.col.memo"%></label>
						<div>
							<input class="form-control" id="memo" name="memo" value="{{memo}}"/>
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
