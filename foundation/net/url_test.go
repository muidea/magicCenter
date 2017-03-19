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
	dir, name := SplitResetAPI(url)
	if dir != "/user" && name != "abc" {
		t.Errorf("SplitResetAPI failed, dir:%s,name:%s", dir, name)
	}

	url = "/user/abc/"
	dir, name = SplitResetAPI(url)
	if dir != "/user" && name != "abc" {
		t.Errorf("SplitResetAPI failed, dir:%s,name:%s", dir, name)
	}

	url = "/user/"
	dir, name = SplitResetAPI(url)
	if dir != "" && name != "user" {
		t.Errorf("SplitResetAPI failed, dir:%s,name:%s", dir, name)
	}
	url = "/user"
	dir, name = SplitResetAPI(url)
	if dir != "" && name != "user" {
		t.Errorf("SplitResetAPI failed, dir:%s,name:%s", dir, name)
	}
}
