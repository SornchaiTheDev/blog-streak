package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosimple/slug"
)

func main() {
	name := flag.String("name", "", "The name of the blog post")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *help {
		fmt.Println("Usage: go run ./cmd/cli/cli.go --name Blog name")
		return
	}

	generatedSlug := slug.Make(*name)

	f, err := os.Create("./blogs/" + generatedSlug + ".md")
	if err != nil {
		log.Fatal("Cannot create the blog")
	}

	defer f.Close()

	layout := "02/01/2006"

	currentDate := time.Now().Format(layout)

	data := fmt.Sprintf("---\nTitle: %s\nPublishedDate: %s\n---", *name, currentDate)

	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal("Cannot write to file")
	}

	fmt.Printf(`✨ Create %s blog post successfully`, *name)
}
