package main

import (
	"blogstreak/components"
	"blogstreak/models"
	"context"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := components.Page(&models.Blog{
			Title:       "Hello",
			Body:        "Hello World",
			PublishedAt: "01/01/2025",
		}, &models.Navigation{
			Previous: models.NavItem{
				Name: "Day Zero",
				Slug: "/day-zero",
			},
			Next: models.NavItem{
				Name: "Day Two",
				Slug: "/day-two",
			},
		})
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/blog/:slug", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "blogs/day-one.md")
	})

	http.HandleFunc("/assets/background.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/background.gif")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
