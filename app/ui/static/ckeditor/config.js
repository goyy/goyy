/*

加粗		斜体		下划线	穿过线	下标字	上标字
['Bold','Italic','Underline','Strike','Subscript','Superscript'],
数字列表	实体列表	 减小缩进	增大缩进
['NumberedList','BulletedList','-','Outdent','Indent'],
左对齐	居中对齐	右对齐	 两端对齐
['JustifyLeft','JustifyCenter','JustifyRight','JustifyBlock'],
超链接	取消超链接		锚点
['Link','Unlink','Anchor'],
图片	flash	表格		水平线	表情		特殊字符	分页符	Iframe
['Image','Flash','Table','HorizontalRule','Smiley','SpecialChar','PageBreak','Iframe'],
样式		格式		字体		 字体大小
['Styles','Format','Font','FontSize'],
文本颜色		背景颜色
['TextColor','BGColor'],
全屏		显示区块
['Maximize', 'ShowBlocks','-']

*/

CKEDITOR.editorConfig = function( config ) {
	config.language = 'zh-cn';
	config.font_names = '宋体/宋体;黑体/黑体;仿宋/仿宋_GB2312;楷体/楷体_GB2312;隶书/隶书;幼圆/幼圆;微软雅黑/微软雅黑;' + config.font_names;
	config.toolbarGroups = 
		[
		 	{ name: 'links' },
		 	{ name: 'insert' },
		 	{ name: 'basicstyles', groups: [ 'basicstyles', 'cleanup' ] },
		 	{ name: 'paragraph', groups: [ 'list', 'indent', 'align' ] },
		 	'/',
		 	{ name: 'styles' },
		 	{ name: 'colors' }
		];
	config.extraPlugins="lineheight";
	// 移除工具栏中不需要的按钮.
	config.removeButtons = 'Anchor,Flash,Table,HorizontalRule,Smiley,SpecialChar,PageBreak,Strike,Subscript,Superscript';

	// 移除图片和链接弹窗的高级tab页.
	config.removeDialogTabs = 'image:advanced;link:advanced;iframe:advanced;iframe:presentation';
};

CKEDITOR.stylesSet.add( 'default', [
 	/* Block Styles */
 	{ name : '首行缩进', element : 'p', styles : { 'text-indent' : '20pt' } },
 	/* Inline Styles */
 	{ name : '标注黄色', element : 'span', styles : { 'background-color' : 'Yellow' } },
 	{ name : '标注绿色', element : 'span', styles : { 'background-color' : 'Lime' } },
 	/* Object Styles */
 	{ name : '图片左对齐', element : 'img', attributes : { 'style' : 'padding: 5px; margin-right: 5px', 'border' : '2', 'align' : 'left' } },
 	{ name : '图片右对齐', element : 'img', attributes : { 'style' : 'padding: 5px; margin-left: 5px', 'border' : '2', 'align' : 'right' } },
 	{ name : '无边界表格', element : 'table', styles: { 'border-style': 'hidden', 'background-color' : '#E6E6FA' } }
]);
