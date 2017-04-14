/*==================时间相关=====================*/

// unix时间格式化 -> yyyy-MM-dd
Handlebars.registerHelper("uyymd", function(time) {
	if ($.isBlank(time) || time <= 0) {
		return "";
	}
	time = "" + time;
	if (time.length == 10) {
		time = time + "000";
	}
	var date = new Date(time - 0);
	return $.format.date(date, "yyyy-MM-dd");
});

// unix时间格式化 -> yyyy-MM-dd HH:mm:ss
Handlebars.registerHelper("uyymdhms", function(time) {
	if ($.isBlank(time) || time <= 0) {
		return "";
	}
	time = "" + time;
	if (time.length == 10) {
		time = time + "000";
	}
	var date = new Date(time - 0);
	return $.format.date(date, "yyyy-MM-dd HH:mm:ss");
});

// unix时间格式化 -> yyyy-MM-dd HH:mm
Handlebars.registerHelper("uyymdhm", function(time) {
	if ($.isBlank(time) || time <= 0) {
		return "";
	}
	time = "" + time;
	if (time.length == 10) {
		time = time + "000";
	}
	var date = new Date(time - 0);
	return $.format.date(date, "yyyy-MM-dd HH:mm");
});

/* ==================与或判断===================== */

// 与判断
Handlebars.registerHelper("and", function(src, dst) {
	return src && dst;
});

// 或判断
Handlebars.registerHelper("or", function(src, dst) {
	return src || dst;
});

/* ==================相等大小比较===================== */

// 判断是否相等
Handlebars.registerHelper("eq", function(src, dst) {
	return src == dst;
});

// 判断是否不相等
Handlebars.registerHelper("ne", function(src, dst) {
	return src != dst;
});

// 判断是否小于
Handlebars.registerHelper("lt", function(src, dst) {
	return src < dst;
});

// 判断是否小于等于
Handlebars.registerHelper("le", function(src, dst) {
	return src <= dst;
});

// 判断是否大于
Handlebars.registerHelper("gt", function(src, dst) {
	return src > dst;
});

// 判断是否大于等于
Handlebars.registerHelper("ge", function(src, dst) {
	return src >= dst;
});

/* ==================字符串相关===================== */

// 字符串md5加密
Handlebars.registerHelper("md5", function(val) {
	var hash = md5(val);
	return hash;
});

// 判断是否包含
Handlebars.registerHelper("contains", function(src, dst) {
	return src.indexOf(dst) != -1;
});

// 将两个或多个字符的文本组合起来，返回一个新的字符串
Handlebars.registerHelper("concat", function(src) {
	for (var i = 1; i < arguments.length; i++) {
		if (typeof (arguments[i]) == "string") {
			src = src.concat(arguments[i]);
		}
	}
	return src;
});

// 返回字符串中一个子串第一处出现的索引（从左到右搜索）。如果没有匹配项，返回 -1
Handlebars.registerHelper("indexOf", function(src, dst) {
	return src.indexOf(dst);
});

// 返回字符串中一个子串最后一处出现的索引（从右到左搜索），如果没有匹配项，返回 -1
Handlebars.registerHelper("lastIndexOf", function(src, dst) {
	return src.indexOf(dst);
});

// 截取字符串,当src.length>end,返回截取结果+"..."
Handlebars.registerHelper("abbr", function(src, end) {
	var len = src.length;
	var value = src.substring(0, end);
	if (len > end) {
		value = value + "...";
	}
	return value;
});

// 得到左边的字符串
Handlebars.registerHelper("left", function(src, len) {
	if (isNaN(len) || len == null) {
		len = src.length;
	} else {
		if (parseInt(len) < 0 || parseInt(len) > src.length) {
			len = src.length;
		}
	}
	return src.substr(0, len);
});

// 得到左边的字符串
Handlebars.registerHelper("right", function(src, len) {
	if (isNaN(len) || len == null) {
		len = src.length;
	} else {
		if (parseInt(len) < 0 || parseInt(len) > src.length) {
			len = src.length;
		}
	}
	return src.substring(src.length - len, src.length);
});

// 返回字符串的一个子串，传入参数是起始位置和结束位置
Handlebars.registerHelper("substring", function(src, begin, end) {
	return src.substring(begin, end);
});

// 返回字符串的一个子串，传入参数是起始位置和长度
Handlebars.registerHelper("substr", function(src, begin, len) {
	return src.substr(begin, len);
});

// 返回字符串的一个子串，传入参数是起始位置和长度
Handlebars.registerHelper("replace", function(src, reg, dst) {
	return src.replace(reg, dst);
});

