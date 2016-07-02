
$(document).ready(function () {
    $(".content-box .content-content .tab-content").hide();    
    $(".content-box .content-content .active").slideToggle("normal");
    
    $(".content-box .content-header ul li a").click(
        // 隐藏兄弟项，显示当前项
        function() {
            var tabId = $(this).attr('href');
            $(this).parent().siblings().removeClass("active");
            $(this).parent().addClass("active");
            
            $(tabId).siblings().hide();
            $(tabId).slideToggle("normal");
            
            return false;
        }
    );
    
    $(".content-box .content-header ul li.active").find("a").trigger("click");
});


