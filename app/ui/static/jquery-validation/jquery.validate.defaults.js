$.validator.setDefaults({
	ignore : "",
	submitHandler: function(form) {
		$(form).disableSubmit();
		// 新版ckeditor取消了内容自动同步到textarea的功能
		// 提交前需手动同步，若不手动同步提交时textarea中的内容为空
		for ( instance in CKEDITOR.instances ) {
			CKEDITOR.instances[instance].updateElement();
		}
		$(form).ajaxSubmit({
			dataType : "json",
			success : function(result) {
				if (result.success) {
					ePage();
					eTabsSwitch("#eListContent", "list");
					$("#eListAlert").showAlert("保存成功");
				} else {
					alert(result.message);
				}
				$(form).enableSubmit();
			}
		});
		return false;
	},
	highlight: function(element) {
		$(element).closest('div').addClass("has-error");
	},
	unhighlight: function(element) {
		$(element).closest('div').removeClass("has-error");
		$(element).qtip('destroy');
	},
	errorPlacement: function(error, element) {
		$(element.context.form).find("div[name='alert']").showAlert("输入有误，请先更正！");
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