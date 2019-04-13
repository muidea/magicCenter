package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/muidea/magicCenter/common/dbhelper"
	"github.com/muidea/magicCenter/module/modules/authority/dal"
	"github.com/muidea/magicCommon/model"
)

var routeReg1 = regexp.MustCompile(`:[^/#?()\.\\]+`)
var routeReg2 = regexp.MustCompile(`\*\*`)

type aclItem struct {
	regex *regexp.Regexp
	acl   model.ACL
}

type aclItemList []*aclItem

type aclHandler struct {
	aclListMap map[string]aclItemList
}

func createACLHandler() *aclHandler {
	return &aclHandler{aclListMap: make(map[string]aclItemList)}
}

func (s *aclHandler) loadACL() {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	acls := dal.QueryAllACL(dbhelper)
	for _, val := range acls {
		pattern := routeReg1.ReplaceAllStringFunc(val.URL, func(m string) string {
			return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
		})
		var index int
		pattern = routeReg2.ReplaceAllStringFunc(pattern, func(m string) string {
			index++
			return fmt.Sprintf(`(?P<_%d>[^#?]*)`, index)
		})
		pattern += `\/?`
		regex := regexp.MustCompile(pattern)

		item := &aclItem{regex: regex, acl: val}

		itemList, ok := s.aclListMap[val.Method]
		if !ok {
			itemList = aclItemList{}
			itemList = append(itemList, item)
			s.aclListMap[val.Method] = itemList
		} else {
			itemList = append(itemList, item)
			s.aclListMap[val.Method] = itemList
		}
	}
}

func (s *aclHandler) filterACL(req *http.Request) (model.ACL, bool) {
	var acl model.ACL
	foundFlag := false

	itemList, ok := s.aclListMap[req.Method]
	if ok {
		url := req.URL.Path
		for _, val := range itemList {
			matches := val.regex.FindStringSubmatch(url)
			if len(matches) > 0 && matches[0] == url {
				foundFlag = true
				acl = val.acl
				break
			}
		}
	}

	return acl, foundFlag
}
