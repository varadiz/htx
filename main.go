package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)
	z := html.NewTokenizer(r)

	indentPosition := 0
	indentSize := 1
	isSkipMode := false
	omitTags := []string{"script", "style"}

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return
		}
		t := z.Token()

		if t.Type.String() == "StartTag" {
			indentPosition += indentSize
		}

		if t.Type.String() == "StartTag" && contains(t.Data, omitTags) {
			isSkipMode = true
		}

		if t.Type.String() == "EndTag" && contains(t.Data, omitTags) {
			isSkipMode = false
		}

		if t.Type.String() == "Text" && isSkipMode == false {
			text := strings.TrimSpace(t.Data)
			if len(text) > 0 {
				fmt.Println(strings.Repeat(" ", indentPosition) + text)
			}
		}

		if t.Type.String() == "EndTag" {
			indentPosition -= indentSize
		}
	}
}
