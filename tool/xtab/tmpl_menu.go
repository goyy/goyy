// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplMenu = `<script type="text/javascript">{{$ := .}}
	$(function(){
		$("#header").permission();
		var winUrl = $.url();
		var project=winUrl.param("project");
		if($.isNotBlank(project)){
			var headerMenu="#header_"+project;
			$(headerMenu).addClass("active");
		}
		$("#loginName").html($.getLoginName());
	});
</script>
<div id="header" class="navbar navbar-default navbar-fixed-top dn" role="navigation">
	<div class="container-fluid">
		<div class="container-fluid">
			<div class="navbar-header">
				<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" href="/sys/user/info.html">{{$.Project.Name}}</a>
			</div>
			<div class="navbar-collapse collapse">
				<ul class="nav navbar-nav">
					<li id="header_sys" data-permissions="sys:area:view,sys:org:view,sys:user:view,sys:menu:view,sys:post:view,sys:role:view,sys:dict:view,sys:log:view,sys:cache:view" class="{%if eq .param 'sys'%}active {%end%}dropdown">
						<a href="#" class="dropdown-toggle" data-toggle="dropdown">系统管理<b class="caret"></b></a>
						<ul class="dropdown-menu">
							<li data-permissions="sys:area:view"><a href="/sys/area/area.html">区域管理</a></li>
							<li data-permissions="sys:org:view"><a href="/sys/org/org.html">机构管理</a></li>
							<li data-permissions="sys:user:view"><a href="/sys/user/user.html">用户管理</a></li>
							<li data-permissions="sys:menu:view,sys:post:view,sys:role:view" class="divider"></li>
							<li data-permissions="sys:menu:view"><a href="/sys/menu/menu.html">菜单管理</a></li>
							<li data-permissions="sys:post:view"><a href="/sys/post/post.html">岗位管理</a></li>
							<li data-permissions="sys:role:view"><a href="/sys/role/role.html">角色管理</a></li>
							<li data-permissions="sys:dict:view" class="divider"></li>
							<li data-permissions="sys:dict:view"><a href="/sys/dict/dict.html">字典管理</a></li>
							<li data-permissions="sys:log:view,sys:cache:view" class="divider"></li>
							<li data-permissions="sys:log:view"><a href="/sys/log/log.html">日志查询</a></li>
							<li data-permissions="sys:cache:view"><a href="/sys/cache/cache.html">缓存查询</a></li>
						</ul>
					</li>
				</ul>
				<ul class="nav navbar-nav navbar-right">
					<li><a id="loginName" href="/sys/user/info.html">&nbsp;</a></li>
					<li><a href="/logout">退出</a></li>
				</ul>
			</div>
		</div>
	</div>
</div>`
