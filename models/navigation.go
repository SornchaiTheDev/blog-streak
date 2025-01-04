package models

type Navigation struct {
	Previous *NavItem
	Next     *NavItem
}

type NavItem struct {
	Name string
	Slug string
}
