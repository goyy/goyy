$(function(){
	$("#sAreaIdTree").tree();
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
	$("#areaIdTree").tree();
	$("#parentIdTree").tree();
}

function selectedArea(id, name) {
	$("#sAreaId").val(id);
	$("#sAreaIdTreeNameTR").val(name);
}