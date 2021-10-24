package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Article struct {
	Company string
	Title   string
	Url     string
}

type Scraper interface {
	Scrape() (Article, bool)
}

type ScraperImpl struct {
	Target            string
	URL               string
	TimeTag           string
	Time              string
	TitleTag          string
	IsBaseURLContains bool
}

func NewScraper(target string, url string, timeTag string, time string, titleTag string, isBaseURLContains bool) Scraper {
	return &ScraperImpl{
		Target:            target,
		URL:               url,
		TimeTag:           timeTag,
		Time:              time,
		TitleTag:          titleTag,
		IsBaseURLContains: isBaseURLContains,
	}
}

func (s *ScraperImpl) Scrape() (Article, bool) {
	res, err := http.Get(s.URL)
	if err != nil {
		log.Println(err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	var article Article
	if strings.Contains(doc.Find(s.TimeTag).First().Text(), s.Time) {
		title := doc.Find(s.TitleTag).First()
		articleLink, exist := title.Attr("href")
		if !exist {
			log.Println("error")
		}

		url := articleLink
		if !s.IsBaseURLContains {
			url = s.URL + articleLink
		}

		article = Article{
			Company: s.Target,
			Title:   title.Text(),
			Url:     url,
		}

		return article, true
	}

	return article, false
}
