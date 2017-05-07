(function($) {
    $._tree = {
        consts: {
            idsuffix: "_elid",
            treeidsuffix: "_Tree",
            tnamesuffix: "_Tname",
            dialogsuffix: "_Dialog",
        },
        datas: {},
        getData: function(key) {
            return this.datas[key];
        },
        setData: function(key, value) {
            this.datas[key] = value;
        },
        buttons: function(prefix, input, reset) {
            var buttons = {};
            if (reset) { // 引入时传入变量 : btns='reset'
                buttons["重置"] = function() {
                    $._tree.treeSetValue(prefix, input, "", "");
                    $(this).dialog("close");
                }
            }
            buttons["取消"] = function() {
                $(this).dialog("close");
            }
            buttons["确定"] = function() {
                var nodes = $._tree.getNodes(prefix, input).getSelectedNodes();
                if (nodes.length == 1) {
                    $._tree.treeSetValue(prefix, input, nodes[0].id, nodes[0].name);
                }
                $(this).dialog("close");
            }
            return buttons;
        },
        getNodes(prefix, input) {
            return $.fn.zTree.getZTreeObj(prefix + input + $._tree.consts.treeidsuffix);
        },
        treeSetValue: function(prefix, input, id, name) {
            var _id = $._tree.getData((prefix + input) + $._tree.consts.idsuffix);
            var treeform = $("#" + _id);
            treeform.find("input[name='" + input + "']").val(id);
            treeform.find("input[name='" + input + $._tree.consts.tnamesuffix + "']").val(name);
        },
        treeOpen: function(prefix, input, reset) {
            var buttons = $._tree.buttons(prefix, input, reset);
            $("#" + prefix + input + $._tree.consts.dialogsuffix).dialog({
                resizable: false,
                modal: true,
                height: 500,
                buttons: buttons
            });
        },
        initData: function(args, data) {
            var setting = {
                data: {
                    simpleData: {
                        enable: true,
                        rootPId: "root"
                    }
                }
            };
            $.fn.zTree.init($("#" + args.prefix + args.input + $._tree.consts.treeidsuffix), setting, data);
            var _id = $._tree.getData((args.prefix + args.input) + $._tree.consts.idsuffix);
            var treeform = $("#" + _id);
            var value = treeform.find("input[name='" + args.input + "']").val();
            if (data.length > 0 && $.isNotBlank(value)) {
                for (var i in data) {
                    if (data[i].id == value) {
                        treeform.find("input[name='" + args.input + $._tree.consts.tnamesuffix + "']").val(data[i].name);
                        var zTree = $._tree.getNodes(args.prefix, args.input);
                        var node = zTree.getNodeByParam("id", value);
                        zTree.selectNode(node);
                        break;
                    }
                }
            }
        },
        create: function($this, args) {
            var dialog = $("#" + args.prefix + args.input + $._tree.consts.dialogsuffix);
            if (dialog.length > 0) {
                dialog.remove();
            }
            var treeHtml = '<div id="' + args.prefix + args.input + $._tree.consts.dialogsuffix + '" title="请选择" style="display: none;"><ul id="' + args.prefix + args.input + $._tree.consts.treeidsuffix + '" class="ztree"></ul></div>';
            $("body").append($(treeHtml));
            var treeform = $this;
            treeform.append('<input type="hidden" name="' + args.input + '" value="' + args.value + '"/>');
            if (args.uitype == "semantic") {
                var thname = '<div class="ui action input"><input readonly="readonly" name="' + args.input + $._tree.consts.tnamesuffix + '" value="" placeholder="' + args.placeholder + '"/>' + '<div id="' + args.prefix + args.input + 'SelectTreeBtn" onclick="$._tree.treeOpen(\'' + args.prefix + '\',\'' + args.input + '\',' + args.reset + ');" class="ui button">选择</div></div>';
                treeform.append(thname);
            } else if (args.uitype == "bootstrap") {
                var thname = '<input class="form-control" readonly="readonly" name="' + args.input + $._tree.consts.tnamesuffix + '" value="" placeholder="' + args.placeholder + '"/>';
                var selectBtn = '<button id="' + args.prefix + args.input + 'SelectTreeBtn" type="button" class="btn btn-default" onclick="$._tree.treeOpen(\'' + args.prefix + '\',\'' + args.input + '\',' + args.reset + ');" >选择</button>';
                treeform.append(thname);
                treeform.append(selectBtn);
            }
            var url = apis + args.url;
            var key = args.prefix + args.input;
            var data = $._tree.getData(key);
            if (data == undefined) {
                $.ajax({
                    type: "get",
                    dataType: "json",
                    url: url,
                    data: {
                        rootId: args.rootId,
                        exclusionId: args.exclusionId,
                        isShowRoot: args.isShowRoot
                    },
                    success: function(result) {
                        if (result.success) {
                            var id = $this.attr("id");
                            if (id == undefined) {
                                id = "id_" + new Date().getTime();
                                $this.attr("id", id);
                            }
                            $._tree.setData(key + $._tree.consts.idsuffix, id);
                            $._tree.setData(key, result.data);
                            $._tree.initData(args, result.data);
                        }
                    }
                });
            } else {
                $._tree.initData(args, data);
            }
        }
    }
    $.fn.tree = function(options) {
        var $t = $(this);
        var data = {//配置参数
            prefix: $t.data("prefix"),
            input: $t.data("input"),
            url: $t.data("url"),
            reset: $t.data("reset"),
            placeholder: $t.data("placeholder"),
            uitype: $t.data("uitype"),
            value: $t.data("value"),
            rootId: $t.data("rootId"),
            exclusionId: $t.data("exclusionId"),
        };
        var args = $.extend({//默认参数
            rootId: "root",
            exclusionId: "-1",
            uitype: "bootstrap",
            isShowRoot: true,
            placeholder: "请选择...",
            reset: false,
            value: "",
            prefix: "e"
        },
        options, data);
        if($.isNotBlank(args.url)){
            $._tree.create(this, args);
        }
    }
})(jQuery);