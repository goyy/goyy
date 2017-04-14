$(function() {
	$.ajaxSetup({
		// 请求失败遇到异常触发
		error : function(xhr, status, e) {
			if (xhr.status == 0) {
				alert("服务器无法响应：请求未初始化！");
				return;
			}
			if ($.isBlank(xhr.responseText)) {
				alert("服务器无法响应：status=" + xhr.status + "！");
				return;
			}
			if (xhr.responseText.indexOf('"message":"') != -1) {
				var result = $.parseJSON(xhr.responseText);
				alert(result.message);
			} else {
				alert(xhr.responseText);
			}
		}
	});
});
