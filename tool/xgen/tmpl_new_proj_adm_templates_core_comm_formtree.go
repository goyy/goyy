// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommFormtree = `<script type="text/javascript">
	function {%.param%}{%.input%}FormTreeSetValue(id, name) {
		var {%.param%}treeform=$("#{%.param%}{%.input%}FormTree");
		{%.param%}treeform.find("input[name='{%.input%}']").val(id);
		{%.param%}treeform.find("input[name='{%.input%}Tname']").val(name);
	}
	function {%.param%}{%.input%}FormTreeOpen() {
		$("#{%.param%}{%.input%}Dialog").dialog({
			resizable: false,
			modal: true,
			height: 500,
			buttons: {
				"<%message "tmpl.core.comm.formtree.cancel"%>": function() {
					$(this).dialog("close");
				},
				"<%message "tmpl.core.comm.formtree.ok"%>": function() {
					var treeObj = $.fn.zTree.getZTreeObj("{%.param%}{%.input%}Tree");
					var nodes = treeObj.getSelectedNodes();
					if (nodes.length == 1) {
						{%.param%}{%.input%}FormTreeSetValue(nodes[0].id, nodes[0].name);
					}
					$(this).dialog("close");
				}
			}
		});
	}
	function {%.param%}{%.input%}Init(data) {
		var setting = {
			data: {
				simpleData: {
					enable: true,
					rootPId: "root"
				}
			}
		};
		$.fn.zTree.init($("#{%.param%}{%.input%}Tree"), setting, data);
		var {%.param%}treeform=$("#{%.param%}{%.input%}FormTree");
		var value={%.param%}treeform.find("input[name='{%.input%}']").val();
		if (data.length > 0 && value.length > 0) {
			for (var i in data) {
				if (data[i].id == value) {
					{%.param%}treeform.find("input[name='{%.input%}Tname']").val(data[i].name);
				}
			}
		}
	}
	
	//页面加载后或模版加载后调用的函数入口 初始化
	function {%.param%}{%.input%}FormTreeInit(){
		var {%.param%}treeform=$("#{%.param%}{%.input%}FormTree");
		var hidden='<input type="hidden" name="{%.input%}" value="'+{%.param%}treeform.data("val")+'"/>';
		var thname='<input class="form-control" readonly="readonly" name="{%.input%}Tname" value=""/>';
		var selectBtn='<button id="{%.param%}{%.input%}SelectTreeBtn" type="button" class="btn btn-default" onclick="{%.param%}{%.input%}FormTreeOpen();" ><%message "tmpl.core.comm.formtree.select"%></button>';
		{%.param%}treeform.append(hidden);
		{%.param%}treeform.append(thname);
		{%.param%}treeform.append(selectBtn);
		$.ajax({
			type: "get",
			dataType: "json",
			url: apis+"{%.url%}",
			success: function(result) {
				{%.param%}{%.input%}Init(result.data);
			}
		});
	}
	
</script>
<div class="hidden">
	<div id="{%.param%}{%.input%}Dialog" title="<%message "tmpl.core.comm.formtree.title"%>">
		<ul id="{%.param%}{%.input%}Tree" class="ztree"></ul>
	</div>
</div>
`
