// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSHome = `$(function(){
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
						$("#eFormAlert").showAlert("<%message "tmpl.form.repwd.alert"%>");
						window.location.href = "/logout";
					}
					$("#eFormPasswd").enableSubmit();
				}
			});
			return false;
		}
	});
	
	// <%message "tmpl.note.user.get"%>
	var url=apis+"/sys/user/now/login/info";
	$.get(url,{},function(result){
		if (result.success) {
			$("#e-info").handlebars("e-info-template",result.data);
		}else{
			$("#e-info-alert").showAlert(result.message);
		}
	});
});
`
