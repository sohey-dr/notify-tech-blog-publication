package main

import (
	"github.com/slack-go/slack"
	"log"
	"notify-tech-blog-publication/scraper"
	"os"
	"sync"
)

func run() {
	articles := concurrentScraping(
		scraper.ScrapeDeNA,
		scraper.ScrapeZOZO,
		scraper.NewScraper("クックパッド", "https://techlife.cookpad.com/", "time", "-10", "time", "15", ".entry-title > a", false).Scrape,
	)
	if len(articles) != 0 {
		err := notifySlack(articles)
		if err != nil {
			log.Println(err)
		}
	}
}

func concurrentScraping(fs ...func() (scraper.Article, bool)) []scraper.Article {
	var (
		articles = make([]scraper.Article, 0, len(fs))
		wg       = sync.WaitGroup{}
	)

	for _, f := range fs {
		wg.Add(1)

		f := f
		go func() {
			defer wg.Done()
			if art, ok := f(); ok {
				articles = append(articles, art)
				log.Println(art)
			}
		}()
	}
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
