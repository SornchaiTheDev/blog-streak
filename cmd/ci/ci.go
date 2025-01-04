package main

import (
	"blogstreak/internal/services"
	"blogstreak/models"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"sync"

	"github.com/gosimple/slug"
)

func main() {

	markdownService := services.NewMarkdownService()
	fmt.Println("⏳ Updating Blogs metadata...")

	dir, err := os.ReadDir("./blogs")

	if err != nil {
		log.Fatal("Cannot read blogs/ dir")
	}

	var wg sync.WaitGroup
	metaChan := make(chan models.Metadata)
	for _, file := range dir {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := os.ReadFile("./blogs/" + file.Name())
			if err != nil {
				log.Fatalf("Cannot read %s file", file.Name())
			}
			metadata, err := markdownService.GetMetadata(data)
			if err != nil {
				log.Fatalf("Cannot get metadata of %s file", file.Name())
			}
			metaChan <- *metadata

		}()
	}

	go func() {
		wg.Wait()
		close(metaChan)
	}()

	var metadatas []models.Metadata

	for metadata := range metaChan {
		metadatas = append(metadatas, metadata)
	}

	dateCmp := func(a, b models.Metadata) int {
		return cmp.Compare(a.PublishedDate, b.PublishedDate) * -1
	}

	slices.SortFunc(metadatas, dateCmp)

	writeData := "{\n"

	currIndex := 0
	for _, metadata := range metadatas {
		generatedSlug := slug.Make(metadata.Title)
		writeData += fmt.Sprintf(`  "%s": "%s"`, metadata.PublishedDate, generatedSlug)
		if currIndex < len(metadatas)-1 {
			writeData += ",\n"
		}
		currIndex++
	}

	writeData += "\n}"

	err = os.WriteFile("./blogs_metadata.json", []byte(writeData), 0666)
	if err != nil {
		log.Fatal("Cannot write to ./blogs_metadata.json file")
	}

	fmt.Println("✨ Completed update blog posts")
}
