package parse

import (
	"github.com/antchfx/htmlquery"
	"log"
	"novel/download"
	"strings"
)

type ParsedResult struct {
	Content string `json:"content"`
	ChapterNum int `json:"chapter_num"`
	Book_id int `json:"book_id"`
}
type Novel struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	Class     string `json:"class"`
	Status    string `json:"status"`
	id int `json:"id"`
}

func ParseIndex(bookurl string) (Novel, []string) {

	var novel Novel
	urls := make([]string, 30000)
	html := download.GetHtml(bookurl)
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
	novel.Author = strings.Replace(infos[0].FirstChild.Data,"    ","",-1)
	novel.Class = strings.Replace( infos[2].FirstChild.Data,"    ","",-1)
	novel.Status =strings.Replace(infos[1].FirstChild.Data,"    ","",-1)
	log.Printf("get a novel :%v", novel)

	index := htmlquery.Find(node, "//dd/a@href")
	if len(index)>=6 {
		index=index[6:]
	}
	for _,item :=range index{
		urls= append(urls, "http://www.piaotian5.com"+item.Attr[0].Val)
	}

	return novel, urls
}

func ParsePage(url string) ParsedResult {
	var parsedRessut ParsedResult

	node, e := htmlquery.LoadURL(url)
	if e != nil {
		log.Printf("fetch page error: %v",e)
	}

	content := htmlquery.Find(node, "//div[@id='content']")

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

	parsedRessut.Content= text

	return parsedRessut
}
