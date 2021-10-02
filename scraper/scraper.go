package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ScrapeDeNA() {
	url := "https://engineer.dena.com/"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	result := doc.Find(".article-list > div > h2 > a").First()
	articleLink, exist := result.Attr("href")
	if !exist {
		log.Println("error")
	}
	articleLink = url + articleLink
	fmt.Println(articleLink)
}
