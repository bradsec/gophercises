package main

import (
	"fmt"
	"strings"

	"github.com/bradsec/gophercises/link"
)

var exampleHtml = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page-one">A link to page one</a>
	<a href="/other-page-two">A link 
	to 
	page two</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
