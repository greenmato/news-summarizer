package source

import (
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func GetSource(month int, year int) string {
	html := getSourceHTML(month, year)

	return stripTags(html)
}

func getSourceUrl(month int, year int) string {
	return fmt.Sprintf("https://en.m.wikipedia.org/wiki/Portal:Current_events/%s_%d", time.Month(month), year)
}

func getSourceHTML(month int, year int) string {
	doc, err := htmlquery.LoadURL(getSourceUrl(month, year))
	if err != nil {
		panic(err)
	}

	parentNodes := htmlquery.Find(doc, "//*[@class='current-events']")

	var result string
	for _, node := range parentNodes {
		result = result + htmlquery.OutputHTML(node, true) + "\n"
	}

	return result
}

func stripTags(s string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(s))
	var result strings.Builder

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		if tokenType == html.TextToken {
			result.WriteString(tokenizer.Token().Data)
		}
	}

	return strings.TrimSpace(result.String())
}
