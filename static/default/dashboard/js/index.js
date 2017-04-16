var magicCenter = {};

function setCookie(c_name, value, expiredays) {
    var exdate = new Date()
    exdate.setDate(exdate.getDate() + expiredays)
    document.cookie = c_name + "=" + escape(value) +
        ((expiredays == null) ? "" : ";expires=" + exdate.toGMTString())
}

function getCookie(c_name) {
    if (document.cookie.length > 0) {
        c_start = document.cookie.indexOf(c_name + "=")
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1
            c_end = document.cookie.indexOf(";", c_start)
            if (c_end == -1) c_end = document.cookie.length
            return unescape(document.cookie.substring(c_start, c_end))
        }
    }
    return ""
}

// 用户登出
magicCenter.logoutUserAction = function(authToken) {
    $.ajax({
        type: "DELETE",
        url: "/cas/user/?authToken=" + authToken,
        data: {},
        dataType: "json",
        success: function(data) {
            setCookie("userName", "", 0);
            setCookie("authToken", "", 0);
            location.href = "/static/dashboard/login.html";
        }
    });
};

$(document).ready(function() {
    magicCenter.viewVM = avalon.define({
        $id: "magicCenter",
        user: []
    });

    var user = getCookie("userName");
    var authToken = getCookie("authToken");

    console.log("user:" + user + ", authToken:" + authToken);
    if (user.length == 0 || authToken.length == 0) {
        location.href = "/static/dashboard/login.html";
        return;
    }
    magicCenter.curUser = user;
    magicCenter.authToken = authToken;

    magicCenter.viewVM.user = user;

    function updatePathNav() {
        var topPath = $("#main-nav>li>a.active").text();
        var subPath = $("#main-nav>li>a.active").parent().find("ul>li>a.active").text();

        $("#path-nav").find("li").remove();
        $("#path-nav").append("<li>" + topPath + "</li>");
        $("#path-nav").append("<li>" + subPath + "</li>");
    }


    // 先把所有的隐藏起来，默认显示当前激活的项
    $("#main-nav li ul").hide();
    $("#main-nav li a.active").parent().find("ul").slideToggle("slow");

    $("#main-nav li a.nav-top-item").click(
        // 隐藏兄弟项，显示当前项
        function() {
            var href = $(this).attr("href");
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

            updatePathNav();
            return false;
        }
    );

    $("#main-nav li ul li a.nav-sub-item").click(
        function() {
            var hrefUrl = $(this).attr('href');
            $("#body-content").load(hrefUrl);

            updatePathNav();
            return false;
        }
    );

    $("#main-nav li a.active").parent().find("ul li a.active").trigger("click");

    $("#profile-links .user-logout").click(
        function() {
            magicCenter.logoutUserAction(magicCenter.authToken);
        }
    );
});