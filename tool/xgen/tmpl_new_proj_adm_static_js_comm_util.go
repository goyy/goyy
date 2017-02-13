// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSCommUtil = `/* jQuery util */
(function($) {
		
	/* 人民币金额转大写 */
	/* var out = $("#inputId").amountToChinese() */
	$.fn.amountToChinese = function() {
		var cnNums = new Array("零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"); // 汉字的数字
		var cnIntRadice = new Array("", "拾", "佰", "仟"); // 基本单位
		var cnIntUnits = new Array("", "万", "亿", "兆"); // 对应整数部分扩展单位
		var cnDecUnits = new Array("角", "分", "毫", "厘"); // 对应小数部分单位
		var cnInteger = "整"; // 整数金额时后面跟的字符
		var cnIntLast = "元"; // 整型完以后的单位
		var maxNum = 999999999999999.9999; // 最大处理的数字

		var IntegerNum; // 金额整数部分
		var DecimalNum; // 金额小数部分
		var ChineseStr = ""; // 输出的中文金额字符串
		var parts; // 分离金额后用的数组，预定义

		var $this = $(this);
		var money = $this.val();
		if (money == "") {
			return "";
		}

		money = parseFloat(money);
		// alert(money);
		if (money >= maxNum) {
			$.alert('<%message "tmpl.form.amount.limit"%>');
			return "";
		}
		if (money == 0) {
			ChineseStr = cnNums[0] + cnIntLast + cnInteger;
			// document.getElementById("show").value=ChineseStr;
			return ChineseStr;
		}
		money = money.toString(); // 转换为字符串
		if (money.indexOf(".") == -1) {
			IntegerNum = money;
			DecimalNum = '';
		} else {
			parts = money.split(".");
			IntegerNum = parts[0];
			DecimalNum = parts[1].substr(0, 4);
		}
		if (parseInt(IntegerNum, 10) > 0) {// 获取整型部分转换
			zeroCount = 0;
			IntLen = IntegerNum.length;
			for (i = 0; i < IntLen; i++) {
				n = IntegerNum.substr(i, 1);
				p = IntLen - i - 1;
				q = p / 4;
				m = p % 4;
				if (n == "0") {
					zeroCount++;
				} else {
					if (zeroCount > 0) {
						ChineseStr += cnNums[0];
					}
					zeroCount = 0; // 归零
					ChineseStr += cnNums[parseInt(n)] + cnIntRadice[m];
				}
				if (m == 0 && zeroCount < 4) {
					ChineseStr += cnIntUnits[q];
				}
			}
			ChineseStr += cnIntLast;
			// 整型部分处理完毕
		}
		if (DecimalNum != '') {// 小数部分
			decLen = DecimalNum.length;
			for (i = 0; i < decLen; i++) {
				n = DecimalNum.substr(i, 1);
				if (n != '0') {
					ChineseStr += cnNums[Number(n)] + cnDecUnits[i];
				}
			}
		}
		if (ChineseStr == '') {
			ChineseStr += cnNums[0] + cnIntLast + cnInteger;
		} else if (DecimalNum == '') {
			ChineseStr += cnInteger;
		}
		return ChineseStr;
	}

	/* 启用提交按钮 */
	$.fn.enableSubmit = function() {
		$(this).find("input[type='submit']").attr('disabled', false);
	}

	/* 禁用提交按钮 */
	$.fn.disableSubmit = function() {
		$(this).find("input[type='submit']").attr('disabled', true);
	}

	/* 显示错误提示 */
	$.fn.showError = function(err) {
		$(this).find(".tipsError").text(err).show();
	}

	/* 隐藏错误提示 */
	$.fn.hideError = function(msg) {
		$(this).find(".tipsError").text(err).hide();
	}
	
	//提示框显示
	$.fn.showAlert=function(msg){
		var box=$(this);
		box.removeClass("hidden");
		box.addClass("show");
		box.children("strong").text(msg);
		box.children("button").unbind("click");
		box.children("button").click(function(){
			box.hideAlert();
		});
	}
	
	//提示框隐藏
	$.fn.hideAlert=function(){
		var box=$(this);
		box.removeClass("show");
		box.addClass("hidden");
		box.children("strong").text("");
	}

	/* 60秒倒计时按钮 */
	$.fn.setInterval = function(text) {
		var $this = $(this);
		$this.attr("disabled", "disabled");
		var limit = 60;
		var time = limit;
		var interval = setInterval(function() {
			if (time > 1) {
				time--;
				$this.val(time + "<%message "tmpl.form.interval.time"%>");
			} else {
				time = limit;
				clearInterval(interval);
				$this.val(text);
				$this.removeAttr("disabled");
			}
		}, 1000);
	}
	
	// handlebars 模版
	$.fn.handlebars=function(templateId, templateData){
		var source = $("#" + templateId).html();
		var template = Handlebars.compile(source);
		var content=template(templateData);
		$(this).hide();
		$(this).html(content);
		$(this).permission(); // 这个方法里会控制显示：不用调用$(this).show();
	};
	
	// 权限判断
	$.fn.permission=function(){
		var n = $.cookie('GSESSIONN');
		if ($.isNotBlank(n)) {
			var pn = "";
			for (i=0;i<=n;i++) {
				pn += $.cookie('GSESSION'+i);
			}
			var ps = Base64.decode(pn);
			var dp = $(this).find("[data-permissions]");
			for (i=0;i<dp.length;i++) {
				$(dp[i]).hide();
				var p = $(dp[i]).data("permissions");
				var pp = p.split(",");
				for (y=0;y<pp.length;y++) {
					if ($.isBlank(pp[y])) {
						continue;
					}
					if (ps.indexOf(pp[y]) >= 0) {
						$(dp[i]).show();
						break;
					}
				}
			}
		} else {
			var dp = $(this).find("[data-permissions]");
			for (i=0;i<dp.length;i++) {
				$(dp[i]).hide();
			}
		}
		if (!$(this).is("[data-permissions]")) {
			// 不能使用$(this).show();否则bootstrap的nav-tabs的函数tab("show")会失效
			$(this).css("display", "");
		}
		
		return true;
	};
	
	// 字典显示
	$.fn.dict=function(options){
		var $this = $(this);
		var args = $.extend({
			genre: "dict",
			url: apis + "/sys/dict/list"
		}, options);
		if ($.isBlank(args.genre)) {
			args.genre = "dict";
		}
		if ($.isBlank(args.url)) {
			args.url = apis + "/sys/dict/list";
		}
		var dd = $this.find("[data-" + args.genre + "]");
		if (dd.length > 0) {
			var params = "[";
			for (i=0;i<dd.length;i++) {
				var id = $(dd[i]).data("id");
				var genre = $(dd[i]).data(args.genre);
				var descr = $(dd[i]).data("descr");
				if ($.isBlank(id) || $.isBlank(genre)) {
					continue;
				}
				var dict = '{"mkey":"'+id+'","genre":"'+genre+'"}';
				if ($.isNotBlank(descr)) {
					dict = '{"mkey":"'+id+'","genre":"'+genre+'","descr":"'+descr+'"}';
				}
				if (i == 0) {
					params = params + dict;
				} else {
					params = params + ',' + dict;
				}
			}
			params = params + "]";
			$.ajax({
				type : "post",
				dataType : "json",
				url : args.url,
				data : {"params":params},
				success : function(result) {
					if (result.data != null && typeof (result.data) != "undefined" && result.data.length > 0) {
						for (i=0;i<result.data.length;i++) {
							if (result.data[i] != null) {
								var did = result.data[i].mkey;
								var dgenre = result.data[i].genre;
								var dname = result.data[i].mval;
								if ($.isNotBlank(dname)) {
									$("[data-id='"+did+"'][data-"+args.genre+"='"+dgenre+"']").html(dname);
								}
							}
						}
					}
				}
			});
		}
	};
	
	// 导出excel
	$.fn.export=function(project,module){
		var options={
			url:apis+"/"+project+"/"+module+"/export",
			type:"get",
			success:function(result,statusText){
				if(result.success){
					window.open(developers+"/export/excel/"+result.data);
				}else{
					alert(result.message);
				}
			},
			error:function(result){
				console.log(result);
				alert('<%message "tmpl.form.util.export.err"%>');
			},
			dataType:"json"
		};
		$(this).ajaxSubmit(options);
		return false;
	};
	
	// 判断是否登录
	$.isLogin=function(){
		var user = $.cookie('GSESSIONUSER');
		if ($.isNotBlank(user)) {
			return true;
		}
		return false;
	};
	
	// 获取登录名
	$.getLoginName=function(){
		var user = $.cookie('GSESSIONUSER');
		if ($.isNotBlank(user)) {
			return user;
		}
		return "";
	};

	/* unix时间格式化 -> yyyy-MM-dd */
	$.uyymd = function(time) {
		time = "" + time;
		if(time.length == 10){
			time = time + "000";
		}
		var date = new Date(time - 0);
		return $.format.date(date, "yyyy-MM-dd");
	}

	/* unix时间格式化 -> yyyy-MM-dd HH:mm:ss */
	$.uyymdhms = function(time) {
		time = "" + time;
		if(time.length == 10){
			time = time + "000";
		}
		var date = new Date(time - 0);
		return $.format.date(date, "yyyy-MM-dd HH:mm:ss");
	}

	/* unix时间格式化 -> yyyy-MM-dd HH:mm */
	$.uyymdhm = function(time) {
		time = "" + time;
		if(time.length == 10){
			time = time + "000";
		}
		var date = new Date(time - 0);
		return $.format.date(date, "yyyy-MM-dd HH:mm");
	}

	/* 判断是否为空 */
	/* var out = $.isBlank("  ") */
	$.isBlank = function(val) {
		if (val != null && typeof (val) != "undefined" && $.trim(val).length > 0) {
			return false;
		}
		return true;
	}

	/* 判断是否不为空 */
	/* var out = $.isNotBlank("  ") */
	$.isNotBlank = function(val) {
		if (val != null && typeof (val) != "undefined" && $.trim(val).length > 0) {
			return true;
		}
		return false;
	}

	/*
	 * 根据长网址获取短网址. $.getShortUrl("http://www.yinongdai.com/cms/article/aboutUs",
	 * function(data) { alert(data[0].url_short); });
	 */
	$.getShortUrl = function(longUrl, callback) {
		var url = "http://api.t.sina.com.cn/short_url/shorten.json?source=1681459862&url_long=" + longUrl;
		$.ajax({
			async : false,
			type : 'get',
			url : url,
			dataType : 'jsonp',
			jsonpCallback : "getShortUrlCallback",
			success : callback
		});
	}

	/**
	 * 根据IP获取相关信息包括城市、提供商等 $.getIpInfo("114.113.223.34", function(){ var
	 * resultData = $.getIpInfoBack(remote_ip_info); //中国北京北京 alert(resultData);
	 * });
	 * 
	 */
	$.getIpInfo = function(ip, callback) {
		var url = 'http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=js&ip=' + ip;
		$.getScript(url, callback);
	}
	$.getIpInfoBack = function(ipInfo, split) {
		if (split == undefined) {
			split = "-";
		}
		var resultData = (ipInfo.country + split + ipInfo.province + split + ipInfo.city);
		return resultData;
	}
	
	/* 截取字符串,当value.length>end,返回截取结果+"..." */
	/* var out=$.abbr("字符串",20); */
	$.abbr=function(value,end){
		var len=value.length;
		value=value.substring(0,end);
		if(len>end){
			value=value+"...";
		}
		return value;
	};

	/* 验证身份证 */
	/* var out = $.checkIdCard("431381198809122734") */
	$.checkIdCard = function(idCard) {
		idCard = idCard.toString();
		// var Errors=new
		// Array("验证通过!","身份证号码位数不对!","身份证号码出生日期超出范围或含有非法字符!","身份证号码校验错误!","身份证地区非法!");
		var Errors = new Array(true, false, false, false, false);
		var area = {
			11 : "北京",
			12 : "天津",
			13 : "河北",
			14 : "山西",
			15 : "内蒙古",
			21 : "辽宁",
			22 : "吉林",
			23 : "黑龙江",
			31 : "上海",
			32 : "江苏",
			33 : "浙江",
			34 : "安徽",
			35 : "福建",
			36 : "江西",
			37 : "山东",
			41 : "河南",
			42 : "湖北",
			43 : "湖南",
			44 : "广东",
			45 : "广西",
			46 : "海南",
			50 : "重庆",
			51 : "四川",
			52 : "贵州",
			53 : "云南",
			54 : "西藏",
			61 : "陕西",
			62 : "甘肃",
			63 : "青海",
			64 : "宁夏",
			65 : "新疆",
			71 : "台湾",
			81 : "香港",
			82 : "澳门",
			91 : "国外"
		}
		var idCard, Y, JYM;
		var S, M;
		var idCard_array = new Array();
		idCard_array = idCard.split("");
		// 地区检验
		if (area[parseInt(idCard.substr(0, 2))] == null)
			return Errors[4];
		// 身份号码位数及格式检验
		switch (idCard.length) {
		case 15:
			if ((parseInt(idCard.substr(6, 2)) + 1900) % 4 == 0 || ((parseInt(idCard.substr(6, 2)) + 1900) % 100 == 0 && (parseInt(idCard.substr(6, 2)) + 1900) % 4 == 0)) {
				ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))[0-9]{3}$/;// 测试出生日期的合法性
			} else {
				ereg = /^[1-9][0-9]{5}[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))[0-9]{3}$/;// 测试出生日期的合法性
			}
			if (ereg.test(idCard))
				return Errors[0];
			else
				return Errors[2];
			break;
		case 18:
			// 18 位身份号码检测
			// 出生日期的合法性检查
			// 闰年月日:((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))
			// 平年月日:((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))
			if (parseInt(idCard.substr(6, 4)) % 4 == 0 || (parseInt(idCard.substr(6, 4)) % 100 == 0 && parseInt(idCard.substr(6, 4)) % 4 == 0)) {
				ereg = /^[1-9][0-9]{5}19[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|[1-2][0-9]))[0-9]{3}[0-9Xx]$/;// 闰年出生日期的合法性正则表达式
			} else {
				ereg = /^[1-9][0-9]{5}19[0-9]{2}((01|03|05|07|08|10|12)(0[1-9]|[1-2][0-9]|3[0-1])|(04|06|09|11)(0[1-9]|[1-2][0-9]|30)|02(0[1-9]|1[0-9]|2[0-8]))[0-9]{3}[0-9Xx]$/;// 平年出生日期的合法性正则表达式
			}
			if (ereg.test(idCard)) {// 测试出生日期的合法性
				// 计算校验位
				S = (parseInt(idCard_array[0]) + parseInt(idCard_array[10])) * 7 + (parseInt(idCard_array[1]) + parseInt(idCard_array[11])) * 9 + (parseInt(idCard_array[2]) + parseInt(idCard_array[12])) * 10 + (parseInt(idCard_array[3]) + parseInt(idCard_array[13])) * 5 + (parseInt(idCard_array[4]) + parseInt(idCard_array[14])) * 8 + (parseInt(idCard_array[5]) + parseInt(idCard_array[15])) * 4 + (parseInt(idCard_array[6]) + parseInt(idCard_array[16])) * 2 + parseInt(idCard_array[7]) * 1 + parseInt(idCard_array[8]) * 6 + parseInt(idCard_array[9]) * 3;
				Y = S % 11;
				M = "F";
				JYM = "10X98765432";
				M = JYM.substr(Y, 1);// 判断校验位
				if (M == idCard_array[17])
					return Errors[0]; // 检测ID的校验位
				else
					return Errors[3];
			} else
				return Errors[2];
			break;
		default:
			return Errors[1];
			break;
		}
	}
})(jQuery);
`
