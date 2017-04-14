$(function(){
	$("#sGenreEQ").combo({
		genre: "sys_role.genre",
		val : "",
		placeholder : "请选择",
		required : false,
		loadable : true
	});
})

var ePostLoadForm=function(){
	var genre=$("#genre").attr("value");
	$("#genre").combo({
		genre: "sys_role.genre",
		val : genre,
		placeholder : "请选择",
		required : true,
		loadable : true
	});
	var classify=$("#classify").attr("value");
	$("#classify").combo({
		genre: "sys_role.classify",
		val : classify,
		placeholder : "请选择",
		required : false,
		loadable : true
	});
	// 初始化角色树结构数据
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
	// 表单校验
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
						$("#eListAlert").showAlert("保存成功");
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