package scraper_test

import (
	"notify-tech-blog-publication/scraper"
	"testing"
)

func TestScraper(t *testing.T) {
	scraper.NewScraper("クックパッド", "https://techlife.cookpad.com/", "time", time.Now().Format("2006-01-02"), ".entry-title > a", true).Scrape()
}