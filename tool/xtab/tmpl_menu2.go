// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplMenu2 = `<script type="text/javascript">{{$ := .}}
	$(function(){
		var winUrl = $.url();
		var project=winUrl.param("project");
		if($.isNotBlank(project)){
			var headerMenu="#header_"+project;
			$(headerMenu).addClass("active");
		}
		$("#loginName").html($.getLoginName());
	});
</script>
<div id="header" class="navbar navbar-default navbar-fixed-top" role="navigation">
	<div class="container-fluid">
		<div class="container-fluid">
			<div class="navbar-header">
				<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" th:href="|/|" th:text="${#proj.title}">JAVAYY</a>
			</div>
			<div class="navbar-collapse collapse">{{range $module := .Modules}}
				<ul class="nav navbar-nav">
					<li id="header_{{$module.ID}}" th:permissions="{{range $table := $.Tables}}{{if eq $table.Module.ID $module.ID}}{{if eq $table.Menu "true"}}{{if ne $table.ID "-"}}{{$table.Permissions}},{{end}}{{end}}{{end}}{{end}}" class="{%if eq .param '{{$module.ID}}'%}active {%end%}dropdown">
						<a href="#" class="dropdown-toggle" data-toggle="dropdown">{{$module.Name}}<b class="caret"></b></a>
						<ul class="dropdown-menu">{{range $table := $.Tables}}{{if eq $table.Module.ID $module.ID}}{{if eq $table.Menu "true"}}{{if eq $table.ID "-"}}
							<li th:permissions="{{$table.Permissions}}" class="divider"></li>{{else}}
							<li th:permissions="{{$table.Permissions}}"><a th:href="{{$table.Href}}(ver=${#ver})">{{$table.Name}}{{message "tmpl.menu.manage"}}</a></li>{{end}}{{end}}{{end}}{{end}}
						</ul>
					</li>
				</ul>{{end}}
				<ul class="nav navbar-nav navbar-right">
					<li><a th:href="|/|" th:text="${#strings.abbreviate(#authentication.name,10)}">&nbsp;</a></li>
					<li><a href="/logout">{{message "tmpl.menu.exit"}}</a></li>
				</ul>
			</div>
		</div>
	</div>
</div>
`
