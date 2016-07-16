var module = {
    moduleList: {},
    defaultModule: '',
    currentModule: {},
};

$(document).ready(function() {

    $("#module-List .button").click(
        function() {
            var enableList = "";
            var defaultModule = "";
            var radioArray = $("#module-List table tbody tr td :radio:checked");
            for (var ii = 0; ii < radioArray.length; ++ii) {
                var radio = radioArray[ii];
                if ($(radio).val() == 1) {
                    enableList += $(radio).attr("name");
                    enableList += ",";
                }
            }
            var defaultArray = $("#module-List .module input:checkbox:checked");
            if (defaultArray.length > 0) {
                var checkBox = defaultArray[0];
                defaultModule = $(checkBox).attr("name");
            }

            $.post("/admin/system/applyModuleSetting/", {
                "module-enableList": enableList,
                "module-defaultModule": defaultModule
            }, function(result) {

                module.moduleList = result.Modules;
                module.defaultModule = result.DefaultModule;

                $("#module-List label").addClass("hidden")
                if (result.ErrCode > 0) {
                    $("#module-List .danger").html(result.Reason);
                    $("#module-List .danger").removeClass("hidden");
                } else {
                    $("#module-List .success").html(result.Reason);
                    $("#module-List .success").removeClass("hidden");
                }

                //module.fillModuleView();	        	
            }, "json");
        }
    );

    // 绑定表单提交事件处理器
    $('#module-Maintain .block .block-Form').submit(function() {
        var options = {
            beforeSubmit: showRequest, // pre-submit callback
            success: showResponse, // post-submit callback
            dataType: 'json' // 'xml', 'script', or 'json' (expected server response type) 
        };

        // pre-submit callback
        function showRequest() {
            //return false;
        }
        // post-submit callback
        function showResponse(result) {
            console.log(result);
            if (result.ErrCode > 0) {
                $("#module-Maintain .alert-Info .content").html(result.Reason);
                $("#module-Maintain .alert-Info").modal();
            } else {
                module.currentModule = result.Module;
                module.fillModuleMaintainView();
            }
        }

        function validate() {
            var result = true
            return result;
        }

        if (!validate()) {
            return false;
        }

        //提交表单
        $(this).ajaxSubmit(options);

        // !!! Important !!!
        // 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
        return false;
    });
});

module.initialize = function() {
    module.fillModuleListView();
};

module.getModuleListView = function() {
    return $("#module-List table")
}

module.getModuleBlockListView = function() {
    return $("#module-Maintain .block table")
}

module.fillModuleListView = function() {
    var moduleListView = module.getModuleListView();
    $(moduleListView).find("tbody tr").remove();
    for (var ii = 0; ii < module.moduleList.length; ++ii) {
        var info = module.moduleList[ii];
        var trContent = module.constructModuleItem(info);

        $(moduleListView).find("tbody").append(trContent);
    }
};

module.fillModuleMaintainView = function() {
    var blockListView = module.getModuleBlockListView();
    $(blockListView).find("tbody tr").remove();
    if (module.currentModule && module.currentModule.Blocks) {
        for (var ii = 0; ii < module.currentModule.Blocks.length; ++ii) {
            var block = module.currentModule.Blocks[ii];
            var trContent = module.constructBlockItem(block);
            $(blockListView).find("tbody").append(trContent);
        }
    }

    $("#module-Maintain .block-Form .block-name").val("");
    $("#module-Maintain .block-Form .block-tag").val("");
    $("#module-Maintain .block-Form .block-style").filter("[value='0']").prop("checked", true);
    $("#module-Maintain .block-Form .module-id").val(module.currentModule.Id);
}

