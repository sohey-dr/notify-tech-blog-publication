package main

import (
	"log"
	"notify-tech-blog-publication/scraper"
)

func main() {
	dena, ok := scraper.ScrapeDeNA()
	if ok {
		log.Println(dena)
	}

	zozo, ok := scraper.ScrapeZOZO()
	if ok {
		log.Println(zozo)
	}
}
