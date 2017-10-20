package dal

import (
	"testing"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
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
	user.Groups = []int{}
	user.Status = 0

	user2 := user
	user2.Account = "t2"

	groups := []int{}

	_, ret := CreateUser(helper, "user", "test", groups)
	if !ret {
		t.Error("CreateUser failed")
		return
	}

	_, ret = CreateUser(helper, "user2", "test", groups)
	if !ret {
		t.Error("CreateUser failed")
		return
	}

	usr, found := QueryUserByAccount(helper, "test", "test")
	if !found {
		t.Errorf("QueryUserByAccount failed, account:%s", "test")
		return
	}

	usr.Name = "testNick"

	_, ret = SaveUser(helper, usr)
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
