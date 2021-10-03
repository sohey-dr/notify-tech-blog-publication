package main

import (
	"log"
	"notify-tech-blog-publication/scraper"
)

func main() {
	t, ok := scraper.ScrapeDeNA()
	if ok {
		log.Println(t)
	}
}
