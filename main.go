package main

import (
	"github.com/slack-go/slack"
	"log"
	"notify-tech-blog-publication/scraper"
	"os"
	"sync"
	"time"
)

func run() {
	articles := concurrentScraping(
		scraper.NewScraper("DeNA", "https://engineer.dena.com/", ".article-list", time.Now().Format("January 02, 2006"), "div > h2 > a", false).Scrape,
		scraper.NewScraper("ZOZO", "https://techblog.zozo.com/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
		// TODO: Slackで読み込み後に見れるため見逃しているがタイトルも取得するようにする
		scraper.NewScraper("Cyber Agent", "https://developers.cyberagent.co.jp/blog/", "time", time.Now().Format("2006/01/02"), ".card__time_title > a", true).Scrape,
		scraper.NewScraper("クックパッド", "https://techlife.cookpad.com/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
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
