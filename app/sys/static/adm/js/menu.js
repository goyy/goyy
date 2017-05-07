var ePostLoadForm=function(){
	$("#radiosTpes").radio();
	$("#radiosHidden").radio();
	$("#parentId").val($("#sParentId").val());
	$("#eForm").validate();
	if(state=="edit"){
		$("#parentIdTree").tree();
	}
}