module.constructModuleItem = function(mod) {
    var tr = document.createElement("tr");
    tr.setAttribute("class", "module");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = mod.Name;
    tr.appendChild(nameTd);

    var descriptionTd = document.createElement("td");
    descriptionTd.innerHTML = mod.Description
    tr.appendChild(descriptionTd);

    var editTd = document.createElement("td");
    var radioGroup = document.createElement("radiobox");
    var enable_radio = document.createElement("input");
    enable_radio.setAttribute("type", "radio");
    enable_radio.setAttribute("name", mod.Id);
    enable_radio.setAttribute("value", "1");
    radioGroup.appendChild(enable_radio);
    var enable_span = document.createElement("span");
    enable_span.innerHTML = "启用";
    radioGroup.appendChild(enable_span);

    var disable_radio = document.createElement("input");
    disable_radio.setAttribute("type", "radio");
    disable_radio.setAttribute("name", mod.Id);
    disable_radio.setAttribute("value", "0");
    radioGroup.appendChild(disable_radio);
    if (mod.EnableFlag) {
        enable_radio.checked = true;
        disable_radio.checked = false;
    } else {
        enable_radio.checked = false;
        disable_radio.checked = true;
    }

    var disable_span = document.createElement("span");
    disable_span.innerHTML = "禁用";
    radioGroup.appendChild(disable_span);

    editTd.appendChild(radioGroup);

    var checkGroup = document.createElement("checkbox");
    var default_check = document.createElement("input");
    default_check.setAttribute("type", "checkbox");
    default_check.setAttribute("name", mod.Id);
    default_check.setAttribute("onclick", "module.selectDefaultModule('" + mod.Id + "');");
    checkGroup.appendChild(default_check);
    if (module.defaultModule == mod.Id) {
        default_check.checked = true;
    } else {
        default_check.checked = false;
    }


    var default_span = document.createElement("span");
    default_span.innerHTML = "设为默认 ";
    checkGroup.appendChild(default_span);

    editTd.appendChild(checkGroup);

    var settingButton = document.createElement("input");
    settingButton.setAttribute("type", "button");
    settingButton.setAttribute("onclick", "module.maintainModule('/admin/system/queryModuleDetail/?id=" + mod.Id + "'); return false;");
    settingButton.setAttribute("value", "模块设置");
    settingButton.setAttribute("class", "button");
    editTd.appendChild(settingButton);
    tr.appendChild(editTd);
    return tr;
};

module.selectDefaultModule = function(defaultModule) {
    $("#module-List .module input:checkbox").prop("checked", false);
    $("#module-List .module input:checkbox[name='" + defaultModule + "']").prop("checked", true);
};

module.maintainModule = function(maintainUrl) {
    $.get(maintainUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#module-List .alert-Info .content").html(result.Reason);
            $("#module-List .alert-Info").modal();
            return
        }

        module.currentModule = result.Module;
        module.fillModuleMaintainView();
        $("#module-Content .content-header .nav .module-Maintain").find("a").trigger("click");
    }, "json");
};

module.constructBlockItem = function(block) {

    var tr = document.createElement("tr");
    tr.setAttribute("class", "block");

    var nameTd = document.createElement("td");
    nameTd.innerHTML = block.Name
    tr.appendChild(nameTd);

    var styleTd = document.createElement("td");
    if (block.Style == 0) {
        styleTd.innerHTML = "链接";
    } else {
        styleTd.innerHTML = "内容";
    }
    tr.appendChild(styleTd);

    var editTd = document.createElement("td");
    var deleteLink = document.createElement("a");
    deleteLink.setAttribute("class", "delete");
    deleteLink.setAttribute("href", "#deleteBlock");
    deleteLink.setAttribute("onclick", "module.deleteBlock('/admin/system/deleteModuleBlock/?id=" + block.Id + "&owner=" + module.currentModule.Id + "'); return false;");
    var deleteImage = document.createElement("img");
    deleteImage.setAttribute("src", "/resources/images/icons/cross.png");
    deleteImage.setAttribute("alt", "Delete");
    deleteLink.appendChild(deleteImage);
    editTd.appendChild(deleteLink);
    tr.appendChild(editTd);
    return tr;
};

module.deleteBlock = function(deleteUrl) {
    $.get(deleteUrl, {}, function(result) {
        if (result.ErrCode > 0) {
            $("#module-List .alert-Info .content").html(result.Reason);
            $("#module-List .alert-Info").modal();
        } else {
            module.currentModule = result.Module;
            module.fillModuleMaintainView();
        }
    }, "json");
};