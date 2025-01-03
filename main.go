package main

import (
	"blogstreak/components"
	"blogstreak/internal/services"
	"blogstreak/models"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	blogService := services.NewBlogService()

	http.HandleFunc("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")

		blog, err := blogService.Get(slug)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong")
			return
		}

		component := components.Page(blog, &models.Navigation{
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

	http.HandleFunc("/assets/background.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/background.gif")
	})

	fmt.Println("Server start litening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
