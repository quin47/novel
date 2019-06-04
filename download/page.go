package download

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHtml(url string) (string, error) {

	resp, err := http.Get(url)

	if err != nil {
		return "", errors.New(fmt.Sprintf("download page failed %v", url))
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	e, err := DetermineEncodingFromReader(bytes.NewReader(b))
	if err != nil {
		return "", errors.New("read html failed ")
	}
	r := transform.NewReader(bytes.NewReader(b), e.NewDecoder())
	all, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("reread bytes error :%v", err)
	}
	return string(all), nil

}

func DetermineEncodingFromReader(r io.Reader) (e encoding.Encoding, err error) {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return nil, err
	}
	e, _, _ = charset.DetermineEncoding(bytes, "")
	return e, err

}
