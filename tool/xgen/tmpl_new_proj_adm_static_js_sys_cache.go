// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysCache = `$(function(){
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
`
