package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
* use http.Get(URL) to get bytes of file from url
* create an empty file using os.Create
* use io.Copy to copy downloaded bytes to file created.
*
* unzip -O shift-jis fire.zip
 */
func main() {
	// Go to author page.
	resp, err := http.Get("https://www.aozora.gr.jp/index_pages/person20.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Make a map of works, and their corresponding links.

	zip, err := os.Create("blep.html")
	if err != nil {
		panic(err)
	}
	defer zip.Close()

	b, err := io.Copy(zip, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("File size: ", b)
}
