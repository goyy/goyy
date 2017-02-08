// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSCommForm = `(function($){
	// 单选框
	$.fn.radio = function(options){
		var $this = $(this);
		var args = $.extend({
			url: $this.data("url"),
			genre: $this.data("genre"),
			val : $this.data("val"),
			name : $this.data("name"),
			required : true,
			loadable : true
		}, options);
		if ($.isBlank(args.url)) {
			args.url = apis + "/sys/dict/box?sOrdinalOA=10&sGenreEQ=" + args.genre;
		}
		var onloaded = function(event, params, noSelectFirst) { // 加载数据方法
			if (!params && typeof(params)!="undefined" && params!=0) {
				params = {};
			}
			$.ajax({
				type : "get",
				dataType : "json",
				url : args.url,
				data : params,
				success : function(result) {
					if (result.data != null && typeof (result.data) != "undefined" && result.data.length > 0) {
						var isMatch = false;
						var first = "";
						for (var i in result.data) {
							var dataid = result.data[i].id;
							var dataname = result.data[i].name;
							if (i == 0) {
								first = dataid;
							}
							if ($.isNotBlank(args.val) && dataid == args.val) { // 判断data与val是否有匹配值
								isMatch = true
								$this.append('<label><input type="radio" checked="checked" name="' + args.name + '" value="' + dataid + '">' + dataname + '&nbsp;</input></label>');
								$this.trigger("change", [dataid]);
							} else {
								$this.append('<label><input type="radio" name="' + args.name + '" value="' + dataid + '">' + dataname + '&nbsp;</input></label>');
							}
						}
						if (args.required && !isMatch) { // 如果是必填项，且所有值都没有匹配上，那么就默认选中第一条记录
							if (noSelectFirst) { // 不自动选中第一行
							} else {
								$("input[name='" + args.name + "'][value='" + first + "']").attr("checked", true);
								$this.trigger("change", [first]);
							}
						}
						$this.trigger("onloaded"); // 加载完数据后派发onloaded事件
					}
				}
			});
		}
		$this.bind("loaded", onloaded);
		$this.bind("clear", function(){ // 清除数据
			$this.html("");
		});
		if (args.loadable) {
			$this.trigger("loaded");
		} else {
			$this.html("");
		}
	}
	
	// 复选框
	$.fn.checkbox = function(options){
		var $this = $(this);
		var args = $.extend({
			url: $this.data("url"),
			genre: $this.data("genre"),
			val : $this.data("val"),
			name : $this.data("name"),
			required : false,
			loadable : true
		}, options);
		if ($.isBlank(args.url)) {
			args.url = apis + "/sys/dict/box?sOrdinalOA=10&sGenreEQ=" + args.genre;
		}
		var onloaded = function(event, params, noSelectFirst) { // 加载数据方法
			if (!params && typeof(params)!="undefined" && params!=0) {
				params = {};
			}
			$.ajax({
				type : "get",
				dataType : "json",
				url : args.url,
				data : params,
				success : function(result) {
					if (result.data != null && typeof (result.data) != "undefined" && result.data.length > 0) {
						var isMatch = false;
						var first = "";
						for (var i in result.data) {
							var dataid = result.data[i].id;
							var dataname = result.data[i].name;
							if (i == 0) {
								first = dataid;
							}
							if ($.isBlank(args.val)) {
								args.val = ",";
							} else {
								args.val = args.val + ",";
							}
							if (args.val.indexOf(dataid+",") >= 0) { // 判断data与val是否有匹配值
								isMatch = true
								$this.append('<label><input type="checkbox" checked="checked" name="' + args.name + '" value="' + dataid + '">' + dataname + '&nbsp;</input></label>');
								$this.trigger("change", [dataid]);
							} else {
								$this.append('<label><input type="checkbox" name="' + args.name + '" value="' + dataid + '">' + dataname + '&nbsp;</input></label>');
							}
						}
						if (args.required && !isMatch) { // 如果是必填项，且所有值都没有匹配上，那么就默认选中第一条记录
							if (noSelectFirst) { // 不自动选中第一行
							} else {
								$("input[name='" + args.name + "'][value='" + first + "']").attr("checked", true);
								$this.trigger("change", [first]);
							}
						}
						$this.trigger("onloaded"); // 加载完数据后派发onloaded事件
					}
				}
			});
		}
		$this.bind("loaded", onloaded);
		$this.bind("clear", function(){ // 清除数据
			$this.html("");
		});
		if (args.loadable) {
			$this.trigger("loaded");
		} else {
			$this.html("");
		}
	}
	
	// 下拉框
	$.fn.combo = function(options){
		var $this = $(this);
		var args = $.extend({
			url: $this.data("url"),
			genre: $this.data("genre"),
			val : $this.data("val"),
			placeholder : "<%message "tmpl.form.combo.placeholder"%>",
			required : false,
			loadable : true
		}, options);
		if ($.isBlank(args.url)) {
			args.url = apis + "/sys/dict/box?sOrdinalOA=10&sGenreEQ=" + args.genre;
		}
		var onloaded = function(event, params, noSelectFirst) { // 加载数据方法
			if (!params && typeof(params)!="undefined" && params!=0) {
				params = {};
			}
			$.ajax({
				type : "get",
				dataType : "json",
				url : args.url,
				data : params,
				success : function(result) {
					if (result.data != null && typeof (result.data) != "undefined" && result.data.length > 0) {
						var isMatch = false;
						var first = "";
						for (var i in result.data) {
							var dataid = result.data[i].id;
							var dataname = result.data[i].name;
							if (i == 0) {
								first = dataid;
								if (args.required) {
									$this.html('');
								} else {
									$this.html('<option value="">' + args.placeholder + '</option>'); //如果是非必填，加入请选择选项
								}
							}
							if ($.isNotBlank(args.val) && dataid == args.val) { // 判断data与val是否有匹配值
								isMatch = true
								$this.append('<option selected="selected" value="' + dataid + '">' + dataname + '</option>');
								$this.trigger("change", [dataid]);
							} else {
								$this.append('<option value="' + dataid + '">' + dataname + '</option>');
							}
						}
						if (args.required && !isMatch) { // 如果是必填项，且所有值都没有匹配上，那么就默认选中第一条记录
							if (noSelectFirst) { // 不自动选中第一行
							} else {
								$this.val(first);
								$this.trigger("change", [first]);
							}
						}
						$this.trigger("onloaded"); // 加载完数据后派发onloaded事件
					}
				}
			});
		}
		$this.bind("loaded", onloaded);
		$this.bind("clear", function(){ // 清除数据
			if (args.required) {
				$this.html("");
			} else {
				$this.html('<option value="">' + args.placeholder + '</option>');
			}
		});
		if (args.loadable) {
			$this.trigger("loaded");
		} else {
			if (args.required) {
				$this.html("");
			} else {
				$this.html('<option value="">' + args.placeholder + '</option>');
			}
		}
	}
	
	// 下拉城市
	$.fn.area = function(options){
		var $this = $(this);
		var selectAreaId = $this.attr("id");
		var selectProvince = "#" + selectAreaId + " select[data-name='selectProvince']";
		var selectCity = "#" + selectAreaId + " select[data-name='selectCity']";
		var selectDistrict = "#" + selectAreaId + " select[data-name='selectDistrict']";
		// 初始化城市参数
		var args = $.extend({
			path : $this.data("path"),
			val : $this.data("val"),
			placeholderProvince : "<%message "tmpl.form.area.province"%>",
			placeholderCity : "<%message "tmpl.form.area.city"%>",
			placeholderDistrict : "<%message "tmpl.form.area.district"%>",
			three : true, // 是否是三级城市
			required : false, // 是否是必填
			loadable : false // 是否默认加载数据
		}, options);
		if (!args.three) { // 如果是二级城市
			selectDistrict = selectCity;
		}
		// 初始化一级城市参数
		var args1 = $.extend({
			val : "",
			placeholder : "<%message "tmpl.form.area.province"%>",
			required : false,
			loadable : false
		}, options);
		args1.val = "";
		if ($.isNotBlank(args1.placeholderProvince)) {
			args1.placeholder = args1.placeholderProvince;
		}
		args1.url = apis + "/sys/area/tree?sOrdinalOA=10&sParentIdEQ=root";
		$(selectProvince).combo(args1);
		// 初始化二级城市参数
		var args2 = $.extend({
			val : "",
			placeholder : "<%message "tmpl.form.area.city"%>",
			required : false,
			loadable : false
		}, options);
		args2.val = "";
		args2.url = apis + "/sys/area/tree?sOrdinalOA=10";
		if (args.three) { // 如果是三级城市
			if ($.isNotBlank(args2.placeholderCity)) {
				args2.placeholder = args2.placeholderCity;
			}
			$(selectCity).combo(args2);
		}
		// 初始化三级城市参数
		var args3 = $.extend({
			val : "",
			placeholder : "<%message "tmpl.form.area.district"%>",
			required : false,
			loadable : false
		}, options);
		if (args.three) { // 如果是三级城市
			if ($.isNotBlank(args3.placeholderDistrict)) {
				args3.placeholder = args3.placeholderDistrict;
			}
		} else {
			if ($.isNotBlank(args3.placeholderCity)) {
				args3.placeholder = args3.placeholderCity;
			}
		}
		args3.url = apis + "/sys/area/tree?sOrdinalOA=10";
		$(selectDistrict).combo(args3);
		
		var val = args.val;
		var initedValue = false;
		var level1AreaId = "nil";
		var level2AreaId = "nil";

		// 一级城市数据加载后执行操作：
		$(selectProvince).bind("onloaded", function() {
			if ($.isNotBlank(val)) { // 反显一级城市
				var url = apis + '/sys/area/parentid?id=' + val;
				if (args.three) { // 如果是三级城市
					url = apis + '/sys/area/parentpid?id=' + val;
				}
				$.ajax({
					async : false,
					type : 'get',
					url : url,
					dataType : 'json',
					success : function(result) {
						if (result.success) {
							setTimeout(function() { // 为了兼容IE6：使用setTimeout
								$(selectProvince).val(result.data);
								if (level1AreaId == "nil") { // 反显的一级城市赋值给level1AreaId：只执行一次，在初始化反显时
									level1AreaId = result.data;
								}
								$(selectProvince).trigger("change", [result.data]);  // 反显后：触发一级城市的onchange事件
							}, 1);
						}
					}
				});
			}
		});

		// 一级城市value改变时执行操作：
		$(selectProvince).change(function(event) {
			// 二级城市重新加载数据：
			var value = $(selectProvince).val();
			if ($.isNotBlank(value)) {
				$(selectCity).trigger("loaded", [{
					'sParentIdEQ' : value
				}, $.isBlank(val) ? false : (!initedValue || value == level1AreaId)]); // 触发二级城市的onloaded事件
			} else {
				// 清除二级城市数据
				$(selectCity).trigger("clear");
				$("#" + args.path).val("");
			}
			if (args.three) { // 如果是三级城市
				// 清除三级城市数据
				$(selectDistrict).trigger("clear");
				$("#" + args.path).val("");
			}
		});

		// 二级城市数据加载后执行操作：
		if (args.three) { // 如果是三级城市
			$(selectCity).bind("onloaded", function() {
				var value = $(selectProvince).val();
				if ((!initedValue || (value == level1AreaId)) && $.isNotBlank(val)) { // 反显二级城市
					$.ajax({
						async : false,
						type : 'get',
						url : apis + '/sys/area/parentid?id=' + val,
						dataType : 'json',
						contentType : "application/json; charset=UTF-8",
						success : function(result) {
							if (result.success) {
								setTimeout(function() { // 为了兼容IE6：使用setTimeout
									$(selectCity).val(result.data);
									if (level2AreaId == "nil") { // 反显的二级城市赋值给level2AreaId：只执行一次，在初始化反显时
										level2AreaId = result.data;
									}
									$(selectCity).trigger("change", [result.data]);  // 反显后：触发一级城市的onchange事件
									initedValue = true;
								}, 1);
							}
						}
					});
				} else {
					initedValue = true;
				}
			});
	
			// 二级城市value改变时执行操作：
			$(selectCity).change(function(event) {
				// 三级城市重新加载数据：
				var value = $(selectCity).val();
				if ($.isNotBlank(value)) {
					$(selectDistrict).trigger("loaded", [{
						'sParentIdEQ' : value
					}, false]);  // 触发三级城市的onloaded事件
				} else {
					// 清除三级城市数据
					$(selectDistrict).trigger("clear");
					$("#" + args.path).val("");
				}
			});
		}

		// 三级城市value改变时执行操作：
		$(selectDistrict).change(function(event) {
			var value = $(selectDistrict).val();
			$("#" + args.path).val(value);
		});

		// 一级城市加载数据：
		$(selectProvince).trigger("loaded", [{}, $.isNotBlank(val)]); // 触发一级城市的onloaded事件
	}
})(jQuery);
`
