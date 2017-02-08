// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysOrg = `$(function(){
	esAreaIdFormTreeInit();
});

var ePostLoadForm=function(){
	$("#parentId").val($("#sParentId").val());
	$("#eForm").validate({
		rules: {
			entityAreaIdTname: {
				required: true
			}
		}
	});
	eareaIdFormTreeInit();
	eparentIdFormTreeInit();
}

function selectedArea(id, name) {
	$("#sAreaId").val(id);
	$("#sAreaIdTreeNameTR").val(name);
}
`
