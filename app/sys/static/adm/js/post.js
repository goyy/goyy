var ePostLoadForm=function(){
	$("#genre").combo({
		genre: "sys_post.genre",
		val : genre,
		placeholder : "请选择",
		required : true,
		loadable : true
	});
	// 初始化岗位树结构数据
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
	
	// 表单校验
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