package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"time"
)

type Article struct {
	Title string
	Url   string
}

func ScrapeDeNA() *Article {
	url := "https://engineer.dena.com/"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	t := time.Now()
	latestArticleContents := doc.Find(".article-list").First().Text()
	if strings.Contains(latestArticleContents, string(t.Month())) && strings.Contains(latestArticleContents, string(t.Day())) {
		result := doc.Find(".article-list > div > h2 > a").First()
		articleLink, exist := result.Attr("href")
		if !exist {
			log.Println("error")
		}

		article = &Article{
			Title: title.Text(),
			Url:   url + articleLink,
		}
	}

	return article
}
