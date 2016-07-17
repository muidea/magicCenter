$(document).ready(function() {

    // 先把所有的隐藏起来，默认显示当前激活的项
    $("#main-nav li ul").hide();
    $("#main-nav li a.active").parent().find("ul").slideToggle("slow");

    $("#main-nav li a.nav-top-item").click(
        // 隐藏兄弟项，显示当前项
        function() {
            $(this).parent().siblings().find("a.nav-top-item").removeClass("active");
            $(this).parent().siblings().find("ul").hide();
            $(this).parent().siblings().find("ul").slideUp("normal");
            $(this).next().slideToggle("normal");
            $(this).addClass("active");
            return false;
        }
    );

    $("#main-nav li ul li a.nav-sub-item").click(
        // 隐藏兄弟项，显示当前项
        function() {
            $(this).parent().siblings().find("a.nav-sub-item").removeClass("active");
            $(this).addClass("active");
            return false;
        }
    );

    $("#main-nav li a.nav-top-item").click(
        function() {
            var hrefUrl = $(this).parent().find("ul li a.active").attr("href");
            $("#body-content").load(hrefUrl);
            return false;
        }
    );

    $("#main-nav li ul li a.nav-sub-item").click(
        function() {
            var hrefUrl = $(this).attr('href');
            $("#body-content").load(hrefUrl);
            return false;
        }
    );

    $("#main-nav li a.active").parent().find("ul li a.active").trigger("click");
});