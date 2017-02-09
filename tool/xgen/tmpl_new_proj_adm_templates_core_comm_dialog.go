// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommDialog = `<script type="text/javascript">
function msgAskDialog(callback, params, message) {
	if(typeof(message)!="undefined"&&message!=""){
		$("#_p_message_").html("<br/>" + message);
	}
	$("#msgAskDialog").dialog({
		resizable: false,
		modal: true,
		buttons: {
			"<%message "tmpl.core.comm.dialog.cancel"%>": function() {
				$(this).dialog("close");
			},
			"<%message "tmpl.core.comm.dialog.ok"%>": function() {
				callback(params);
				$(this).dialog("close");
			}
		}
	});
}
</script>
<div class="hidden">
	<div id="msgAskDialog" title="<%message "tmpl.core.comm.dialog.tips"%>">
		<p id="_p_message_"><br/>&nbsp;</p>
	</div>
</div>
`