// 返回字符串的长度，所谓字符串的长度是指其包含的字符的个数
Handlebars.registerHelper("len", function(src) {
	return src.length;
});

// 去除前后空格
Handlebars.registerHelper("trim", function(src) {
	return src.replace(/(^\s*)|(\s*$)/g, "");
});

// 去除左边的空格
Handlebars.registerHelper("ltrim", function(src) {
	return src.replace(/(^\s*)/g, "");
});

// 去除右边的空格
Handlebars.registerHelper("rtrim", function(src) {
	return src.replace(/(\s*$)/g, "");
});

// 将整个字符串转成小写字母
Handlebars.registerHelper("toLower", function(src) {
	return src.toLowerCase();
});

// 将整个字符串转成大写字母
Handlebars.registerHelper("toUpper", function(src) {
	return src.toUpperCase();
});

// 判断是否为空
Handlebars.registerHelper("isBlank", function(src) {
	if (src != null && typeof (src) != "undefined" && $.trim(src) != "") {
		return false;
	}
	return true;
});

// 判断是否不为空
Handlebars.registerHelper("isNotBlank", function(src) {
	if (src != null && typeof (src) != "undefined" && $.trim(src) != "") {
		return true;
	}
	return false;
});

/* ==================数字相关===================== */

// 数字相加
Handlebars.registerHelper("add", function(src, dst) {
	return src + dst;
});

// 数字相减
Handlebars.registerHelper("sub", function(src, dst) {
	return src - dst;
});

// 数字相乘
Handlebars.registerHelper("mul", function(src, dst) {
	return src * dst;
});

// 数字相除
Handlebars.registerHelper("div", function(src, dst) {
	return src / dst;
});

// 千位格式化 fmt money 1,000,000
// val 值
Handlebars.registerHelper("fmtMoney", function(val) {
	val = String(val);
	var reg = /(-?\d+)(\d{3})/;
	while (reg.test(val)) {
		val = val.replace(reg, "$1,$2");
	}
	return val;
});

// fmt money fiexd 1000.00
// val 值
// fiexd 保留位数
Handlebars.registerHelper("fmtMoneyFiexd", function(val, fiexd) {
	val = String(val);
	var valArr = val.split(".");
	var w = fiexd, arr1 = "", zero = "";
	if (valArr.length > 1) {
		arr1 = valArr[1];
		w = fiexd - arr1.length;
	}
	var result = valArr[0] + ".";
	for (var i = 0; i < w; i++) {
		zero += "0";
	}
	val = result + arr1 + zero;
	return val;
});

// 取模 x%y==0 return bool
Handlebars.registerHelper("divisible", function(x, y) {
	if (x % y == 0) {
		return true;
	}
	return false;
});
/* ==================系统模块相关===================== */

// 根据id获取用户姓名
Handlebars.registerHelper("username", function(id) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/user/name",
		data : {
			id : id,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				value = result.data;
			}
		}
	});
	return value;
});

// 通过字典类型和键获取值
Handlebars.registerHelper("dictval", function(genre, mkey) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/dict/mval",
		data : {
			genre : genre,
			mkey : mkey,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				value = result.data;
			}
		}
	});
	return value;
});

// 通过字典类型和键列表字符串（逗号分隔）获取值列表字符串（逗号分隔）
Handlebars.registerHelper("dictvals", function(genre, mkeys) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/dict/mvals",
		data : {
			genre : genre,
			mkeys : mkeys,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				value = result.data;
			}
		}
	});
	return value;
});

// 根据id获取区域名称
Handlebars.registerHelper("areaname", function(id) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/area/name",
		data : {
			id : id,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				value = result.data;
			}
		}
	});
	return value;
});

// 根据id获取区域全称
Handlebars.registerHelper("areafullname", function(id) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/area/fullname",
		data : {
			id : id,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				if (result.data != null && result.data.length > 0) {
					value = result.data;
					value = value.replace(" - ", "");
					value = value.replace(" - ", "");
				}
			}
		}
	});
	return value;
});

// 根据id获取区域省份名称
Handlebars.registerHelper("areappname", function(id) {
	var value = "";
	$.ajax({
		async : false,
		url : apis + "/sys/area/parentpname",
		data : {
			id : id,
			random : Math.random()
		},
		type : "GET",
		dataType : 'json',
		success : function(result) {
			if (result.success) {
				if (result.data != null && result.data.length > 0) {
					value = result.data;
				}
			}
		}
	});
	return value;
});
