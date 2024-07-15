package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"htx/filter"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	z := html.NewTokenizer(r)

	// TODO: can only read once, so it's either html.Parse or html.NewTokenizer
	// doc, err := html.Parse(r)
	// if err != nil {
	// 	panic(err)
	// }
	// html.Render(os.Stdout, doc)

	text, err := filter.ExtractText(z)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, text)

}
