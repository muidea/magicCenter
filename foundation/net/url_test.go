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
