package main

import (
	"github.com/notify-tech-blog-publication/scraper"
	"testing"
)

func TestScrapeDeNA(t *testing.T) {
	dena, ok := scraper.ScrapeDeNA()
	if dena.Company == "" && ok {
		t.Fatal("failed test")
	}
}

func TestScrapeZOZO(t *testing.T) {
	zozo, ok := scraper.ScrapeZOZO()
	if zozo.Company == "" && ok {
		t.Fatal("failed test")
	}
}

func TestScrapeCookpad(t *testing.T) {
	cookpad, ok := scraper.ScrapeCookpad()
	if cookpad.Company == "" && ok {
		t.Fatal("failed test")
	}
}
