package user

import (
	"webcenter/datamanager"
)

type User struct {
	datamanager.User
}

func (this *User)Valid() bool {
	return this.Id > 0 && len(this.Name) > 0 && len(this.Password) > 0
}

func (this *User)ValidPassword(password string) bool {
	return this.Password == password
}

func FindUserByEmail(email string) (User, bool) {
	user := User{}
	
    userManager := datamanager.GetUserManager()
    userInfo, found:= userManager.FindUserByEMail(email)
    if found {
    	user.Id = userInfo.Id
    	user.Name = userInfo.Name
    	user.Password = userInfo.Password
    	user.Email = userInfo.Email
    	user.Group = user.Group
    }
    
    return user,found
}

