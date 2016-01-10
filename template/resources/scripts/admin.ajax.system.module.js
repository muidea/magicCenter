

var module = {
	accesscode:'',
	moduleList:{}
};

module.initialize = function() {
	module.fillModuleView();
};

module.fillModuleView = function() {
	
	$("#module-content .content-list table tbody tr").remove();
	for (var ii =0; ii < module.moduleList.length; ++ii) {
		var info = module.moduleList[ii];
		var trContent = module.constructModuleItem(info);		
		$("#module-content .content-list table tbody").append(trContent);
	}
	
	$("#module-content .content-list table tbody tr:even").addClass("alt-row");	
};

module.constructModuleItem = function(module) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","module");
	
	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","view");
	nameLink.setAttribute("href", module.Uri );
	nameLink.innerHTML = module.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var descriptionTd = document.createElement("td");
	descriptionTd.innerHTML = module.Description
	tr.appendChild(descriptionTd);
	
	var editTd = document.createElement("td");
	var radioGroup = document.createElement("radiobox");
	var enable_radio = document.createElement("input");
	enable_radio.setAttribute("type","radio");
	enable_radio.setAttribute("name","enable-" + module.Id);
	enable_radio.setAttribute("value","1");
	radioGroup.appendChild(enable_radio);	
	var enable_span = document.createElement("span");
	enable_span.innerHTML ="启用";
	radioGroup.appendChild(enable_span);
	
	var disable_radio = document.createElement("input");
	disable_radio.setAttribute("type","radio");
	disable_radio.setAttribute("name","enable-" + module.Id);
	disable_radio.setAttribute("value","0");
	radioGroup.appendChild(disable_radio);
	if (module.Enable) {
		enable_radio.checked = true;
		disable_radio.checked = false;
	} else {
		enable_radio.checked = false;
		disable_radio.checked = true;		
	}
	
	var disable_span = document.createElement("span");
	disable_span.innerHTML ="禁用";
	radioGroup.appendChild(disable_span);
	
	editTd.appendChild(radioGroup);
	
	var checkGroup = document.createElement("checkbox");
	var default_check = document.createElement("input");
	default_check.setAttribute("type","checkbox");
	default_check.setAttribute("name","default-" + module.Id);
	checkGroup.appendChild(default_check);
	if (module.Default) {
		default_check.checked = true;
	} else {
		default_check.checked = false;
	}
	
	
	var default_span = document.createElement("span");
	default_span.innerHTML ="设为默认 ";
	checkGroup.appendChild(default_span);
	
	editTd.appendChild(checkGroup);	
	
	if(module.Internal == 0) {
		var uninstall = document.createElement("input");
		uninstall.setAttribute("type","button");
		uninstall.setAttribute("class","button");
		uninstall.setAttribute("value","卸载模块");
		editTd.appendChild(uninstall);
	}
	
	tr.appendChild(editTd);	
	return tr;
};


