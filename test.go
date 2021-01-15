package main

import (
	"fmt"
	"strings"
)

func main() {
	a := `yo ablsf
	sfjksj <ol>
	<li><a href="go here">ようこそ</a></li>
	<li><a href="go there">よ</a></li>
	<li><a href="go over there">そ</a></li>
	</ol>
	<h2>HERE WE GO</h2>`
	if strings.Contains(a, "<ol>") {
		// Get rid of everything before the ordered list.
		b := strings.Split(a, "<ol>")
		// Delete everything after the ordered list.
		c := strings.Split(b[1], "</ol>")
		// Now creaete a list (tokenize) of each list element.
		arr := strings.Split(c[0], "</li>")
		// Now tokenize on each " character.
		//links := []string{}
		for i := 0; i < len(arr)-1; i++ {
			//temp := strings.Split(arr[i], "\"") // Get the link.
			//temp := strings.Split(arr[i])
			fmt.Println(temp)
		}
		// for i, v := range links {
		// 	fmt.Println(i, v)
		// }
		// for _, val := range arr {
		// 	links := strings.Split(val, "\"")
		// }
		//fmt.Println(arr[3])
		//links := make(map[string]string)

	}

}
