package dal

import (
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/common/initializer"
)

type idHolder struct {
	articleID int
	catalogID int
	linkID    int
	mediaID   int
	commentID int
}

func (s *idHolder) Handle() {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	s.articleID = loadArticleID(dbhelper)
	s.catalogID = loadCatalogID(dbhelper)
	s.linkID = loadLinkID(dbhelper)
	s.mediaID = loadMediaID(dbhelper)
}

func (s *idHolder) AllocArticleID() int {
	s.articleID++
	return s.articleID
}

func (s *idHolder) AllocCatalogID() int {
	s.catalogID++
	return s.catalogID
}

func (s *idHolder) AllocLinkID() int {
	s.linkID++
	return s.linkID
}

func (s *idHolder) AllocMediaID() int {
	s.mediaID++
	return s.mediaID
}

func (s *idHolder) AllocCommentID() int {
	s.commentID++
	return s.commentID
}

var holder = &idHolder{}

func init() {
	initializer.RegisterHandler(holder)
}

func allocArticleID() int {
	return holder.AllocArticleID()
}

func allocCatalogID() int {
	return holder.AllocCatalogID()
}

func allocLinkID() int {
	return holder.AllocLinkID()
}

func allocMediaID() int {
	return holder.AllocMediaID()
}

func allocCommentID() int {
	return holder.AllocCommentID()
}
