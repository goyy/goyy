// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSLogin = `$(function() {
	var winUrl=$.url();
	var redirect=winUrl.param("redirect");
	if($.isNotBlank(redirect)){
		$("#redirect").val(redirect);
	}
	$.fn.loginqtip=function(msg){
		$(this).qtip({
			position: {
				my: 'bottom left',
				at: 'top center'
			},
			style: {
				classes: 'qtip-red'
			},
			content: msg
		});
	}
	$("#eForm").validate({
		rules : {
			loginName : {
				required : true
			},
			passwd : {
				required : true
			},
			captcha : {
				required : true
			}
		},
		messages : {
			loginName : {
				required : "<%message "tmpl.valid.login.name"%>"
			},
			passwd : {
				required : "<%message "tmpl.valid.login.passwd"%>"
			},
			captcha : {
				required : "<%message "tmpl.valid.login.captcha"%>"
			}
		},
		submitHandler: function(form) {
			$(form).ajaxSubmit({
				dataType : "json",
				success : function(result) {
					if(result.success){
						window.location.href=result.data;
					}else{
						if(result.code=="-1"){
							$("#captcha").parent().addClass("has-error");
							$("#captcha").loginqtip(result.message);
						}else if(result.code=="-100"){
							$("#loginName").parent().addClass("has-error");
							$("#loginName").loginqtip(result.message);
							$("#passwd").parent().addClass("has-error");
							$("#passwd").loginqtip(result.message);
						}else{
							alert(result.message);
						}
						resetCaptcha();
					}
					$(form).enableSubmit();
				}
			});
			return false;
		},
		errorPlacement: function(error, element) {
			var tips = $(element);
			if (element.context.type == "checkbox") {
				tips = $(element).parent().parent().parent();
			}
			if (element.context.type == "hidden") {
				tips = $(element).parent();
			}
			tips.qtip({
				position: {
					my: 'bottom left',
					at: 'top center'
				},
				style: {
					classes: 'qtip-red'
				},
				content: $(error).text()
			});
		}
	});
});

// <%message "tmpl.note.captcha.reset"%>
function resetCaptcha(){
	$("#captchaImg").attr("src", "/captcha/build?random=" + Math.random());
}
`
