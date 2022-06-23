package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func toInt64(strVal string) string {
	rex := regexp.MustCompile("[0-9]+")
	strVal = rex.FindString(strVal)

	return strVal
}
func getRanking(userpage_url string) string {
	res, err := http.Get(userpage_url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(body)
	html := buf.String()

	//要らない文字を消す
	idx := strings.Index(html, `<div class="tableItem rank"><span>`)
	rank := toInt64(html[idx : idx+50])

	return rank
}

func main() {
	ranking := getRanking("https://ibispaint.com/artist1/1086978798250123264/")
	fmt.Printf("%#v 位です", ranking)
}
