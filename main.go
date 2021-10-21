package main

import (
	"github.com/notify-tech-blog-publication/scraper"
	"log"
	"os"
	"sync"
	"time"

	"github.com/slack-go/slack"
)

func Run() {
	start := time.Now()
	articles := concurrentScraping(
		scraper.ScrapeDeNA,
		scraper.ScrapeZOZO,
		scraper.ScrapeCookpad,
	)
	if len(articles) != 0 {
		err := notifySlack(articles)
		if err != nil {
			log.Println(err)
		}
	}
	end := time.Now()
	log.Printf("%f 秒時間がかかりました\n", (end.Sub(start)).Seconds())
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
	Run()
}

//type Scraper interface {
//	Scrape() (scraper.Article, bool)
//}
//
//type ScraperImpl struct {
//	URL string
//	// 取ってきたい tag 要素とか
//}
//
//func NewScpaper(url string) Scraper {
//	return &ScraperImpl{
//		URL: url,
//	}
//}
//
//func (s *ScraperImpl) Scrape() (scraper.Article, bool) {
//	// 処理
//}
