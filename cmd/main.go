package main

import (
	"AozoraScraper/scraper"
	"fmt"
)

func main() {
	// Go to author page and get the HTML response.
	ap := "https://www.aozora.gr.jp/index_pages/person20.html"
	body := scraper.FetchHTML(ap)
	defer body.Close()

	// Tokenize the author page.
	mm := scraper.TokenizeAP(body)
	i := 1
	for key, val := range mm {
		fmt.Println(i, key, val)
		i++
	}
}
