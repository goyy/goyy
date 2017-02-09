// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommNavtabsList = `<ul id="{%.param%}NavTabs" class="nav nav-tabs dn" >
	<li go:data-permissions="{%project%}:{%module%}:view" class="active">
		<a href="#{%.param%}ListContent" go:data-state="list" onclick="{%.param%}action(this,'')">{%title%}<%message "tmpl.core.comm.navtabs.list"%></a>
	</li>
</ul>
<script type="text/javascript">
	$(function(){
		$("#{%.param%}NavTabs").permission();
	});
</script>
`
