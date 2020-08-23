package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://www.uta-net.com/search/?Keyword=REPLACE_HERE&x=26&y=21&Aselect=2&Bselect=3"

func main() {
	flag.Parse()
	searchKeyWord := flag.Args()
	searchQuery := ""
	for _, query := range(searchKeyWord) {
		searchQuery += url.QueryEscape(query + " ")
	}
	searchURL := strings.Replace(baseURL, "REPLACE_HERE", searchQuery[:len(searchQuery) - 2], 1)

	response, err := http.Get(searchURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		panic(err)
	}
	lyricURL, isExist := doc.Find("#ichiran > div.result_table.last > table > tbody > tr:nth-child(1) > td.side.td1 > a").Attr("href")
	if (!isExist) {
		panic("FUCK YOU")
	}

	lyricResponse, err := http.Get("https://www.uta-net.com/" + lyricURL)
	if err != nil {
		panic(err)
	}
	defer lyricResponse.Body.Close()

	lylicDoc, err := goquery.NewDocumentFromResponse(lyricResponse)
	if err != nil {
		panic(err)
	}
	lylic := lylicDoc.Find("#kashi_area").Text()
	fmt.Println(lylic)
}
