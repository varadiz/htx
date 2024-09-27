package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"htx/filter"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	z := html.NewTokenizer(r)

	shouldExtractText := flag.Bool("text", false, "Flag to extract text")
	shouldExtractPrettyText := flag.Bool("prettytext", false, "Flag to extract pretty text")

	flag.Parse()

	// default should extract text
	if *shouldExtractText || (*shouldExtractText == false && *shouldExtractPrettyText == false) {
		text, err := filter.ExtractText(z)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stdout, text)
	}

	if *shouldExtractPrettyText {
		text, err := filter.ExtractPrettyText(z)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stdout, text)
	}
}
