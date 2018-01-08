// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplJsMain = `
function ePostLoadForm() {<%range $i, $e := .Entities%><%range $f := $e.Fields%><%if notblank $f.Dict%>
	var <%$f.Name%>=$("#<%$f.Name%>").val();
	$("#<%$f.Name%>").combo({
		genre: "<%$f.Dict%>",
		val : <%$f.Name%>,
		placeholder : "<%message "tmpl.form.combo.placeholder"%>",
		required : true,
		loadable : true
	});<%end%><%end%><%end%>
	$("#eForm").validate();
}
`
