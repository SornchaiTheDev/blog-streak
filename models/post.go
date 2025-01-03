package models

import "github.com/a-h/templ"

type Blog struct {
	Title         string
	Body          templ.Component
	PublishedDate string
}
