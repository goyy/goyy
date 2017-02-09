// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommNavtabs = `<ul id="{%.param%}NavTabs" class="nav nav-tabs dn" >
	<li go:data-permissions="{%project%}:{%module%}:view" class="active">
		<a href="#{%.param%}ListContent" go:data-state="list" onclick="{%.param%}Action(this,'')">{%title%}<%message "tmpl.core.comm.navtabs.list"%></a>
	</li>
	<li go:data-permissions="{%project%}:{%module%}:add">
		<a href="#{%.param%}FormContent" go:data-state="add" go:data-url="{%apis%}/{%project%}/{%module%}/add" data-template="{%.param%}FormTemplate" 
			onclick="{%.param%}Action(this,'',{%.param%}PostLoadForm)" >{%title%}<span id="{%.param%}TabTitle"><%message "tmpl.core.comm.navtabs.add"%></span></a>
	</li>
</ul>
<script type="text/javascript">
	$(function(){
		$("#{%.param%}NavTabs").permission();
	});
</script>
`
