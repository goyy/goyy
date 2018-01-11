$(function(){
	$("#eFormPasswd").validate({
		rules : {
			oldPasswd : {
				passwd : true,
				required : true
			},
			newPasswd : {
				passwd : true,
				required : true
			},
			okPasswd : {
				passwd : true,
				required : true,
				equalTo : "#newPasswd"
			}
		},
		submitHandler: function(form) {
			$("#eFormPasswd").disableSubmit();
			$("#eFormPasswd").ajaxSubmit({
				dataType : "json",
				success : function(result) {
					if (result.success) {
						$("#eFormAlert").showAlert("修改密码成功！");
						window.location.href = "/logout";
					}
					$("#eFormPasswd").enableSubmit();
				}
			});
			return false;
		}
	});
	
	//获取当前用户登录信息
	var url=apis+"/sys/user/principal";
	$.get(url,{},function(result){
		if (result.success) {
			$("#e-info").handlebars("e-info-template",result.data);
		}else{
			$("#e-info-alert").showAlert(result.message);
		}
	});
});