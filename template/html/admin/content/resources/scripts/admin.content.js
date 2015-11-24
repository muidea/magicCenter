

var content = {
		article :{
			articleInfo :{}
		},
		catalog :{
			catalogInfo :{}
		},
		link :{
			linkInfo :{}
		},
		image :{
			imageInfo :{}
		}
};

content.initialize = function(accessCode, view) {
	content.accessCode = accessCode;
	
	$.post("/content/admin/queryAllContent/", {
		accesscode: accessCode
	}, function(result){
		content.errCode = result.ErrCode;
		content.reason = result.Reason;
		
		content.article.articleInfo = result.ArticleInfo;		
		content.catalog.catalogInfo = result.Catalog;
		content.link.linkInfo = result.Link;
		content.image.imageInfo = result.Image;
		
		content.fillArticleView();
		content.fillCatalogView();
		content.fillLinkView();
		content.fillImageView();
		
	}, "json");
	
    // 绑定表单提交事件处理器
    $("#article-content .article-Form").submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  "json"        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //alert("before submit");
        } 
        // post-submit callback
        function showResponse(result) {
        	$("#article-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#article-Edit div.error div").html(result.Reason);
        		$("#article-Edit div.error").show();
        	} else {
        		$("#article-Edit div.success div").html(result.Reason);
        		$("#article-Edit div.success").show();

        		content.refreshArticle();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#article-content .article-Form .article-title").parent().find("span").remove();
        	var title = $("#article-content .article-Form .article-title").val();
        	if (title.length == 0) {
        		$("#article-content .article-Form .article-title").parent().append("<span class=\"input-notification error png_bg\">请输入标题</span>");
        		result = false;
        	}
        	
    		$("#article-content .article-Form .article-content").parent().find("span").remove();        			
        	var content = $("#article-content .article-Form .article-content").val();
        	if (content.length == 0) {
        		$("#article-content .article-Form .article-content").parent().append("<span class=\"input-notification error png_bg\">请输入内容</span>");
        		result = false;
        	}

    		$("#article-content .article-Form .article-catalog").parent().find("span").remove();        			
    		var catalog = $("#article-content .article-Form .article-catalog").val();
        	if (catalog == -1) {
        		$("#article-content .article-Form .article-catalog").parent().append("<span class=\"input-notification error png_bg\">请选择分类</span>");
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
    
    // 绑定表单提交事件处理器
    $('#catalog-content .catalog-Form').submit(function() {
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
        	$("#catalog-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#catalog-Edit div.error div").html(result.Reason);
        		$("#catalog-Edit div.error").show();
        	} else {
        		$("#catalog-Edit div.success div").html(result.Reason);
        		$("#catalog-Edit div.success").show();

        		content.refreshCatalog();
        	}        	
        }
        
        function validate() {
        	var result = true
        	
        	$("#catalog-content .catalog-Form .catalog-name").parent().find("span").remove();
        	var title = $("#catalog-content .catalog-Form .catalog-name").val();
        	if (title.length == 0) {
        		$("#catalog-content .catalog-Form .catalog-name").parent().append("<span class=\"input-notification error png_bg\">请输入分类名</span>");
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
    
    // 绑定表单提交事件处理器
    $('#link-content .link-Form').submit(function() {
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
        	$("#link-Edit div.notification").hide();
        	
        	if (result.ErrCode > 0) {
        		$("#link-Edit div.error div").html(result.Reason);
        		$("#link-Edit div.error").show();
        	} else {
        		$("#link-Edit div.success div").html(result.Reason);
        		$("#link-Edit div.success").show();

        		content.refreshLink();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#link-content .link-Form .link-name").parent().find("span").remove();
        	var name = $("#link-content .link-Form .link-name").val();
        	if (name.length == 0) {
        		$("#link-content .link-Form .link-name").parent().append("<span class=\"input-notification error png_bg\">请输入站点名</span>");
        		result = false;
        	}
        	$("#link-content .link-Form .link-url").parent().find("span").remove();
        	var url = $("#link-content .link-Form .link-url").val();
        	if (url.length == 0) {
        		$("#link-content .link-Form .link-url").parent().append("<span class=\"input-notification error png_bg\">请输入站点网址</span>");
        		result = false;
        	}
        	
        	var value = $("#link-content .link-Form .link-style").val();
        	
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

        		content.refreshImage();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#image-content .image-Form .image-name").parent().find("span").remove();
        	var name = $("#image-content .image-Form .image-name").val();
        	if (name.length == 0) {
        		$("#image-content .image-Form .image-name").parent().append("<span class=\"input-notification error png_bg\">请选择图片</span>");
        		result = false;
        	}
        	$("#image-content .image-Form .image-desc").parent().find("span").remove();
        	var desc = $("#image-content .image-Form .image-desc").val();
        	if (desc.length == 0) {
        		$("#image-content .image-Form .image-url").parent().append("<span class=\"input-notification error png_bg\">请输入图片描述</span>");
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

content.refreshArticle = function() {
	$.post("/content/admin/queryAllArticle/", {
		accesscode: content.accessCode
	}, function(result){
		content.errCode = result.ErrCode;
		content.reason = result.Reason;
		
		content.article.articleInfo = result.ArticleInfo;		
		
		content.fillArticleView();
	}, "json");	
}

content.refreshCatalog = function() {
	$.post("/content/admin/queryAllCatalog/", {
		accesscode: content.accessCode
	}, function(result){
		content.errCode = result.ErrCode;
		content.reason = result.Reason;
		
		content.catalog.catalogInfo = result.Catalog;		
		
		content.fillCatalogView();		
	}, "json");	
}

content.refreshLink = function() {
	$.post("/content/admin/queryAllLink/", {
		accesscode: content.accessCode
	}, function(result){
		content.errCode = result.ErrCode;
		content.reason = result.Reason;
		
		content.link.linkInfo = result.Link;		
		
		content.fillLinkView();		
	}, "json");	
}

content.refreshImage = function() {
	$.post("/content/admin/queryAllImage/", {
		accesscode: content.accessCode
	}, function(result){
		content.errCode = result.ErrCode;
		content.reason = result.Reason;
		
		content.image.imageInfo = result.Image;		
		
		content.fillImageView();		
	}, "json");	
}

content.fillArticleView = function() {
	$("#article-List div.notification").hide();
	
	if (content.errCode > 0) {
		$("#article-List div.error div").html(content.reason);
		$("#article-List div.error").show();
		
		$("#article-List table").hide();
		return;
	}
	
	$("#article-List table tbody tr").remove();
	var articleInfoList = content.article.articleInfo;
	for (var ii =0; ii < articleInfoList.length; ++ii) {
		var articleInfo = articleInfoList[ii];
		var trContent = content.constructArticleItem(articleInfo);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#article-List table tbody").append(trContent);
	}
	$("#article-List table").show();
	
	$("#article-Edit div.notification").hide();
	$("#article-Edit .article-Form .article-id").val(-1);
	$("#article-Edit .article-Form .article-title").val("");
	$("#article-Edit .article-Form .article-content").wysiwyg("setContent", "");
	$("#article-Edit .article-Form .article-catalog").empty();	
	$("#article-Edit .article-Form .article-catalog").append("<option value=-1>请选择分类</option>");
	for (var ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
		catalog = content.catalog.catalogInfo[ii];
		
		$("#article-Edit .article-Form .article-catalog").append("<option value="+  catalog.Id + ">" + catalog.Name + "</option>");
	}
}

content.constructArticleItem = function(articleInfo) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","article");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var titleTd = document.createElement("td");
	var titleLink = document.createElement("a");
	titleLink.setAttribute("class","edit");
	titleLink.setAttribute("href","#queryArticle");
	titleLink.setAttribute("onclick","content.editArticle('/content/admin/editArticle/?id=" + articleInfo.Id + "'); return false;" );
	titleLink.innerHTML = articleInfo.Title;
	titleTd.appendChild(titleLink);
	tr.appendChild(titleTd);

	var cataLogTd = document.createElement("td");
	cataLogTd.innerHTML = articleInfo.Catalog.Name;
	tr.appendChild(cataLogTd);

	var authorTd = document.createElement("td");
	authorTd.innerHTML = articleInfo.Author.NickName;
	tr.appendChild(authorTd);
	
	var createDateTd = document.createElement("td");
	createDateTd.innerHTML = articleInfo.CreateDate;
	tr.appendChild(createDateTd);

	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#queryArticle");
	editLink.setAttribute("onclick","content.editArticle('/content/admin/editArticle/?id=" + articleInfo.Id + "'); return false" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteArticle" );
	deleteLink.setAttribute("onclick","content.deleteArticle('/content/admin/deleteArticle/?id=" + articleInfo.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
}

content.fillCatalogView = function() {
	$("#catalog-List div.notification").hide();
	
	if (content.errCode > 0) {
		$("#catalog-List div.error div").html(content.reason);
		$("#catalog-List div.error").show();
		
		$("#catalog-List table").hide();
		return;
	}
	
	$("#catalog-List table tbody tr").remove();
	var catalogList = content.catalog.catalogInfo;
	for (var ii =0; ii < catalogList.length; ++ii) {
		var catalog = catalogList[ii];
		var trContent = content.constructCatalogItem(catalog);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#catalog-List table tbody").append(trContent);
	}
	$("#catalog-List table").show();
	
	$("#catalog-Edit div.notification").hide();		
	$("#catalog-Edit .catalog-Form .catalog-id").val(-1);
	$("#catalog-Edit .catalog-Form .catalog-name").val("");
	$("#catalog-Edit .catalog-Form .catalog-parent").empty();
	$("#catalog-Edit .catalog-Form .catalog-parent").append("<option value=-1>请选择父类</option>");
	for (var ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
		catalog = content.catalog.catalogInfo[ii];
		
		$("#catalog-Edit .catalog-Form .catalog-parent").append("<option value="+  catalog.Id + ">" + catalog.Name + "</option>");
	}	
};


content.constructCatalogItem = function(catalog) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","catalog");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var titleTd = document.createElement("td");
	var titleLink = document.createElement("a");
	titleLink.setAttribute("class","edit");
	titleLink.setAttribute("href","#editCatalog" );
	titleLink.setAttribute("onclick","content.editCatalog('/content/admin/editCatalog/?id=" + catalog.Id + "'); return false;" );
	titleLink.innerHTML = catalog.Name;
	titleTd.appendChild(titleLink);
	tr.appendChild(titleTd);

	var parentName = "-";
	for (var ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
		c = content.catalog.catalogInfo[ii];
		if (c.Id == catalog.Pid) {
			parentName = c.Name;
			break;
		}
	}
	var parentTd = document.createElement("td");
	parentTd.innerHTML = parentName;
	tr.appendChild(parentTd);
	
	var createrTd = document.createElement("td");
	createrTd.innerHTML = catalog.Creater.NickName;
	tr.appendChild(createrTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editCatalog" );
	editLink.setAttribute("onclick","content.editCatalog('/content/admin/editCatalog/?id=" + catalog.Id + "'); return false;" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteCatalog" );
	deleteLink.setAttribute("onclick","content.deleteCatalog('/content/admin/deleteCatalog/?id=" + catalog.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};


content.fillLinkView = function() {
	$("#link-List div.notification").hide();
	
	if (content.errCode > 0) {
		$("#link-List div.error div").html(content.reason);
		$("#link-List div.error").show();
		
		$("#link-List table").hide();
		return;
	}
	
	$("#link-List table tbody tr").remove();
	var linkList = content.link.linkInfo;
	for (var ii =0; ii < linkList.length; ++ii) {
		var link = linkList[ii];
		var trContent = content.constructLinkItem(link);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#link-List table tbody").append(trContent);
	}
	$("#link-List table").show();
	
	$("#link-Edit div.notification").hide();		
	$("#link-Edit .link-Form .link-id").val(-1);
	$("#link-Edit .link-Form .link-name").val("");
	$("#link-Edit .link-Form .link-url").val("");
	$("#link-Edit .link-Form .link-logo").val("");
};


content.constructLinkItem = function(link) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","link");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var nameTd = document.createElement("td");
	var nameLink = document.createElement("a");
	nameLink.setAttribute("class","edit");
	nameLink.setAttribute("href","#editLink" );
	nameLink.setAttribute("onclick","content.editLink('/content/admin/editLink/?id=" + link.Id + "'); return false;" );
	nameLink.innerHTML = link.Name;
	nameTd.appendChild(nameLink);
	tr.appendChild(nameTd);
	
	var createrTd = document.createElement("td");
	createrTd.innerHTML = link.Creater.NickName;
	tr.appendChild(createrTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editLink" );
	editLink.setAttribute("onclick","content.editLink('/content/admin/editLink/?id=" + link.Id + "'); return false;" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteLink" );
	deleteLink.setAttribute("onclick","content.deleteLink('/content/admin/deleteLink/?id=" + link.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};


content.fillImageView = function() {
	$("#image-List div.notification").hide();
	
	if (content.errCode > 0) {
		$("#image-List div.error div").html(content.reason);
		$("#image-List div.error").show();
		
		$("#image-List table").hide();
		return;
	}
	
	$("#image-List table tbody tr").remove();
	var imageList = content.image.imageInfo;
	for (var ii =0; ii < imageList.length; ++ii) {
		var image = imageList[ii];
		var trContent = content.constructImageItem(image);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		$("#image-List table tbody").append(trContent);
	}
	$("#image-List table").show();
	
	$("#image-Edit div.notification").hide();		
	$("#image-Edit .image-Form .image-name").val("");
	$("#image-Edit .image-Form .image-desc").val("");
};


content.constructImageItem = function(image) {
	var tr = document.createElement("tr");
	tr.setAttribute("class","image");

	var checkBoxTd = document.createElement("td");
	var checkBox = document.createElement("input");
	checkBox.setAttribute("type","checkbox");
	
	checkBoxTd.appendChild(checkBox);
	tr.appendChild(checkBoxTd);

	var imageTd = document.createElement("td");
	var imageUrl = document.createElement("img");
	imageUrl.setAttribute("src", image.Url);
	imageUrl.setAttribute("alt", image.Desc);
	imageUrl.setAttribute("width", "50");
	imageUrl.setAttribute("height", "50");
	imageTd.appendChild(imageUrl);
	tr.appendChild(imageTd);
	
	var createrTd = document.createElement("td");
	createrTd.innerHTML = image.Creater.NickName;
	tr.appendChild(createrTd);
	
	var editTd = document.createElement("td");
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteImage" );
	deleteLink.setAttribute("onclick","content.deleteImage('/content/admin/deleteImage/?id=" + image.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};

content.editArticle = function(editUrl) {
	$.post(editUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#article-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#article-List div.error div").html(result.Reason);
			$("#article-List div.error").show();
			return
		}
		
		$("#article-Edit .article-Form .article-id").val(result.Article.Id);
		$("#article-Edit .article-Form .article-title").val(result.Article.Title);
		$("#article-Edit .article-Form .article-content").wysiwyg("setContent", result.Article.Content);
		$("#article-Edit .article-Form .article-catalog").empty();
		
		$("#article-Edit .article-Form .article-catalog").append("<option value=-1>请选择分类</option>");
		for (var ii =0; ii < result.Catalog.length; ++ii) {
			catalog = result.Catalog[ii];
			
			$("#article-Edit .article-Form .article-catalog").append("<option value="+  catalog.Id + ">" + catalog.Name + "</option>");
			if (catalog.Id == result.Article.Catalog.Id) {
				$("#article-Edit .article-Form .article-catalog").get(0).selectedIndex = ii +1;				
			}
		}
				
		$("#article-content .content-box-tabs li a").removeClass('current');
		$("#article-content .content-box-tabs li a.article-Edit-tab").addClass('current');
		$("#article-Edit").siblings().hide();
		$("#article-Edit").show();
	}, "json");
}

content.deleteArticle = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#article-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#article-List div.error div").html(result.Reason);
			$("#article-List div.error").show();
			return
		}
		
		$("#article-List div.success div").html(result.Reason);
		$("#article-List div.success").show();
		
		content.refreshArticle();
	}, "json");
}

content.editCatalog = function(editUrl) {
	$.post(editUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#catalog-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#catalog-List div.error div").html(result.Reason);
			$("#catalog-List div.error").show();
			return
		}
		
		$("#catalog-Edit .catalog-Form .catalog-id").val(result.Catalog.Id);
		$("#catalog-Edit .catalog-Form .catalog-name").val(result.Catalog.Name);
		$("#catalog-Edit .catalog-Form .catalog-parent").empty();
		
		$("#catalog-Edit .catalog-Form .catalog-parent").append("<option value=-1>请选择父类</option>");
		for (var ii =0; ii < result.ParentCatalog.length; ++ii) {
			catalog = result.ParentCatalog[ii];
			
			$("#catalog-Edit .catalog-Form .catalog-parent").append("<option value="+  catalog.Id + ">" + catalog.Name + "</option>");
			if (catalog.Id == result.Catalog.Pid) {
				$("#catalog-Edit .catalog-Form .catalog-parent").get(0).selectedIndex = ii +1;				
			}
		}
		
		$("#catalog-content .content-box-tabs li a").removeClass('current');
		$("#catalog-content .content-box-tabs li a.catalog-Edit-tab").addClass('current');
		$("#catalog-Edit").siblings().hide();
		$("#catalog-Edit").show();		
	}, "json");
}

content.deleteCatalog = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#catalog-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#catalog-List div.error div").html(result.Reason);
			$("#catalog-List div.error").show();
			return;
		}
		
		$("#catalog-List div.success div").html(result.Reason);
		$("#catalog-List div.success").show();
		
		content.refreshCatalog();
	}, "json");
}


content.editLink = function(editUrl) {
	$.post(editUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#link-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#link-List div.error div").html(result.Reason);
			$("#link-List div.error").show();
			return
		}
		
		$("#link-Edit .link-Form .link-id").val(result.Link.Id);
		$("#link-Edit .link-Form .link-name").val(result.Link.Name);
		$("#link-Edit .link-Form .link-url").val(result.Link.Url);
		$("#link-Edit .link-Form .link-logo").val(result.Link.Logo);
		
		var linkStyles = $("#link-Edit .link-Form .link-style");
		for (var ii =0; ii < linkStyles.length; ++ii) {
			var style = linkStyles[ii];
			var value = style.getAttribute("value");
			if(value == result.Link.Style) {
				style.checked = true;
				break;
			}
		}
		
		$("#link-content .content-box-tabs li a").removeClass('current');
		$("#link-content .content-box-tabs li a.link-Edit-tab").addClass('current');
		$("#link-Edit").siblings().hide();
		$("#link-Edit").show();		
	}, "json");
}

content.deleteLink = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#link-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#link-List div.error div").html(result.Reason);
			$("#link-List div.error").show();
			return;
		}
		
		$("#link-List div.success div").html(result.Reason);
		$("#link-List div.success").show();
		
		content.refreshLink();
	}, "json");
}


content.deleteImage = function(deleteUrl) {
	$.post(deleteUrl, {
		accesscode: content.accessCode
	}, function(result) {
		$("#image-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#image-List div.error div").html(result.Reason);
			$("#image-List div.error").show();
			return;
		}
		
		$("#image-List div.success div").html(result.Reason);
		$("#image-List div.success").show();
		
		content.refreshImage();
	}, "json");
}
