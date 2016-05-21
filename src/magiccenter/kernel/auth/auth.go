package auth

import (
	"log"
	"net/http"
	"martini"
	"magiccenter/router"
    "magiccenter/session"
    "magiccenter/configuration"
	"magiccenter/kernel/account/model"
	"magiccenter/kernel/account/bll"
)

func AdminAuthVerify() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request) bool {
		authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
		if !found {
			panic("unexpected, can't fetch authorith id")
		}
		
		result := false
		session := session.GetSession(res, req)		
		user, found := session.GetOption(authId)
		if found {
			for _, gid := range user.(model.UserDetail).Groups {
				group, found := bll.QueryGroupById(gid)
				if found && group.AdminGroup() {
					result = true
				}
			}
		}
		
		return result
	}
}

func LoginAuthVerify() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request) bool {
		authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
		if !found {
			panic("unexpected, can't fetch authorith id")
		}
		
		session := session.GetSession(res, req)		
		_, found = session.GetOption(authId)
		return found
	}
}


func Authority() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		
		if !router.VerifyAuthority(res, req) {
			http.Redirect(res, req, "/", http.StatusFound)
			return 			
		}
		
		c.Next()
	}
}
