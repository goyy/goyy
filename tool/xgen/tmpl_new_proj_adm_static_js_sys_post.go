// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysPost = `var ePostLoadForm=function(){
	$("#genre").combo({
		genre: "sys_post.genre",
		val : genre,
		placeholder : "<%message "tmpl.form.combo.placeholder"%>",
		required : true,
		loadable : true
	});
	// <%message "tmpl.note.post.init"%>
	var setting = {
		check: {
			enable: true
		},
		data: {
			simpleData: {
				enable: true,
				rootPId: "root"
			}
		}
	};
	var data = $.parseJSON($("#menuIds").val());
	$.fn.zTree.init($("#treeMenu"), setting, data);
	
	// <%message "tmpl.note.form.valid"%>
	$("#eForm").validate({
		submitHandler: function(form) {
			var menuIds = "";
			var treeObj = $.fn.zTree.getZTreeObj("treeMenu");
			var nodes = treeObj.getCheckedNodes(true);
			if (nodes.length > 0) {
				for (var i in nodes) {
					if (i > 0) {
						menuIds += ","
					}
					menuIds += nodes[i].id
				}
			}
			$("#menuIds").val(menuIds);
			$("#eForm").disableSubmit();
			$("#eForm").ajaxSubmit({
				dataType : "json",
				success : function(result) {
					if (result.success) {
						ePage();
						eTabsSwitch("#eListContent");
						$("#eListAlert").showAlert("<%message "tmpl.form.submit.alert"%>");
					} else {
						alert(result.message);
					}
					$("#eForm").enableSubmit();
				}
			});
			return false;
		}
	});
}
`
