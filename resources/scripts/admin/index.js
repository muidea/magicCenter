$(document).ready(function() {

    // Sidebar Accordion Menu:
    $("#main-nav li ul").hide(); // Hide all sub menus
    $("#main-nav li a.current").parent().find("ul").slideToggle("slow"); // Slide down the current

    $("#main-content .main-content-box").hide();
    var currrentDiv = $("#main-nav li a.current").attr('href');
    $(currrentDiv).slideToggle("slow");

    // menu item's sub menu
    $("#main-nav li a.nav-top-item").click( // When a top menu item is clicked...
        function() {
            $(this).parent().siblings().find("a.nav-top-item").removeClass("current");
            $(this).parent().siblings().find("ul").hide();
            $(this).parent().siblings().find("ul").slideUp("normal"); // Slide up all sub menus except the one clicked
            $(this).next().slideToggle("normal"); // Slide down the clicked sub menu
            $(this).addClass("current");

            return false;
        }
    );

    $("#main-nav li ul li a.nav-sub-item").click(
        // When a top menu item is clicked...
        function() {
            $(this).parent().siblings().find("a.nav-sub-item").removeClass("current"); // Remove class from all tabs
            $(this).addClass("current"); // Add class "current" to clicked tab

            return false;
        }
    );

    // Sidebar Accordion Menu Hover Effect:
    $("#main-nav li .nav-top-item").hover(
        function() {
            $(this).stop().animate({
                paddingRight: "25px"
            }, 200);
        },
        function() {
            $(this).stop().animate({
                paddingRight: "15px"
            });
        }
    );

    // menu item's sub menu
    $("#main-nav li a.nav-top-item").click( // When a top menu item is clicked...
        function() {
            var hrefUrl = $(this).parent().find("ul li a.current").attr("href");
            $("#body-content .body-content-box").load(hrefUrl);

            return false;
        }
    );

    $("#main-nav li ul li a.nav-sub-item").click(
        // When a top menu item is clicked...
        function() {
            var hrefUrl = $(this).attr('href');
            $("#body-content .body-content-box").load(hrefUrl);

            return false;
        }
    );

    $("#main-nav li:eq(0) ul li:eq(0) a").trigger("click");
});