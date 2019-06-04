package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"log"
	"novel/parse"
	"novel/repository"
	"sync"
	"testing"
)

func Test_ANovel(t *testing.T) {
	var group = sync.WaitGroup{}

	novel, urls := parse.ParseIndex("http://www.piaotian5.com/book/12269.html")

	novel.ID = bson.NewObjectId().Hex()
	repository.Infos.Insert(novel)
	group.Add(len(urls))

	for i, v := range urls {
		go func(pageI int, url string) {
			page, er := parse.ParsePage(url)
			if er != nil {
				log.Printf("got an error %v", er)
			}
			page.BookId = novel.ID
			page.ChapterNum = pageI + 1
			repository.Chapters.Insert(page)
			log.Printf("finised %v: %v", page.ChapterNum, page.ChapterName)
			group.Done()

		}(i, v)
	}

	group.Wait()
	assert.NotEmpty(t, novel.Name)
	t.Logf("finished book %v", novel.Name)
}
