// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommDisable = `<script type="text/javascript">
	function disable(obj, id, callback) {
		$("#disableAskDialog").dialog({
			resizable: false,
			modal: true,
			buttons: {
				"<%message "tmpl.core.comm.disable.cancel"%>": function() {
					$(this).dialog("close");
				},
				"<%message "tmpl.core.comm.disable.ok"%>": function() {
					{%.param%}Action(obj, id, callback);
					$(this).dialog("close");
				}
			}
		});
	}
</script>
<div class="hidden">
	<div id="disableAskDialog" title="<%message "tmpl.core.comm.disable.tips"%>">
		<p><br/><%message "tmpl.core.comm.disable.ask"%></p>
	</div>
</div>
`
