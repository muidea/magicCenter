

var image = {
	errCode:0,
	reason:'',
	imageInfo:{},
	catalogInfo:{}
};


image.initialize = function() {
	
	image.refreshCatalog();
	
	image.fillImageView();
	
	
    // 绑定表单提交事件处理器
    $('#image-content .image-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //return false;
        } 
        // post-submit callback
        function showResponse(result) {
        	$("#image-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#image-Edit div.error div").html(result.Reason);
        		$("#image-Edit div.error").show();
        	} else {
        		$("#image-Edit div.success div").html(result.Reason);
        		$("#image-Edit div.success").show();

        		image.refreshImage();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#image-content .image-Form .image-url").parent().find("span").remove();
        	var url = $("#image-content .image-Form .image-url").val();
        	if (url.length == 0) {
        		$("#image-content .image-Form .image-url").parent().append("<span class=\"input-notification error png_bg\">请输入分类名</span>");
        		result = false;
        	}
        	
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
};

image.assignDefaltName = function() {
	var url = $("#image-Edit .image-Form .image-url").val();
	if (url.length == 0) {
		return;
	}
	
	var arr = url.split('\\');
	var fileName = arr[arr.length-1];
	arr = fileName.split('.');
	$("#image-Edit .image-Form .image-name").val(arr[0]);
}

image.refreshCatalog = function() {
	$("#image-Edit .image-Form .image-catalog").children().remove();
	for (var ii =0; ii < image.catalogInfo.length; ++ii) {
		var catalog = image.catalogInfo[ii];
		$("#image-Edit .image-Form .image-catalog").append("<input type='checkbox' name='image-catalog' value=" +  catalog.Id + "> </input> <span>" + catalog.Name + "</span> ");
	}		
};


image.refreshImage = function() {
	$.get("/admin/content/queryAllImage/", {
	}, function(result){
		image.errCode = result.ErrCode;
		image.reason = result.Reason;
		image.imageInfo = result.ImageInfo;
		
		image.fillImageView();
	}, "json");
};


image.fillImageView = function() {
	
	$("#image-List div.notification").hide();
	
	if ( image.errCode > 0) {
		$("#image-List div.error div").html(image.reason);
		$("#image-List div.error").show();
		
		$("#image-List table").hide();
		return;
	}
	
	$("#image-List table tbody tr").remove();
	if (image.imageInfo) {
		for (var ii =0; ii < image.imageInfo.length; ++ii) {
			var info = image.imageInfo[ii];
			var trContent = image.constructImageItem(info);
			$("#image-List table tbody").append(trContent);
		}		
	}
	$("#image-List table tbody tr:even").addClass("alt-row");
	$("#image-List table").show();
	
	$("#image-Edit div.notification").hide();		
	$("#image-Edit .image-Form .image-id").val(-1);
	$("#image-Edit .image-Form .image-name").val("");
	$("#image-Edit .image-Form .image-url").val("");
	$("#image-Edit .image-Form .image-desc").wysiwyg("setContent", "");
};


image.constructImageItem = function(img) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","image");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);
	
	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","edit");
	nameLink.setAttribute("href","#editCatalog" );
	nameLink.setAttribute("onclick","image.editImage('/admin/content/editImage/?id=" + img.Id + "'); return false;" );
	nameLink.innerHTML = img.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);

	var urlTd = document.createElement("td");
	urlTd.innerHTML = img.Url;
	tr.appendChild(urlTd);

	var descTd = document.createElement("td");
	descTd.innerHTML = img.Desc;
	tr.appendChild(descTd);
	
	var catalogTd = document.createElement("td");
	var catalogs = ""
	for (var ii =0; ii < img.Catalog.length; ) {
		catalogs += img.Catalog[ii++].Name
		if (ii < img.Catalog.length) {
			catalogs += ","
		} else {
			break;
		}
	}	
	catalogs = catalogs.length == 0 ? '-' :catalogs;	
	catalogTd.innerHTML = catalogs;
	tr.appendChild(catalogTd);	
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editImage" );
	editLink.setAttribute("onclick","image.editImage('/admin/content/editImage/?id=" + img.Id + "'); return false;" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteImage" );
	deleteLink.setAttribute("onclick","image.deleteImage('/admin/content/deleteImage/?id=" + img.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};

image.editImage = function(editUrl) {
	$.get(editUrl, {
	}, function(result) {
		$("#image-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#image-List div.error div").html(result.Reason);
			$("#image-List div.error").show();
			return
		}
		
		$("#image-Edit .image-Form .image-id").val(result.Image.Id);
		$("#image-Edit .image-Form .image-url").val(result.Image.Url);
		$("#image-Edit .image-Form .image-desc").wysiwyg("setContent", result.Image.Desc);
		
		$("#image-Edit .image-Form .image-catalog input").prop("checked", false);
		if (result.Image.Catalog) {
			for (var ii =0; ii < result.Image.Catalog.length; ++ii) {
				var ca = result.Image.Catalog[ii];
				$("#image-Edit .image-Form .image-catalog input").filter("[value="+ ca +"]").prop("checked", true);			
			}
		}
		
		$("#image-content .content-box-tabs li a").removeClass('current');
		$("#image-content .content-box-tabs li a.image-Edit-tab").addClass('current');
		$("#image-Edit").siblings().hide();
		$("#image-Edit").show();		
	}, "json");
};

image.deleteImage = function(deleteUrl) {
	$.get(deleteUrl, {
	}, function(result) {
		$("#image-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#image-List div.error div").html(result.Reason);
			$("#image-List div.error").show();
			return;
		}
		
		$("#image-List div.success div").html(result.Reason);
		$("#image-List div.success").show();
		
		image.refreshImage();
	}, "json");
};


