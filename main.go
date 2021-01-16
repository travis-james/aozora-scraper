package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var site string = `https://www.aozora.gr.jp`

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

	body := respToString(resp)

	// Strip the OL tags.
	bodyArr := strings.Split(body, "<ol>")
	bodyArr = strings.Split(bodyArr[1], "</ol>")
	// for i, v := range bodyArr {
	// 	fmt.Println(i, v)
	// }

	mymap := makeMap(bodyArr[0])
	i := 1
	for k, v := range mymap {
		fmt.Println(i, k, v)
		i++
	}

	// zip, err := os.Create("blep.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer zip.Close()

	// b, err := io.Copy(zip, resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File size: ", b)
}

func respToString(resp *http.Response) string {
	// Convert the response body to a string.
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodystr := string(bodyBytes)
		return bodystr
	}
	return ""
}

func makeMap(ps string) map[string]string {
	sakuhin := make(map[string]string)

	// Get the names of the works: <a href....>WN</a>
	regWN := regexp.MustCompile(`\">.*?\</a>`)
	resWN := regWN.FindAllString(ps, -1)
	// Get the corresponding links to those works: <a href="../LINK">
	regL := regexp.MustCompile(`\".*?\"`)
	resL := regL.FindAllString(ps, -1)
	for i := 0; i < len(resWN); i++ {
		// Sometimes there's an inline links to other people, ignore those links.
		if strings.Contains(resL[i], "person") {
			continue
		}

		wn := strings.Trim(resWN[i], "\">") // Get rid of tag elements.
		wn = strings.Trim(wn, "</a>")

		addr := strings.Trim(resL[i], "\"") // Strip trailing "
		addr = strings.Trim(addr, "..\"")   // Strip leading .."

		sakuhin[wn] = site + addr
	}
	return sakuhin
}
