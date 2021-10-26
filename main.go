package main

import (
	"github.com/slack-go/slack"
	"log"
	"notify-tech-blog-publication/scraper"
	"os"
	"strings"
	"sync"
	"time"
)

func run() {
	articles := concurrentScraping(
		scraper.NewScraper("DeNA", "https://engineer.dena.com/", ".article-list", time.Now().Format("January 02, 2006"), "div > h2 > a", false).Scrape,
		scraper.NewScraper("ZOZO", "https://techblog.zozo.com/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
		// TODO: Slackで読み込み後に見れるため見逃しているがタイトルも取得するようにする(CA, メルカリ, エルレカ, Yahoo)
		scraper.NewScraper("Cyber Agent", "https://developers.cyberagent.co.jp/blog/", "time", time.Now().Format("2006/01/02"), ".card__time_title > a", true).Scrape,
		scraper.NewScraper("メルカリ", "https://engineering.mercari.com/", "time", time.Now().Format("2006/01/02"), ".post-list__item > a", true).Scrape,
		scraper.NewScraper("Sansan", "https://buildersbox.corp-sansan.com/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
		scraper.NewScraper("マネーフォワード", "https://moneyforward.com/engineers_blog/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
		scraper.NewScraper("ラクスル", "https://tech.raksul.com/", ".c-posts__post__date", time.Now().Format("2006.01.02"), ".c-posts__post__info__title > a", true).Scrape,
		scraper.NewScraper("Yahoo!", "https://techblog.yahoo.co.jp/", ".date", time.Now().Format("2006.01.02"), ".panel-horizontal > a", true).Scrape,
		scraper.NewScraper("エルレカ", "https://medium.com/eureka-engineering/", "time", time.Now().Format("Jan 02"), ".postMetaInline > a", true).Scrape,
		scraper.NewScraper("Gunosy", "https://tech.gunosy.io/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
		scraper.NewScraper("Retty", "https://engineer.retty.me/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape,
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
		text += "\n" + article.Company + ": <" + article.Url + "|" + strings.ReplaceAll(article.Title, " ", "") + ">"
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
