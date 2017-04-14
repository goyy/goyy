$(function(){
	esAreaIdFormTreeInit();
	esOrgIdFormTreeInit();
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
	eareaIdFormTreeInit();
	eorgIdFormTreeInit();
}