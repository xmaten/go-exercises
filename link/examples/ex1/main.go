package main

import (
	"fmt"
	"github.com/xmaten/link"
	"strings"
)

var exampleHtml = `
	<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link
to
	other page <span> some span  </span></a>
		<a href="/second-page">A link to second page</a>
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