// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesSysUser = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta name="description" content="">
<meta name="author" content="">
<!--#settings project="sys" module="user" title="<%message "tmpl.sys.user.title"%>"--><!--#endsettings-->
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
					<form id="{%.prefix%}ListSform" class="form-inline" method="post" go:action="{%apis%}/{%project%}/{%module%}/index">
						<input type="hidden" id="sDeletionU0" name="sDeletionU0" value="0" />
						<input type="hidden" id="sCreatedOD" name="sCreatedOD" value="10" />
						<input type="hidden" id="sModifiedOD" name="sModifiedOD" value="20" />
						<input type="hidden" id="sPageNoTR" name="sPageNoTR" value="1" />
						<input type="hidden" id="sPageSizeTR" name="sPageSizeTR" value="" />
						<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="sAreaId" url="/sys/area/tree"--><!--#endinclude-->
						<div id="{%.prefix%}sAreaIdFormTree" class="form-group" data-val="">
							<label class="control-label"><%message "tmpl.sys.area.title"%></label>
						</div>
						<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="sOrgId" url="/sys/org/tree"--><!--#endinclude-->
						<div id="{%.prefix%}sOrgIdFormTree" class="form-group" data-val="">
							<label class="control-label"><%message "tmpl.sys.org.title"%></label>
						</div>
						<div class="form-group">
							<label class="control-label"><%message "tmpl.sys.col.name"%></label>
							<input class="form-control" id="sNameLK" name="sNameLK" value="">
						</div>
						<button type="button" id="{%.prefix%}ListSbtn" class="btn btn-primary"><%message "tmpl.sys.btn.query"%></button>
					</form>
				</div>
				<!--#include file="/core/comm/alert.html" param="{%.prefix%}List"--><!--#endinclude-->
				<table class="table table-bordered table-condensed">
					<thead>
						<tr>
							<th class="f-w50"><%message "tmpl.sys.col.id"%></th>
							<th><%message "tmpl.sys.col.area"%></th>
							<th><%message "tmpl.sys.col.org"%></th>
							<th><%message "tmpl.sys.col.login.name"%></th>
							<th><%message "tmpl.sys.col.name2"%></th>
							<th><%message "tmpl.sys.col.email"%></th>
							<th><%message "tmpl.sys.col.mobile"%></th>
							<th><%message "tmpl.sys.col.tel"%></th>
							<th><%message "tmpl.sys.col.memo"%></th>
							<th class="f-w150"><%message "tmpl.sys.col.reg.time"%></th>
							<th><%message "tmpl.sys.col.login.ip"%></th>
							<th class="f-w150"><%message "tmpl.sys.col.login.time"%></th>
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
							<td data-id="{{areaId}}" data-dict="sys_area.id">&nbsp;</td>
							<td data-id="{{orgId}}" data-dict="sys_org.id">&nbsp;</td>
							<td>{{loginName}}</td>
							<td>{{name}}</td>
							<td>{{email}}</td>
							<td>{{mobile}}</td>
							<td>{{tel}}</td>
							<td>{{memo}}</td>
							<td>{{uyymdhms created}}</td>
							<td>{{loginIp}}</td>
							<td>{{uyymdhms loginTime}}</td>
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
			<!--#include file="/core/comm/formtree.html" param="{%.prefix%}" input="orgId" url="/sys/org/tree"--><!--#endinclude-->
			<script id="{%.prefix%}FormTemplate" type="text/x-handlebars-template">
				<form id="{%.prefix%}Form" method="post" go:action="{%apis%}/{%project%}/{%module%}/save">
					<br/>
					<!--#include file="/core/comm/alert.html" param="{%.prefix%}Form"--><!--#endinclude-->
					<input type="hidden" id="id" name="id" value="{{id}}" />
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.area"%></label>
						<div id="{%.prefix%}areaIdFormTree" class="form-inline" data-val="{{areaId}}">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.org"%></label>
						<div id="{%.prefix%}orgIdFormTree" class="form-inline" data-val="{{orgId}}">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.login.name"%></label>
						<div>
							<input class="form-control required account" id="loginName" name="loginName" value="{{loginName}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.name2"%></label>
						<div>
							<input class="form-control required realname" id="name" name="name" value="{{name}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.passwd"%></label>
						<div>
							<input type="password" class="form-control {{#if (eq state 'add')}} required {{/if}}passwd" id="passwd" name="passwd" placeholder="●●●●●●"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.email"%></label>
						<div>
							<input class="form-control required email" id="email" name="email" value="{{email}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.sys.col.mobile"%></label>
						<div>
							<input class="form-control mobile" id="mobile" name="mobile" value="{{mobile}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label"><%message "tmpl.sys.col.tel"%></label>
						<div>
							<input class="form-control tel" id="tel" name="tel" value="{{tel}}"/>
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.genre"%></label>
						<div id="radiosType" data-genre="sys_user.genre" data-name="genre" data-val="{{genre}}" class="well well-sm radio">
						</div>
					</div>
					<div class="form-group">
						<label class="control-label">*<%message "tmpl.sys.col.role"%></label>
						<div id="radiosRole" go:data-url="{%apis%}/sys/role/box?sOrdinalOA=10" data-name="roleIds" data-val="{{roleIds}}" class="well well-sm checkbox">
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
