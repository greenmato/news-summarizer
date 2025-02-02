package main

import (
	"flag"
	"fmt"

	"github.com/greenmato/news-summarizer/pkg/article"
	"github.com/greenmato/news-summarizer/pkg/source"
)

func main() {
	var month int
	var year int

	// Define command-line flags
	flag.IntVar(&month, "month", 1, "Month")
	flag.IntVar(&year, "year", 2025, "Year")

	// Parse command-line flags
	flag.Parse()

	source := source.GetSource(month, year)
	// fmt.Println(source)

	article := article.GenerateArticle(source)
	fmt.Println(article)
}
