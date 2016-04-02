

var article = {
	errCode:0,
	reason:"",
	articleInfo:{},
	catalogInfo:{}
};

article.initialize = function() {
	article.refreshCatalog();
	article.fillArticleView();
	
    // 绑定表单提交事件处理器
    $("#article-content .article-Form").submit(function() {
        var options = { 
                beforeSubmit:  showRequest,  // pre-submit callback
                success:       showResponse,  // post-submit callback
                dataType:  "json"        // 'xml', 'script', or 'json' (expected server response type) 
            };
        
        // pre-submit callback
        function showRequest() {
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

        		article.refreshArticle();
        	}
        }
        
        function validate() {
        	var result = true
        	
        	$("#article-content .article-Form .article-title").parent().find("span").remove();
        	var title = $("#article-content .article-Form .article-title").val();
        	if (title.length == 0) {
        		$("#article-content .article-Form .article-title").parent().append("<span class=\"input-notification error\">请输入标题</span>");
        		result = false;
        	}
        	
    		$("#article-content .article-Form .article-content").parent().find("span").remove();        			
        	var content = $("#article-content .article-Form .article-content").val();
        	if (content.length == 0) {
        		$("#article-content .article-Form .article-content").parent().append("<span class=\"input-notification error\">请输入内容</span>");
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

article.refreshCatalog = function() {
		$("#article-Edit .article-Form .article-catalog").children().remove();
		for (var ii =0; ii < article.catalogInfo.length; ++ii) {
			var catalog = article.catalogInfo[ii];
			$("#article-Edit .article-Form .article-catalog").append("<input type='checkbox' name='article-catalog' value=" +  catalog.Id + "> </input> <span>" + catalog.Name + "</span> ");
		}
};

article.refreshArticle = function() {
	$.get("/admin/content/queryAllArticle/", {
	}, function(result){
		article.errCode = result.ErrCode;
		article.reason = result.Reason;
		
		article.articleInfo = result.Articles;		
		
		article.fillArticleView();
	}, "json");	
};

article.fillArticleView = function() {
	$("#article-List div.notification").hide();
	
	if (article && article.errCode > 0) {
		$("#article-List div.error div").html(article.reason);
		$("#article-List div.error").show();
		
		$("#article-List table").hide();
		return;
	}
	
	$("#article-List table tbody tr").remove();
	var articleInfoList = article.articleInfo;
	for (var ii =0; ii < articleInfoList.length; ++ii) {
		var articleInfo = articleInfoList[ii];
		var trContent = article.constructArticleItem(articleInfo);
		$("#article-List table tbody").append(trContent);
	}
	
	$("#article-List table tbody tr:even").addClass("alt-row");
	$("#article-List table").show();
	
	$("#article-Edit div.notification").hide();
	$("#article-Edit .article-Form .article-id").val(-1);
	$("#article-Edit .article-Form .article-title").val("");
	$("#article-Edit .article-Form .article-content").wysiwyg("setContent", "");
};

article.constructArticleItem = function(articleInfo) {
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
	titleLink.setAttribute("onclick","article.editArticle('/admin/content/editArticle/?id=" + articleInfo.Id + "'); return false;" );
	titleLink.innerHTML = articleInfo.Title;
	titleTd.appendChild(titleLink);
	tr.appendChild(titleTd);

	var cataLogTd = document.createElement("td");
	var catalogs = "";
	if (articleInfo.Catalog) {
		for (var ii =0; ii < articleInfo.Catalog.length;) {
			catalogs += articleInfo.Catalog[ii++].Name
			if (ii < articleInfo.Catalog.length) {
				catalogs += ","
			} else {
				break;
			}
		}		
	}
	catalogs = catalogs.length == 0 ? '-' :catalogs;	
	cataLogTd.innerHTML = catalogs;
	tr.appendChild(cataLogTd);

	var authorTd = document.createElement("td");
	authorTd.innerHTML = articleInfo.Author.Name;
	tr.appendChild(authorTd);
	
	var createDateTd = document.createElement("td");
	createDateTd.innerHTML = articleInfo.CreateDate;
	tr.appendChild(createDateTd);

	var editTd = document.createElement("td");
	var editLink = document.createElement("a");
	editLink.setAttribute("class","edit");
	editLink.setAttribute("href","#editArticle");
	editLink.setAttribute("onclick","article.editArticle('/admin/content/editArticle/?id=" + articleInfo.Id + "'); return false" );
	var editImage = document.createElement("img");
	editImage.setAttribute("src","/resources/images/icons/pencil.png");
	editImage.setAttribute("alt","Edit");
	editLink.appendChild(editImage);	
	editTd.appendChild(editLink);
	
	var deleteLink = document.createElement("a");
	deleteLink.setAttribute("class","delete");
	deleteLink.setAttribute("href","#deleteArticle" );
	deleteLink.setAttribute("onclick","article.deleteArticle('/admin/content/deleteArticle/?id=" + articleInfo.Id + "'); return false;" );
	var deleteImage = document.createElement("img");
	deleteImage.setAttribute("src","/resources/images/icons/cross.png");
	deleteImage.setAttribute("alt","Delete");
	deleteLink.appendChild(deleteImage);	
	editTd.appendChild(deleteLink);
	
	tr.appendChild(editTd);
	
	return tr;
};

article.editArticle = function(editUrl) {
	$.get(editUrl, {
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
		$("#article-Edit .article-Form .article-catalog input").prop("checked", false);
		
		if (result.Article.Catalog) {
			for (var ii =0; ii < result.Article.Catalog.length; ++ii) {
				var ca = result.Article.Catalogs[ii];
				$("#article-Edit .article-Form .article-catalog input").filter("[value="+ ca.Id +"]").prop("checked", true);			
			}			
		}
		
		$("#article-content .content-box-tabs li a").removeClass('current');
		$("#article-content .content-box-tabs li a.article-Edit-tab").addClass('current');
		$("#article-Edit").siblings().hide();
		$("#article-Edit").show();
	}, "json");
};

article.deleteArticle = function(deleteUrl) {
	$.get(deleteUrl, {
	}, function(result) {
		$("#article-List div.notification").hide();
		
		if (result.ErrCode > 0) {
			$("#article-List div.error div").html(result.Reason);
			$("#article-List div.error").show();
			return
		}
		
		$("#article-List div.success div").html(result.Reason);
		$("#article-List div.success").show();
		
		article.refreshArticle();
	}, "json");
};

