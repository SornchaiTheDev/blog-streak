package main

import (
	"blogstreak/internal/services"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosimple/slug"
)

func isExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func main() {
	name := flag.String("name", "", "The name of the blog post")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *help {
		fmt.Println("Usage: go run ./cmd/cli/cli.go --name Blog name")
		return
	}

	generatedSlug := slug.Make(*name)

	fileName := "./blogs/" + generatedSlug + ".md"

	if isExists(fileName) {
		fmt.Println("❌ The blog is already exists")
		os.Exit(0)
	}

	f, err := os.Create(fileName)
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

	streakService := services.NewStreakService()

	streakService.Update()

	fmt.Printf(`✨ Create %s blog post successfully`, *name)
}
