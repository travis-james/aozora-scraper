package main

import (
	"AozoraScraper/scraper"
	"log"
	"os"
)

func main() {
	// Go to author page and get the HTML response.
	ap := "https://www.aozora.gr.jp/index_pages/person20.html"
	body := scraper.FetchHTML(ap)
	defer body.Close()

	// Tokenize the author page into a map of URLs.
	mm := scraper.ParseAP(body)

	// Create a directory (directory name) to save the work to.
	dn := "works/"
	err := os.Mkdir(dn, 0755)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for key, val := range mm {
		// Get the response from a single work's link.
		body = scraper.FetchHTML(val)
		// Then on that web page, find the link to the zip of the work.
		zl := scraper.GetZipLink(body, val)
		fn := dn + key + ".zip"
		scraper.DownloadFile(fn, zl)
		if i == 1 {
			break
		}
		i++
	}
}
