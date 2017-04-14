$(function(){
	$("#eListSbtn").click(function(){
		$("#eListSform").disableSubmit();
		$("#eListSform").ajaxSubmit({
			dataType : "json",
			success : function(result) {
				if (result.success) {
					$("#data").text(result.data);
				} else {
					alert(result.message);
				}
				$("#eListSform").enableSubmit();
			}
		});
	});
});