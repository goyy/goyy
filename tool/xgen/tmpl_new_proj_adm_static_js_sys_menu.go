// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysMenu = `var ePostLoadForm=function(){
	$("#radiosTpes").radio();
	$("#radiosHidden").radio();
	$("#parentId").val($("#sParentId").val());
	$("#eForm").validate();
	if(state=="edit"){
		eparentIdFormTreeInit();
	}
}
`
