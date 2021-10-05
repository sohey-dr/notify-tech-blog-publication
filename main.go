package main

import (
	"log"
	"notify-tech-blog-publication/scraper"
	"sync"
	"time"
)

const siteNum int = 2

func main() {
	start := time.Now()
	concurrent()
	end := time.Now()
	log.Printf("%f 秒時間がかかりました\n", (end.Sub(start)).Seconds())
}

func concurrentScraping() []scraper.Article {
	var articles []scraper.Article
	wg := &sync.WaitGroup{}
	wg.Add(siteNum)
	go func() {
		defer wg.Done()
		dena, ok := scraper.ScrapeDeNA()
		if ok {
			articles = append(articles, dena)
			log.Println(dena)
		}
	}()

	go func() {
		defer wg.Done()
		zozo, ok := scraper.ScrapeZOZO()
		if ok {
			articles = append(articles, zozo)
			log.Println(zozo)
		}
	}()

	wg.Wait()

	return articles
}
