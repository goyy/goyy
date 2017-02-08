// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysUser = `$(function(){
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
`
