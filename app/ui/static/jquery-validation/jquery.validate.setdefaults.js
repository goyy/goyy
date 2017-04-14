$.validator.setDefaults({
	ignore : "",
	invalidHandler: function(form, validator) {
		//只报错第一个
		$.each(validator.invalid,function(key,value){
			$(".tipsError").text(value).show();
			return false;
		}); 
	},
	errorPlacement: function(error, element) {
	}
});