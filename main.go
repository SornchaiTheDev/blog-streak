package main

import (
	"blogstreak/components"
	"blogstreak/internal/services"
	"blogstreak/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	blogService := services.NewBlogService()
	streakService := services.NewStreakService()

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

	http.HandleFunc("/api/streaks", func(w http.ResponseWriter, r *http.Request) {
		amount := streakService.Get()
		tag := strconv.Itoa(amount)
		eTag := r.Header.Get("If-None-Match")

		if eTag == tag {
			w.WriteHeader(http.StatusNotModified)
		}

		w.Header().Set("Cache-Control", "public, max-age=3600")

		oneHourExp := time.Now().UTC().Add(time.Hour * 1)
		w.Header().Set("Expires", oneHourExp.Format(http.TimeFormat))

		w.Header().Set("Etag", tag)

		fmt.Fprintf(w, "%d", amount)
	})

	http.HandleFunc("/assets/background.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/background.gif")
	})

	fmt.Println("Server start litening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
