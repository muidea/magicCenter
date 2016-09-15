package dal

import (
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
	"testing"
)

func TestUser(t *testing.T) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	user := model.UserDetail{}
	user.Account = "test"
	user.Name = "nickName"
	user.Email = "test@126.com"
	user.Status = 0
	user.Groups = append(user.Groups, 10)
	user.Groups = append(user.Groups, 11)
	user.Groups = append(user.Groups, 12)
	user.Groups = append(user.Groups, 13)

	user2 := user
	user2.Account = "t2"

	ret := CreateUser(helper, user, "test")
	if !ret {
		t.Error("CreateUser failed")
		return
	}

	ret = CreateUser(helper, user2, "test")
	if !ret {
		t.Error("CreateUser failed")
		return
	}

	usr, found := QueryUserByAccount(helper, "test")
	if !found {
		t.Errorf("QueryUserByAccount failed, account:%s", "test")
		return
	}

	usr.Name = "testNick"

	ret = SaveUser(helper, usr)
	if !ret {
		t.Errorf("SaveUser failed, id=%d", usr.ID)
		return
	}

	newUsr, found := QueryUserByID(helper, usr.ID)
	if !found {
		t.Errorf("QueryUserByID failed, id=%d", usr.ID)
		return
	}
	if newUsr.Name != "testNick" {
		t.Error("invalid user name")
		return
	}

	vU, ret := VerifyUserByAccount(helper, "test", "test")
	if !ret {
		t.Error("VerifyUserByAccount failed")
		return
	}
	if vU.Email != "test@126.com" {
		t.Error("VerifyUserByAccount failed")
		return
	}

	u1 := QueryUserByGroup(helper, 10)
	if len(u1) == 0 {
		t.Errorf("QueryUserByGroup failed, group=%d", 10)
		return
	}
	u2 := QueryUserByGroup(helper, 11)
	if len(u2) == 0 {
		t.Errorf("QueryUserByGroup failed, group=%d", 11)
		return
	}
	u3 := QueryUserByGroup(helper, 13)
	if len(u3) == 0 {
		t.Errorf("QueryUserByGroup failed, group=%d", 13)
		return
	}

	users := QueryAllUser(helper)
	if len(users) < 2 {
		t.Error("QueryAllUser failed")
		return
	}

	ret = DeleteUserByAccount(helper, "t2", "test")
	if !ret {
		t.Errorf("DeleteUserByAccount failed,%s", "t2")
		return
	}

	ret = DeleteUser(helper, newUsr.ID)
	if !ret {
		t.Errorf("DeleteUser failed, id=%d", newUsr.ID)
		return
	}

}
