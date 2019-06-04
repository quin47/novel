package repository

import (
	"github.com/globalsign/mgo/bson"
	"novel/parse"
	"testing"
)

func Test_save(t *testing.T) {

	novel, urls := parse.ParseIndex("http://www.piaotian5.com/book/12651.html")
	novel.ID = bson.NewObjectId().Hex()
	Infos.Insert(novel)
	page, _ := parse.ParsePage(urls[0])
	page.BookId = novel.ID
	Chapters.Insert(page)

}
