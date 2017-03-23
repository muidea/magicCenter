package handler

type aclManager struct {
}

func (i *aclManager) addACLRoute(url string) bool {
	return true
}

func (i *aclManager) delACLRoute(url string) bool {
	return true
}

func (i *aclManager) adjustACLAuthGroup(url string, authGroup []int) bool {
	return true
}

func (i *aclManager) verifyAuthGroup(url string, authGroup []int) bool {
	return true
}
