package scraper

import (
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Default URL for creating links to works.
var defURL string = "https://www.aozora.gr.jp"

// FetchHTML takes in a url and returns the responses body.
func FetchHTML(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error when fetching URL: %v\n", err)
	}
	return resp.Body
}

// ParseAP takes a response body of an author's page then parses it
// to build a map of titles with their links.
func ParseAP(body io.ReadCloser) map[string]string {
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
					// I want to record all <a href=""></a> so begin by finding
					// individual <a>
					for tt != html.StartTagToken || token.Data != "a" {
						tt = z.Next()
						token = z.Token()
						// If the closing </ol> tag is found, we're done, return.
						if tt == html.EndTagToken && token.Data == "ol" {
							return retval
						}
					}
					// Get the link first. <a href=....
					link := token.Attr
					// Now move onto get the name of the work, <a>TITLE</a>.
					tt = z.Next()
					token = z.Token()
					title := token.Data
					// Add the title with it's corresponding link to the map, if it exists.
					if len(link) > 0 {
						if link[0].Key == "href" {
							wl := link[0].Val
							// The webiste has inline link elements that are not works.
							// Do NOT add those.
							if !(strings.Contains(wl, "person")) {
								wl = defURL + strings.TrimLeft(wl, "..")
								retval[title] = wl
							}
						}
					}
				}
			}
		}
	}
}

// GetZipLink takes a response from a url, and parses the HTML elements
// to find the zip link for the author's work. The url to the zip is
// returned as a string. The 'url' parameter is used to build the returned
// url.
func GetZipLink(body io.ReadCloser, url string) string {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			log.Fatal("GetZipLink failure in parsing HTML.")
		}
		// I'm looking for <a>
		if tt == html.StartTagToken {
			token := z.Token()
			if token.Data == "a" {
				link := token.Attr
				if len(link) > 0 {
					if link[0].Key == "href" {
						zl := link[0].Val
						return zl
					}
				}
			}
		}
	}
}
