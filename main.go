package main

import (
	"flag"
	"fmt"
	"os"

	source "github.com/greenmato/news-summarizer/pkg"
	"github.com/greenmato/news-summarizer/pkg/article"
)

func main() {
	var month int
	var year int

	// Define command-line flags
	flag.IntVar(&month, "month", 1, "Month")
	flag.IntVar(&year, "year", 2025, "Year")

	// Add help text
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nA simple greeting CLI tool that demonstrates basic Go CLI functionality.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Parse command-line flags
	flag.Parse()

	source := source.GetSource(month, year)
	// fmt.Println(source)

	article := article.GenerateArticle(source)
	fmt.Println(article)
}
