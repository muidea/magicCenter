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

func TestSplitParam(t *testing.T) {
	url := "key1=value1&key2=value2"
	ret := SplitParam(url)
	if len(ret) != 2 {
		t.Errorf("SplitParam failed, ret:%d", len(ret))
	}
	value, found := ret["key1"]
	if !found || value != "value1" {
		t.Error("SplitParam failed")
	}

	url = "key1=value1&key2"
	ret = SplitParam(url)
	if len(ret) != 1 {
		t.Errorf("SplitParam failed, ret:%d", len(ret))
	}
	value, found = ret["key1"]
	if !found || value != "value1" {
		t.Error("SplitParam failed")
	}

	url = "key1=value1&key2="
	ret = SplitParam(url)
	if len(ret) != 1 {
		t.Errorf("SplitParam failed, ret:%d", len(ret))
	}
	value, found = ret["key1"]
	if !found || value != "value1" {
		t.Error("SplitParam failed")
	}
	value, found = ret["key2"]
	if found {
		t.Error("SplitParam failed")
	}
}

func TestParseRestAPIUrl(t *testing.T) {
	url := "/user/abc/?key1=value1&key2=value2"
	id, param, ret := ParseRestAPIUrl(url)
	if !ret || id != "abc" || len(param) != 2 {
		t.Errorf("ParseRestAPIUrl failed, ret=%d, id:%s, param len:%d", ret, id, len(param))
	}
	if id != "abc" {
		t.Errorf("ParseRestAPIUrl failed, id:%s", id)
	}

	v, b := param["key2"]
	if !b || v != "value2" {
		t.Error("ParseResetAPIUrl failed")
	}
}
