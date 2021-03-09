package scraper

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Default URL for creating links to works.
var defURL string = "https://www.aozora.gr.jp"

// ParseAP takes a response body of an author's page then parses it
// to build a map of titles with their links.
func ParseAP(body io.ReadCloser) (map[string]string, error) {
	retval := make(map[string]string)
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return nil, errors.New("Error in ParseAP html.ErrorToken")
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
							return retval, nil
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

// GetZipLink takes a response (body) from a url, and parses the HTML elements
// to find the zip link for the author's work. The url to the zip is
// returned as a string. The baseURL parameter is used to build the returned
// url.
func GetZipLink(body io.ReadCloser, baseURL string) (string, error) {
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return "", errors.New("func GetZipLink failure in parsing HTML")
		}
		// I'm looking for <a>
		if tt == html.StartTagToken {
			token := z.Token()
			if token.Data == "a" {
				link := token.Attr
				if len(link) > 0 {
					if link[0].Key == "href" {
						zl := link[0].Val
						if strings.Contains(zl, "zip") {
							zl = strings.TrimLeft(zl, "./")
							bURL := strings.Split(baseURL, "/")
							// Hideous way of doing things, but it works. I just want to remove the /cardxxx.html
							// part at the end from https://www.aozora.gr.jp/cards/000020/cardxxx.html
							// and put the zl at the end.
							url := bURL[0] + "//" + bURL[2] + "/" + bURL[3] + "/" + bURL[4] + "/" + zl
							return url, nil
						}
					}
				}
			}
		}
	}
}

// DownloadFile takes a url and saves the response to a file (filename fn).
// Returns a nil error on success.
// This is from: https://golangcode.com/download-a-file-from-a-url/
// I google'd how to download a file in Golang, and this is exactly
// what I needed :)
func DownloadFile(fn string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadWorks takes a directory name (dn), the title of a work, then a link for that work.
// Two channels are passed: chIsFinished to keep track of when a routine finishes.
// chFailed allows passing of strings to give an error and the title of the work that failed.
func DownloadWorks(dn, title, link string, chFailed chan string, chIsFinished chan bool) {
	defer func() {
		chIsFinished <- true
	}() // Signal channel is done on exit.

	// Get the response from a single work's link.
	resp, err := http.Get(link)
	if err != nil {
		chFailed <- title + " : " + err.Error()
		return
	}

	// Then on that web page, find the link to the zip of the work.
	zl, err := GetZipLink(resp.Body, link)
	resp.Body.Close()
	if err != nil {
		chFailed <- title + " : " + err.Error()
		return
	}

	fn := dn + "/" + title + ".zip"
	err = DownloadFile(fn, zl)
	if err != nil {
		chFailed <- title + " : " + err.Error()
		return
	}
}
