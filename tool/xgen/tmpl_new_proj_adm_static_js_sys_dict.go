// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSSysDict = `$(function(){
	$("#sGenreEQ").combo({url:apis+"/sys/dict/genres"});
});

var ePostLoadForm=function(){
	$("#eForm").validate();
}
`
