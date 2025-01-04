package services

import (
	"blogstreak/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/gosimple/slug"
)

type metadataService struct {
	markdownService MarkdownService
}

func NewMetadataService(markdownService MarkdownService) *metadataService {
	return &metadataService{
		markdownService: markdownService,
	}
}

func createFile() ([]byte, error) {
	file, err := os.Create("./blogs_metadata.json")
	if err != nil {
		return nil, errors.New("Cannot create file blogs_metadata.json")
	}
	defer file.Close()

	_, err = file.WriteString("{}")
	if err != nil {
		return nil, errors.New("Cannot write to file blogs_metadata.json")
	}

	return []byte("{}"), nil
}

func getAll() (map[string]string, error) {
	data, err := os.ReadFile("./blogs_metadata.json")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			data, err = createFile()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Error reading file blogs_metadata.json")
		}
	}

	var metadatas map[string]string
	err = json.Unmarshal(data, &metadatas)
	if err != nil {
		return nil, errors.New("Cannot parse the metadata.")
	}

	return metadatas, nil
}

func getCurrentDate(currentSlug string) (*time.Time, error) {
	metadatas, err := getAll()
	if err != nil {
		return nil, err
	}

	currDate := "Not found"

	for date, slug := range metadatas {
		if slug == currentSlug {
			currDate = date
			break
		}
	}

	layout := "02/01/2006"
	currTime, err := time.Parse(layout, currDate)
	if err != nil {
		return nil, errors.New("Cannot parse the current Date")
	}

	return &currTime, nil
}

func (s *metadataService) GetPrevious(currentSlug string) (*models.NavItem, error) {
	currentTime, err := getCurrentDate(currentSlug)

	layout := "02/01/2006"
	prevDate := currentTime.Add(-(time.Hour * 24)).Format(layout)

	metadatas, err := getAll()
	if err != nil {
		return nil, err
	}

	prevSlug, ok := metadatas[prevDate]
	if !ok {
		return nil, nil
	}

	data, err := os.ReadFile(fmt.Sprintf("./blogs/%s.md", prevSlug))
	if err != nil {
		return nil, errors.New("Cannot get the blog you request")
	}

	metadata, err := s.markdownService.GetMetadata(data)
	if err != nil {
		return nil, err
	}

	generatedSlug := slug.Make(metadata.Title)

	return &models.NavItem{
		Name: metadata.Title,
		Slug: generatedSlug,
	}, nil

}

func (s *metadataService) GetNext(currentSlug string) (*models.NavItem, error) {
	currentTime, err := getCurrentDate(currentSlug)

	layout := "02/01/2006"
	nextDate := currentTime.Add(time.Hour * 24).Format(layout)

	metadatas, err := getAll()
	if err != nil {
		return nil, err
	}

	nextSlug, ok := metadatas[nextDate]
	if !ok {
		return nil, nil
	}

	data, err := os.ReadFile(fmt.Sprintf("./blogs/%s.md", nextSlug))
	if err != nil {
		return nil, errors.New("Cannot get the blog you request")
	}

	metadata, err := s.markdownService.GetMetadata(data)
	if err != nil {
		return nil, err
	}

	generatedSlug := slug.Make(metadata.Title)

	return &models.NavItem{
		Name: metadata.Title,
		Slug: generatedSlug,
	}, nil

}
