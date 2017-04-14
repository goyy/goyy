$(function(){
	$("#sGenreEQ").combo({url:apis+"/sys/dict/genres"});
});

var ePostLoadForm=function(){
	$("#eForm").validate();
}