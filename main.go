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

func main() {
	// Go to author page and get the HTML response.
	body := fetchHTML("https://www.aozora.gr.jp/index_pages/person20.html")
	defer body.Close()

	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()

		// If an error token, it's the end of the file, stop computing.
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			// Otherwise the html is malformed, quit it.
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}

		//process the token according to the token type...
		//if this is a start tag token...
		if tokenType == html.StartTagToken {
			//get the token
			token := tokenizer.Token()
			//if the name of the element is "title"
			if "title" == token.Data {
				//the next token should be the page title
				tokenType = tokenizer.Next()
				//just make sure it's actually a text token
				if tokenType == html.TextToken {
					//report the page title and break out of the loop
					fmt.Println(tokenizer.Token().Data)
					break
				}
			}
		}
	}
}
