package main

import (
	"AozoraScraper/scraper"
)

func main() {
	// Go to author page and get the HTML response.
	ap := "https://www.aozora.gr.jp/index_pages/person20.html"
	body := scraper.FetchHTML(ap)
	defer body.Close()

	// Tokenize the author page.
	mm := scraper.ParseAP(body)

	i := 0
	for key, val := range mm {
		// Get the response from a single work's link.
		body = scraper.FetchHTML(val)
		// Then on that web page, find the link to the zip of the work.
		zl := scraper.GetZipLink(body, val)
		fn := key + ".zip"
		scraper.DownloadFile(fn, zl)
		if i == 3 {
			break
		}
		i++
	}
}
