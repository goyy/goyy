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
				<a class="navbar-brand" go:href="/home.html?{%ver%}">{{$.Project.Name}}</a>
			</div>
			<div class="navbar-collapse collapse">{{range $module := .Modules}}
				<ul class="nav navbar-nav">
					<li id="header_{{$module.Id}}" go:data-permissions="{{range $table := $.Tables}}{{if eq $table.Module.Id $module.Id}}{{if eq $table.Menu "true"}}{{if ne $table.Id "-"}}{{$table.Permissions}}{{end}}{{end}}{{end}}{{end}}" class="{%if eq .param '{{$module.Id}}'%}active {%end%}dropdown">
						<a href="#" class="dropdown-toggle" data-toggle="dropdown">{{$module.Name}}<b class="caret"></b></a>
						<ul class="dropdown-menu">{{range $table := $.Tables}}{{if eq $table.Module.Id $module.Id}}{{if eq $table.Menu "true"}}{{if eq $table.Id "-"}}
							<li data-permissions="{{$table.Permissions}}" class="divider"></li>{{else}}
							<li data-permissions="{{$table.Permissions}}"><a go:href="{{$table.Href}}?{%ver%}">{{$table.Name}}{{message "tmpl.menu.manage"}}</a></li>{{end}}{{end}}{{end}}{{end}}
						</ul>
					</li>
				</ul>{{end}}
				<ul class="nav navbar-nav navbar-right">
					<li><a id="loginName" go:href="/home.html?{%ver%}">&nbsp;</a></li>
					<li><a href="/logout">{{message "tmpl.menu.exit"}}</a></li>
				</ul>
			</div>
		</div>
	</div>
</div>`
