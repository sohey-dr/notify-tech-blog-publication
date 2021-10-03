package main

import (
	"fmt"
	"notify-tech-blog-publication/scraper"
)

func main() {
	t, ok := scraper.ScrapeDeNA()
	if ok {
		fmt.Println(t)
	}
}
