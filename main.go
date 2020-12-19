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
 */
func main() {
	resp, err := http.Get("https://i.pximg.net/img-master/img/2019/01/01/10/36/56/72426894_p0_master1200.jpg")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	img, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	b, err := io.Copy(img, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("File size: ", b)
}
