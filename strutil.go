package strutil

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

// Filters out all HTML tags from given string
func FilterHtmlTags(s string) string {
	var buffer bytes.Buffer

	dom := html.NewTokenizer(strings.NewReader(s))
	previousStartToken := dom.Token()
loopDomTest:
	for {
		tt := dom.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartToken = dom.Token()
		case tt == html.TextToken:
			if previousStartToken.Data == "script" {
				continue
			}
			textContent := strings.TrimSpace(html.UnescapeString(string(dom.Text())))
			if len(textContent) > 0 {
				buffer.WriteString(textContent)
			}
		}
	}

	return buffer.String()
}
