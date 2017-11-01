package net

import "testing"

func TestJoinURL(t *testing.T) {
	pre := "aa"
	sub := "bb"
	ret := JoinURL(pre, sub)
	if ret != "aa/bb" {
		t.Error("JoinURL unexpect, ret:" + ret)
	}

	pre = "aa/"
	sub = "bb"
	ret = JoinURL(pre, sub)
	if ret != "aa/bb" {
		t.Error("JoinURL unexpect, ret:" + ret)
	}

	pre = "/aa//"
	sub = "bb"
	ret = JoinURL(pre, sub)
	if ret != "/aa/bb" {
		t.Error("JoinURL unexpect, ret:" + ret)
	}
	pre = "/aa/"
	sub = "/bb"
	ret = JoinURL(pre, sub)
	if ret != "/aa/bb" {
		t.Error("JoinURL unexpect, ret:" + ret)
	}
	pre = "/aa/"
	sub = "/bb/"
	ret = JoinURL(pre, sub)
	if ret != "/aa/bb/" {
		t.Error("JoinURL unexpect, ret:" + ret)
	}
}

func TestParseRestAPIUrl(t *testing.T) {
	url := "/user/abc"
	dir, name := SplitRESTAPI(url)
	if dir != "/user/" && name != "abc" {
		t.Errorf("SplitRESTAPI failed, dir:%s,name:%s", dir, name)
	}

	url = "/user/abc/"
	dir, name = SplitRESTAPI(url)
	if dir != "/user/abc/" && name != "" {
		t.Errorf("SplitRESTAPI failed, dir:%s,name:%s", dir, name)
	}

	url = "/user/"
	dir, name = SplitRESTAPI(url)
	if dir != "/user/" && name != "" {
		t.Errorf("SplitRESTAPI failed, dir:%s,name:%s", dir, name)
	}
	url = "/user"
	dir, name = SplitRESTAPI(url)
	if dir != "/" && name != "user" {
		t.Errorf("SplitRESTAPI failed, dir:%s,name:%s", dir, name)
	}
}

func TestFormatRoutePattern(t *testing.T) {
	url := "/user/"
	id := "abc"
	pattern := FormatRoutePattern(url, id)
	if pattern != "/user/:id" {
		t.Errorf("FormatRoutePattern failed, url:%s, id:%s", url, id)
	}

	url = "/user/abc"
	id = "ef"
	pattern = FormatRoutePattern(url, id)
	if pattern != "/user/abc/:id" {
		t.Errorf("FormatRoutePattern failed, url:%s, id:%s", url, id)
	}

	url = "/user/abc"
	id = ""
	pattern = FormatRoutePattern(url, id)
	if pattern != "/user/abc/" {
		t.Errorf("FormatRoutePattern failed, url:%s, id:%s", url, id)
	}

	url = "/user/"
	id = ""
	pattern = FormatRoutePattern(url, id)
	if pattern != "/user/" {
		t.Errorf("FormatRoutePattern failed, url:%s, id:%s", url, id)
	}
}
