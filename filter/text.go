package filter

import (
	"htx/utils"
	"strings"

	"golang.org/x/net/html"
)

func ExtractText(z *html.Tokenizer) (string, error) {

	isSkipMode := false
	omitTags := []string{"script", "style"}
	output := ""

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		t := z.Token()

		if t.Type.String() == "StartTag" && utils.ArrContains(t.Data, omitTags) {
			isSkipMode = true
		}

		if t.Type.String() == "EndTag" && utils.ArrContains(t.Data, omitTags) {
			isSkipMode = false
		}

		if t.Type.String() == "Text" && isSkipMode == false {
			text := strings.TrimSpace(t.Data)
			if len(text) > 0 {
				output += text + "\n"
			}
		}
	}

	return output, nil
}
