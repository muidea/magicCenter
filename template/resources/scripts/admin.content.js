

var content = {
		article :{
			articleInfo :{}
		},
		catalog :{
			catalogInfo :{}
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
		
		content.fillArticleView();
		content.fillCatalogView();
		
	}, "json");
	
    // 绑定表单提交事件处理器
    $('#article-content .article-Form').submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                resetForm:	 true,
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
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
                resetForm:	 true,
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //alert("before submit");
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

content.fillArticleView = function() {
	$("#article-List div.notification").hide();
	
	if (content.errCode > 0) {
		$("#article-List div.error div").html(result.Reason);
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
	titleLink.setAttribute("onclick","content.editArticle('/content/admin/queryArticle/?id=" + articleInfo.Id + "'); return false;" );
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
	editLink.setAttribute("onclick","content.editArticle('/content/admin/queryArticle/?id=" + articleInfo.Id + "'); return false" );
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
		$("#catalog-List div.error div").html(result.Reason);
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
	titleLink.setAttribute("onclick","content.editCatalog('/content/admin/queryCatalog/?id=" + catalog.Id + "'); return false;" );
	titleLink.innerHTML = catalog.Name;
	titleTd.appendChild(titleLink);
	tr.appendChild(titleTd);

	var createrTd = document.createElement("td");
	createrTd.innerHTML = catalog.Creater.NickName;
	tr.appendChild(createrTd);
	
	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editCatalog" );
	editLink.setAttribute("onclick","content.editCatalog('/content/admin/queryCatalog/?id=" + catalog.Id + "'); return false;" );
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
		for (var ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
			catalog = content.catalog.catalogInfo[ii];
			
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
			return;
		}
		
		$("#catalog-Edit .catalog-Form .catalog-id").val(result.Catalog.Id);
		$("#catalog-Edit .catalog-Form .catalog-name").val(result.Catalog.Name);
		
		$("#catalog-Edit .content-box-tabs li a").removeClass('current');
		$("#catalog-Edit .content-box-tabs li a.catalog-Edit-tab").addClass('current');
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






