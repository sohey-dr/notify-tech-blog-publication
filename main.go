package main

import (
	"log"
	"notify-tech-blog-publication/scraper"
	"os"
	"sync"
	"time"

	"github.com/slack-go/slack"
)

// scrape target site
const siteNum int = 3

func run() {
	start := time.Now()
	articles := concurrentScraping()
	if articles != nil {
		err := notifySlack(articles)
		if err != nil {
			log.Println(err)
		}
	}
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

	go func() {
		defer wg.Done()
		cookpad, ok := scraper.ScrapeCookpad()
		if ok {
			articles = append(articles, cookpad)
			log.Println(cookpad)
		}
	}()

	wg.Wait()

	return articles
}

func notifySlack(articles []scraper.Article) error {
	text := "*公開された記事がありました！*\n"

	for _, article := range articles {
		text += "\n" + article.Company + ": <" + article.Url + "|" + article.Title + ">"
	}

	msg := slack.WebhookMessage{
		Text: text,
	}

	incomingWebHookURL := os.Getenv("NOTIFY_INCOMING_WEBHOOK")

	return slack.PostWebhook(incomingWebHookURL, &msg)
}

func main() {
	run()
}
