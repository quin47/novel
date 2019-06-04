package download

import "testing"

func Test_GetHtml(t *testing.T) {

	html, _ := GetHtml("http://www.piaotian5.com/book/134.html")

	t.Logf("html value %v", html)
}
