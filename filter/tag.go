package filter

import (
	"fmt"

	"golang.org/x/net/html"
)

func FilterByTag(want, nodes []*html.Node) {
	for _, node := range nodes {
		fmt.Println(node)

	}

}
