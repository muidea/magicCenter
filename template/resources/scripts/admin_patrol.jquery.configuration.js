var wysiwyg_editor = null;
function edit_route(url) {
	$
			.get(
					url,
					function(ret) {
						$(
								'.patrol-line .patrol-line-header .content-box-tabs a.routelineList')
								.removeClass('current');
						$(
								'.patrol-line .patrol-line-header .content-box-tabs a.routelineEdit')
								.addClass('current');
						$('.patrol-line .patrol-line-content #routelineList')
								.hide();
						$('.patrol-line .patrol-line-content #routelineEdit')
								.show();

						console.log(JSON.stringify(ret));
						$(
								".patrol-line .patrol-line-content #routelineEdit #routeline-id")
								.val(ret.Route.Id);
						$(
								".patrol-line .patrol-line-content #routelineEdit #routeline-name")
								.val(ret.Route.Name);
						$(
								".patrol-line .patrol-line-content #routelineEdit #routeline-description")
								.wysiwyg("setContent", ret.Route.Description);
					}, "json");
}

// pre-submit callback
function showRequest() {
	//alert("before submit");
}
// post-submit callback
function showResponse(re) {
	if (re.Result > 0) {
		// $("#alertInfo #content").html(re.Reason);
		// $("#alertInfo #content").show();
		//alert("NOK");
	} else {
		// location.href = re.Data
		//alert("OK");
	}
}

function delete_route(url) {
	alert(url);
}

function edit_route_meta(url) {
	alert(url);
}

function show_route_meta() {
	var content = $("#routeline-description").wysiwyg("getContent").val();
	alert(content);
}

