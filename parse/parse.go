package parse

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"log"
	"novel/download"
	"strings"
	"time"
)

type ParsedResult struct {
	ID          string `json:"_id" bson:"_id,omitempty"`
	ChapterName string `json:"chapter_name"`
	Content     string `json:"content"`
	ChapterNum  int    `json:"chapter_num"`
	BookId      string `json:"book_id" bson:"book_id"`
}
type Novel struct {
	ID     string `json:"_id" bson:"_id,omitempty"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Class  string `json:"class"`
	Status string `json:"status"`
}

func ParseIndex(bookurl string) (Novel, []string) {

	var novel Novel
	urls := make([]string, 0, 30000)
	html, _ := download.GetHtml(bookurl)
	if html == "" {
		log.Printf("read page failed : %v", bookurl)
	}

	node, e := htmlquery.Parse(strings.NewReader(html))
	if e != nil {
		log.Printf("query html failed %v", e)
	}

	bookInfo := htmlquery.Find(node, "//div[@id='info']")
	bookName := htmlquery.Find(bookInfo[0], "//h1")
	infos := htmlquery.Find(bookInfo[0], "//p")
	novel.Name = bookName[0].FirstChild.Data
	novel.Author = strings.Replace(infos[0].FirstChild.Data, "    ", "", -1)
	novel.Class = strings.Replace(infos[2].FirstChild.Data, "    ", "", -1)
	novel.Status = strings.Replace(infos[1].FirstChild.Data, "    ", "", -1)
	log.Printf("get a novel :%v", novel)

	index := htmlquery.Find(node, "//dd/a@href")
	if len(index) >= 6 {
		index = index[6:]
	}
	for _, item := range index {
		urls = append(urls, "http://www.piaotian5.com"+item.Attr[0].Val)
	}

	return novel, urls
}

func ParsePage(url string) (ParsedResult, error) {
	var parsedRessut ParsedResult
	html, err := download.GetHtml(url)

	if err != nil {

		log.Printf("download page failed %v", url)
		for i := 1; i < 4; i++ {
			html, err = download.GetHtml(url)
			time.Sleep(time.Second * 3)
			log.Printf("retry  %v:download page:%v", i, url)
			if err == nil {
				break
			}
		}
	}

	if err == nil {
		log.Printf("download page failed after retries:  %v", url)
	}

	reader := strings.NewReader(html)
	node, e := htmlquery.Parse(reader)
	if e != nil {
		log.Printf("fetch page error: %v", e)
	}

	chapterName := htmlquery.Find(node, "//div[@class='content']/h1")
	if len(chapterName) != 1 {
		return parsedRessut, errors.New("get title failed")
	}

	parsedRessut.ChapterName = chapterName[0].FirstChild.Data

	content := htmlquery.Find(node, "//div[@id='content']")

	if len(content) < 1 {
		return parsedRessut, errors.New("get content failed")
	}

	/**
	remove last four line  (ads)
	*/
	content[0].RemoveChild(content[0].LastChild)
	content[0].RemoveChild(content[0].LastChild)

	content[0].RemoveChild(content[0].LastChild)

	content[0].RemoveChild(content[0].LastChild)

	text := htmlquery.InnerText(content[0])

	/**
	remove empty line
	*/
	text = strings.Replace(text, "\n\n", "\n", -1)

	parsedRessut.Content = text

	return parsedRessut, nil
}
