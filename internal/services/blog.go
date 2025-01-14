package services

import (
	"blogstreak/models"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

type blogService struct {
	markdownService MarkdownService
}

func NewBlogService(markdownService MarkdownService) *blogService {
	return &blogService{
		markdownService: markdownService,
	}
}

func (s *blogService) GetAll() ([]string, error) {
	files, err := os.ReadDir("./blogs")
	if err != nil {
		return nil, err
	}

	blogs := []string{}

	for _, file := range files {
		blogs = append(blogs, file.Name())
	}

	return blogs, nil
}

func (s *blogService) Get(name string) (*models.Blog, error) {
	blogName := name + ".md"
	if ok := s.validate(blogName); !ok {
		return nil, errors.New("The blog that you request does not found.")
	}

	data, err := os.ReadFile("blogs/" + blogName)
	if err != nil {
		return nil, err
	}

	return s.markdownService.ParseMD(data)

}

func (s *blogService) validate(name string) bool {
	blogs, err := s.GetAll()
	if err != nil {
		return false
	}

	isFound := false
	for _, blog := range blogs {
		if name == blog {
			isFound = true
			break
		}
	}

	return isFound
}

func (s *blogService) New(name string) {
	path := fmt.Sprintf("./blogs/%s", name)
	_, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		_, err := os.Create(path)
		if err != nil {
			log.Fatal("Cannot create new blog")
		}
	}

}