$(document)
		.ready(
				function() {

					// Sidebar Accordion Menu:

					$("#main-nav li ul").hide(); // Hide all sub menus
					$("#main-nav li a.current").parent().find("ul")
							.slideToggle("slow"); // Slide down the current
													// menu item's sub menu

					$("#main-nav li a.nav-top-item").click(
							// When a top menu item is clicked...
							function() {
								$(this).parent().siblings().find(
										"a.nav-top-item")
										.removeClass("current");
								$(this).parent().siblings().find("ul").hide();
								$(this).parent().siblings().find("ul").slideUp(
										"normal"); // Slide up all sub menus
													// except the one clicked
								$(this).next().slideToggle("normal"); // Slide
																		// down
																		// the
																		// clicked
																		// sub
																		// menu
								$(this).addClass("current");
								window.location.href = (this.href); // Just open
																	// the link
																	// instead
																	// of a sub
																	// menu
								return false;
							});

					$("#main-nav li ul li a.nav-sub-item").click(
							// When a top menu item is clicked...
							function() {
								$(this).parent().siblings().find(
										"a.nav-sub-item")
										.removeClass("current"); // Remove
																	// "current"
																	// class
																	// from all
																	// tabs
								$(this).addClass("current"); // Add class
																// "current" to
																// clicked tab
								return false;
							});

					$("#main-nav li ul li a.patrol-line")
							.click(
									// When a top menu item is clicked...
									function() {
										$("#main-content").children(
												".content-box").children(
												".content-box-header")
												.children(".content-box-tabs")
												.hide();
										$("#main-content").children(
												".content-box").children(
												".content-box-content").hide();

										$(
												"#main-content .patrol-line .content-box-header .content-box-tabs")
												.show();
										$(
												"#main-content .patrol-line .content-box-content")
												.show();

										return false;
									});

					$("#main-nav li ul li a.patrol-point")
							.click(
									// When a top menu item is clicked...
									function() {
										$("#main-content").children(
												".content-box").children(
												".content-box-header")
												.children(".content-box-tabs")
												.hide();
										$("#main-content").children(
												".content-box").children(
												".content-box-content").hide();

										$(
												"#main-content .patrol-point .content-box-header .content-box-tabs")
												.show();
										$(
												"#main-content .patrol-point .content-box-content")
												.show();

										return false;
									});

					$("#main-nav li a.no-submenu").click( // When a menu item
															// with no sub menu
															// is clicked...
					function() {
						window.location.href = (this.href); // Just open the
															// link instead of a
															// sub menu
						return false;
					});

					// Sidebar Accordion Menu Hover Effect:

					$("#main-nav li .nav-top-item").hover(function() {
						$(this).stop().animate({
							paddingRight : "25px"
						}, 200);
					}, function() {
						$(this).stop().animate({
							paddingRight : "15px"
						});
					});

					// Minimize Content Box

					$(".content-box-header h3").css({
						"cursor" : "s-resize"
					}); // Give the h3 in Content Box Header a different cursor
					$(".closed-box .content-box-content").hide(); // Hide the
																	// content
																	// of the
																	// header if
																	// it has
																	// the class
																	// "closed"
					$(".closed-box .content-box-tabs").hide(); // Hide the tabs
																// in the header
																// if it has the
																// class
																// "closed"

					$(".content-box-header h3").click( // When the h3 is
														// clicked...
					function() {
						$(this).parent().next().toggle(); // Toggle the
															// Content Box
						$(this).parent().parent().toggleClass("closed-box"); // Toggle
																				// the
																				// class
																				// "closed-box"
																				// on
																				// the
																				// content
																				// box
						$(this).parent().find(".content-box-tabs").toggle(); // Toggle
																				// the
																				// tabs
					});

					// Content box tabs:

					$('.content-box .content-box-content div.tab-content')
							.hide(); // Hide the content divs
					$('ul.content-box-tabs li a.default-tab').addClass(
							'current'); // Add the class "current" to the
										// default tab
					$('.content-box-content div.default-tab').show(); // Show
																		// the
																		// div
																		// with
																		// class
																		// "default-tab"

					$('.content-box ul.content-box-tabs li a').click(
							// When a tab is clicked...
							function() {
								$(this).parent().siblings().find("a")
										.removeClass('current'); // Remove
																	// "current"
																	// class
																	// from all
																	// tabs
								$(this).addClass('current'); // Add class
																// "current" to
																// clicked tab
								var currentTab = $(this).attr('href'); // Set
																		// variable
																		// "currentTab"
																		// to
																		// the
																		// value
																		// of
																		// href
																		// of
																		// clicked
																		// tab
								$(currentTab).siblings().hide(); // Hide all
																	// content
																	// divs
								$(currentTab).show(); // Show the content div
														// with the id equal to
														// the id of clicked tab
								return false;
							});

					$('#routelineList .routeline .edit').click(function() {
						edit_route($(this).attr('href'));
						return false;
					});

					$('#routelineList .routeline .delete').click(function() {
						delete_route($(this).attr('href'));
						return false;
					});

					// 绑定表单提交事件处理器
					$('#routelineEdit #routeline').submit(function() {
						var options = {
								beforeSubmit : showRequest, // pre-submit callback
								success : showResponse, // post-submit callback
								dataType : 'json' // 'xml', 'script', or 'json'
													// (expected server response type)
							};

						$(this).ajaxSubmit(options);
						//submit_route($(this).attr('action'));
						// !!! Important !!!
						// 为了防止普通浏览器进行表单提交和产生页面导航（防止页面刷新？）返回false
						return false;
					});
					// Close button:

					$(".close").click(function() {
						$(this).parent().fadeTo(400, 0, function() { // Links
																		// with
																		// the
																		// class
																		// "close"
																		// will
																		// close
																		// parent
							$(this).slideUp(400);
						});
						return false;
					});

					// Alternating table rows:

					$('tbody tr:even').addClass("alt-row"); // Add class "alt-row" to even table rows

					// Check all checkboxes when the one in a table head is checked:

					$('.check-all').click(
							function() {
								$(this).parent().parent().parent().parent()
										.find("input[type='checkbox']").attr(
												'checked',
												$(this).is(':checked'));
							});

					// Initialise Facebox Modal window:

					$('a[rel*=modal]').facebox(); // Applies modal window to any link with attribute rel="modal"

					// Initialise jQuery WYSIWYG:

					$(".wysiwyg").wysiwyg(); // Applies WYSIWYG editor to any textarea with the class "wysiwyg"
					//wysiwyg_editor = $.wysiwyg({element: $(".wysiwyg")});

				});
