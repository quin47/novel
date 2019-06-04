package parse

import (
	"testing"
	"unicode/utf8"
)

func Test_ParseIndex(t *testing.T) {

	novel, strings := ParseIndex("http://www.piaotian5.com/book/12651.html")

	t.Logf("test book :%v, urls: %v ", novel, len(strings))

}

func Test_ParsePAge(t *testing.T) {

	page, _ := ParsePage("http://www.piaotian5.com/book/12651/7886363.html")

	t.Logf("test content :%v, word_count: %v ", page.Content, utf8.RuneCountInString(page.Content))

}
