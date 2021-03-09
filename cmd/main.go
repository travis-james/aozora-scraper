package main

import (
	"AozoraScraper/scraper"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	ap := flag.String("ap", "https://www.aozora.gr.jp/index_pages/person11.html", "The url to the author's page")
	dn := flag.String("dn", "works", "The directory you want to save the author's work, too. Must be a new folder")
	flag.Parse()

	// Go to author page to get the HTML response.
	resp, err := http.Get(*ap)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Tokenize the author page into a map of URLs.
	mm, err := scraper.ParseAP(resp.Body)
	if err != nil {
		panic(err)
	}

	// Create a directory (directory name) to save the work to.
	err = os.Mkdir(*dn, 0755)
	if err != nil {
		panic(err)
	}

	// Followed the concurrency model found here: https://juliensalinas.com/en/how-to-speed-up-web-scraping-with-go-golang-concurrency/
	chFailed := make(chan string)
	chIsFinished := make(chan bool)

	// Now download all the zips from that map of links and save to the provided
	// directory name.
	for title, link := range mm {
		go scraper.DownloadWorks(*dn, title, link, chFailed, chIsFinished)
	}
	failedTitles := make([]string, 0)
	for i := 0; i < len(mm); {
		select {
		case title := <-chFailed:
			failedTitles = append(failedTitles, title)
		case <-chIsFinished:
			i++
		}
	}
	if len(failedTitles) > 0 {
		fmt.Println("The following failed.... ", failedTitles)
	}
	fmt.Println("Program finished.")
}
