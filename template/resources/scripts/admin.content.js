

var content = {
		article :{
			view :{},
			articleInfo :{}
		},
		catalog :{
			view :{},
			catalogInfo :{}
		}
};

content.initialize = function(accessCode, view) {
	var articleView = view.find("#article-content");
	var catalogView = view.find("#catalog-content");
	
	content.accessCode = accessCode;
	content.article.view = articleView;		
	content.catalog.view = catalogView;
	
	$.post("/content/admin/queryAllContent/", {
		accesscode: accessCode
	}, function(result){
		content.ErrCode = result.ErrCode;
		content.Reason = result.Reason;
		
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
                dataType:  'json'        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
            //alert("before submit");
        } 
        // post-submit callback
        function showResponse(result) {
        	if (result.ErrCode > 0) {
        		var notificationDiv = $(content.article.view).find("#article-Edit div.error");
        		notificationDiv.children("div").html(result.Reason);
        		notificationDiv.show();
        	} else {
        		var notificationDiv = $(content.article.view).find("#article-Edit div.success");
        		notificationDiv.children("div").html(result.Reason);
        		notificationDiv.show();
        		
        		content.refreshArticle();
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
		content.ErrCode = result.ErrCode;
		content.Reason = result.Reason;
		
		content.article.articleInfo = result.ArticleInfo;		
		
		content.fillArticleView();
	}, "json");	
}

content.fillArticleView = function() {
	var articleTable = content.article.view.find("#article-List").children("table");
	var notificationDiv = content.article.view.find("#article-List").children("div");
	if (content.ErrCode > 0) {
		notificationDiv.children("div").html(content.Reason);
		notificationDiv.show();
		articleTable.hide();
		return;
	}
	
	var articleListBody = articleTable.children("tbody");
	articleListBody.children("tr").remove();
	
	var articleInfoList = content.article.articleInfo;
	for (ii =0; ii < articleInfoList.length; ++ii) {
		var articleInfo = articleInfoList[ii];
		var trContent = content.constructArticleItem(articleInfo);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		articleListBody.append(trContent);
	}
	
	$("#article-Edit .article-Form .article-id").val(-1);
	$("#article-Edit .article-Form .article-title").val("");
	$("#article-Edit .article-Form .article-content").wysiwyg("setContent", "");
	$("#article-Edit .article-Form .article-catalog").empty();	
	$("#article-Edit .article-Form .article-catalog").append("<option value=-1>请选择分类</option>");
	for (ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
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
	titleLink.setAttribute("onclick","content.editArticle('/content/admin/queryArticle/?id=" + articleInfo.Id + "')" );
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
	editLink.setAttribute("onclick","content.editArticle('/content/admin/queryArticle/?id=" + articleInfo.Id + "')" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteArticle" );
	deleteLink.setAttribute("onclick","content.deleteArticle('/content/admin/deleteArticle/?id=" + articleInfo.Id + "')" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
}

content.fillCatalogView = function() {
	var catalogTable = content.catalog.view.find("#catalog-List").children("table");
	var notificationDiv = content.catalog.view.find("#catalog-List").children("div");
	if (content.ErrCode > 0) {
		notificationDiv.children("div").html(content.Reason);
		notificationDiv.show();
		catalogTable.hide();
		return;
	}
	
	var catalogListBody = catalogTable.children("tbody");
	catalogListBody.children("tr").remove();
	
	var catalogList = content.catalog.catalogInfo;
	for (ii =0; ii < catalogList.length; ++ii) {
		var catalog = catalogList[ii];
		var trContent = content.constructCatalogItem(catalog);
		if (ii % 2 == 1) {
			trContent.setAttribute("class","alt-row");
		}
		catalogListBody.append(trContent);
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
	titleLink.setAttribute("onclick","content.editCatalog('/content/admin/queryCatalog/?id=" + catalog.Id + "')" );
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
	editLink.setAttribute("onclick","content.editCatalog('/content/admin/queryCatalog/?id=" + catalog.Id + "')" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteCatalog" );
	deleteLink.setAttribute("onclick","content.deleteCatalog('/content/admin/deleteCatalog/?id=" + catalog.Id + "')" );
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
		if (result.ErrCode > 0) {
			$("#article-List .notification div").html(result.Reason);
			$("#article-List .notification").show();
			return
		}
		
		$("#article-Edit .article-Form .article-id").val(result.Article.Id);
		$("#article-Edit .article-Form .article-title").val(result.Article.Title);
		$("#article-Edit .article-Form .article-content").wysiwyg("setContent", result.Article.Content);
		$("#article-Edit .article-Form .article-catalog").empty();
		
		$("#article-Edit .article-Form .article-catalog").append("<option value=-1>请选择分类</option>");
		for (ii =0; ii < content.catalog.catalogInfo.length; ++ii) {
			catalog = content.catalog.catalogInfo[ii];
			
			$("#article-Edit .article-Form .article-catalog").append("<option value="+  catalog.Id + ">" + catalog.Name + "</option>");
			if (catalog.Id == result.Article.Catalog.Id) {
				$("#article-Edit .article-Form .article-catalog").get(0).selectedIndex = ii +1;				
			}
		}
				
		$(content.article.view).find(".content-box-tabs li a").removeClass('current');
		$(content.article.view).find(".content-box-tabs li a.article-Edit-tab").addClass('current');
		$("#article-Edit").siblings().hide();
		$("#article-Edit").show();
	}, "json");
}

content.submitArticle = function() {
	
}

content.deleteArticle = function(deleteUril) {
	$.post(deleteUril, {
		accesscode: content.accessCode
	}, function(result) {
		if (result.ErrCode > 0) {
			$("#article-List .notification div.error").html(result.Reason);
			$("#article-List .notification").show();
			return
		}
		
		content.refreshArticle();
	}, "json");
}

content.editCatalog = function(editUrl) {
	$.post(editUrl, {
		accesscode: content.catalog.accesscode
	}, function(result){
		alert(result.ErrCode);
		//content.catalog.catalogInfo = result;
		
		//console.log("content.catalog.queryAllCatalog, ret:%d", result.ErrCode);
		
		//fillCatalogViewCallBack();
	}, "json");
}

content.deleteCatalog = function(deleteUrl) {
	alert(deleteUrl);
	return false;
}






