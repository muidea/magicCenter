package dal

import "muidea.com/magicCenter/common/dbhelper"

var articleID int
var catalogID int
var linkID int
var mediaID int

func init() {
	dbhelper, _ := dbhelper.NewHelper()
	defer dbhelper.Release()

	articleID = loadArticleID(dbhelper)
	catalogID = loadCatalogID(dbhelper)
	linkID = loadLinkID(dbhelper)
	mediaID = loadMediaID(dbhelper)
}

func allocArticleID() int {
	articleID = articleID + 1
	return articleID
}

func allocCatalogID() int {
	catalogID = catalogID + 1
	return catalogID
}

func allocLinkID() int {
	linkID = linkID + 1
	return linkID
}

func allocMediaID() int {
	mediaID = mediaID + 1
	return mediaID
}
