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
	mm := scraper.ParseAP(body)

	for _, val := range mm {
		body = scraper.FetchHTML(val)
		fmt.Println(scraper.GetZipLink(body, val))
		break
	}
}
