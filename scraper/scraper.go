package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// FetchHTML takes in a url and returns the responses body.
func FetchHTML(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error when fetching URL: %v\n", err)
	}
	return resp.Body
}

// TokenizeAP takes a response body of an author's page then tokenizes it
// to build a map of titles with their links.
func TokenizeAP(body io.ReadCloser) map[string]string {
	retval := make(map[string]string)
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return retval
		}

		if tt == html.StartTagToken { // I'm looking for first occurence of <ol>
			token := z.Token()
			if token.Data == "ol" {
				for { // Inside <ol></ol>
					// Find me all <a>
					for tt != html.StartTagToken || token.Data != "a" {
						tt = z.Next()
						token = z.Token()
						// If the closing </ol> tag is found, we're done, return.
						if tt == html.EndTagToken && token.Data == "ol" {
							fmt.Println(token.Data)
							return retval
						}
					}
					// Get the link first. <a href=....
					link := token.Attr
					// Now move onto get the name of the work, <a>TITLE</a>.
					tt = z.Next()
					token = z.Token()
					title := token.Data
					// Add the title with it's corresponding link to the map.
					if len(link) > 0 {
						if link[0].Key == "href" {
							retval[title] = link[0].Val
						}
					}
				}
			}
		}
	}
}
