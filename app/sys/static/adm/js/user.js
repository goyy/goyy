$(function(){
	$("#sAreaIdTree").tree();
	$("#sOrgIdTree").tree();
});


var ePostLoadForm=function(){
	$("#radiosType").radio();
	$("#radiosRole").checkbox();
	$("#eForm").validate({
		rules: {
			eGenre: {
				required: true
			},
			eAreaIdTname: {
				required: true
			},
			eOrgIdTname: {
				required: true
			}
		}
	});
	$("#areaIdTree").tree();
	$("#orgIdTree").tree();
}