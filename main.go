package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type jobResult struct {
	id       string
	location string
	title    string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=python"

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}
func getPage(pageNum int) {
	pageUrl := baseURL + "&recruitPage=" + strconv.Itoa(pageNum) + "&recruitSort=relation" + "&recruitPageCount=" + strconv.Itoa(10)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)
	doc, err2 := goquery.NewDocumentFromReader(res.Body)
	checkErr(err2)
	doc.Find(".item_recruit").Each(func(i int, s *goquery.Selection) {
		taget, _ := s.Attr("value")
		title, _ := s.Find(".job_tit>a").Attr("title")
		location := s.Find(".job_condition>span:first-child>a").Text()
		result := taget + " " + title + " " + location
		cleaning := cleanString(result)
		fmt.Println(cleaning)
	})
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(str), "")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status :: ", res.StatusCode)
	}
}
