package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	url := "https://qiita.com/Toshinori_Hayashi/items/5b0a72dc64ced91717c0"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	rslt := doc.Find("h1").Text()
	fmt.Println(rslt)
}
