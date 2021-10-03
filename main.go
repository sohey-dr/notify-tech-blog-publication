package main

import (
	"log"
	"notify-tech-blog-publication/scraper"
)

func main() {
	start := time.Now()
	concurrent()
	end := time.Now()
	log.Printf("%f 秒時間がかかりました\n", (end.Sub(start)).Seconds())
}

func concurrent() {
	wg := &sync.WaitGroup{}
	wg.Add(siteNum)
	go func() {
		defer wg.Done()
		dena, ok := scraper.ScrapeDeNA()
		if ok {
			log.Println(dena)
		}
	}()

	go func() {
		defer wg.Done()
		zozo, ok := scraper.ScrapeZOZO()
		if ok {
			log.Println(zozo)
		}
	}()

	wg.Wait()
}
