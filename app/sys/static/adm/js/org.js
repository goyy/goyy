$(function(){
	esAreaIdFormTreeInit();
});

var ePostLoadForm=function(){
	$("#parentId").val($("#sParentId").val());
	$("#eForm").validate({
		rules: {
			entityAreaIdTname: {
				required: true
			}
		}
	});
	eareaIdFormTreeInit();
	eparentIdFormTreeInit();
}

function selectedArea(id, name) {
	$("#sAreaId").val(id);
	$("#sAreaIdTreeNameTR").val(name);
}