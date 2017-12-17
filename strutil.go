package strutil

import (
	"bytes"
	"fmt"
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

// Trims the suffix of a tree example : INPUT "tree+" AND SUFFIX "+" OUTPUT: "tree"
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// builds a string formed by the elements of the array of strings using separator
func MakeString(strings []string) string {
	return CombineStrings(strings, "")
}

// combine strings of array to a unique string split by separator
func CombineStrings(strings []string, separator string) string {
	if strings == nil || len(strings) == 0 {
		return ""
	}

	var buffer bytes.Buffer

	for i := 0; i < len(strings); i++ {
		buffer.WriteString(fmt.Sprintf("%s%s", strings[i], separator))
	}

	return TrimSuffix(buffer.String(), separator)
}
