package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var site string = `https://www.aozora.gr.jp`

// fetchHTML takes in a url and returns the responses body.
func fetchHTML(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error when fetching URL: %v\n", err)
	}
	return resp.Body
}

func tokenize(body io.ReadCloser) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}

		if tt == html.StartTagToken { // I'm looking for first occurence of <ol>
			if z.Token().Data == "ol" {
				z.Next()
				z.Next()
				z.Next()
				z.Next()
				fmt.Println(z.Token().Data)
				fmt.Println("TEST")
			}
			return
		}
	}
}

func main() {
	// Go to author page and get the HTML response.
	body := fetchHTML("https://www.aozora.gr.jp/index_pages/person20.html")
	defer body.Close()

	tokenize(body)
}
