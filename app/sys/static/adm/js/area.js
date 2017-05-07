var ePostLoadForm = function(){
	$("#parentId").val($("#sParentId").val());
	$("#eForm").validate();
	if(state=="edit"){
		$("#parentIdTree").tree();
	}
}