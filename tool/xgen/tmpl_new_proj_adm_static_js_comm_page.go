// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmStaticJSCommPage = `(function($){
	var ms = {
		init: function(obj, args){
			return (function(){
				obj.empty();
				// 上一页
				if (args.pageNo > 1) {
					var func = args.pageFunc + "("+(args.pageNo-1)+", "+args.pageSize+")";
					obj.append('<li><a href="javascript:'+func+';">&laquo;</a></li>');
				} else {
					obj.append('<li class="disabled"><a href="javascript:">&laquo;</a></li>');
				}
				// 中间页码
				if (args.pageNo != 1 && args.pageNo >= 4 && args.totalPages != 4) {
					var func = args.pageFunc + "(1, "+args.pageSize+")";
					obj.append('<li><a href="javascript:'+func+';">'+1+'</a></li>');
				}
				if (args.pageNo-2 > 2 && args.pageNo <= args.totalPages && args.totalPages > 5) {
					obj.append('<li class="disabled"><a href="javascript:">...</a></li>');
				}
				var start = args.pageNo -2,end = args.pageNo+2;
				if ((start > 1 && args.pageNo < 4)||args.pageNo == 1) {
					end++;
				}
				if (args.pageNo > args.totalPages-4 && args.pageNo >= args.totalPages) {
					start--;
				}
				for (;start <= end; start++) {
					if(start <= args.totalPages && start >= 1){
						if(start != args.pageNo){
							var func = args.pageFunc + "("+start+", "+args.pageSize+")";
							obj.append('<li><a href="javascript:'+func+';">'+ start +'</a></li>');
						}else{
							obj.append('<li class="active"><a href="javascript:">'+ start +'</a></li>');
						}
					}
				}
				if(args.pageNo + 2 < args.totalPages - 1 && args.pageNo >= 1 && args.totalPages > 5){
					obj.append('<li class="disabled"><a href="javascript:">...</a></li>');
				}
				if(args.pageNo != args.totalPages && args.pageNo < args.totalPages -2  && args.totalPages != 4){
					var func = args.pageFunc + "("+args.totalPages+", "+args.pageSize+")";
					obj.append('<li><a href="javascript:'+func+';">'+args.totalPages+'</a></li>');
				}
				// 下一页
				if(args.pageNo < args.totalPages){
					var func = args.pageFunc + "("+(args.pageNo+1)+", " + args.pageSize+")";
					obj.append('<li><a href="javascript:'+func+';">&raquo;</a></li>');
				}else{
					obj.append('<li class="disabled"><a href="javascript:">&raquo;</a></li>');
				}
				// 转到指定页
				obj.append('<li class="disabled"><a href="javascript:" style="height:34px;"><%message "tmpl.form.page.current"%><input style="width:40px;height:20px;" type="text" value="'+args.pageNo+'" onkeydown="var e=e||event;var c=e.keyCode||e.which||e.charCode;if(c==13) '+args.pageFunc+'(this.value,'+args.pageSize+');" onclick="this.select();"/> / <input style="width:40px;height:20px;" type="text" value="'+args.pageSize+'" onkeydown="var e=e||event;var c=e.keyCode||e.which||e.charCode;if(c==13) '+args.pageFunc+'('+args.pageNo+',this.value);" onclick="this.select();"/><%message "tmpl.form.page.row"%>，<%message "tmpl.form.page.total"%>'+args.totalElements+'<%message "tmpl.form.page.row"%></a></li>');
			})();
		}
	}
	$.fn.pagination = function(options){
		var args = $.extend({
			pageNo : 1,
			pageSize : 10,
		    pageFunc: "page",
			totalPages : 0,
			totalElements : 0
		}, options);
		ms.init(this, args);
	}
})(jQuery);
`
