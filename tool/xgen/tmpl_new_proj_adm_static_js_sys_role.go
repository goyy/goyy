// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysRole = `$(function(){
	$("#sGenreEQ").combo({
		genre: "sys_role.genre",
		val : "",
		placeholder : "<%message "tmpl.form.combo.placeholder"%>",
		required : false,
		loadable : true
	});
})

var ePostLoadForm=function(){
	var genre=$("#genre").attr("value");
	$("#genre").combo({
		genre: "sys_role.genre",
		val : genre,
		placeholder : "<%message "tmpl.form.combo.placeholder"%>",
		required : true,
		loadable : true
	});
	var classify=$("#classify").attr("value");
	$("#classify").combo({
		genre: "sys_role.classify",
		val : classify,
		placeholder : "<%message "tmpl.form.combo.placeholder"%>",
		required : false,
		loadable : true
	});
	// <%message "tmpl.note.role.init"%>
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
	var data = $.parseJSON($("#postIds").val());
	$.fn.zTree.init($("#treeMenu"), setting, data);
	// <%message "tmpl.note.form.valid"%>
	$("#eForm").validate({
		submitHandler: function(form) {
			var postIds = "";
			var treeObj = $.fn.zTree.getZTreeObj("treeMenu");
			var nodes = treeObj.getCheckedNodes(true);
			if (nodes.length > 0) {
				for (var i in nodes) {
					if (i > 0) {
						postIds += ","
					}
					postIds += nodes[i].id
				}
			}
			$("#postIds").val(postIds);
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
