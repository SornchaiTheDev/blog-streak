package main

import (
	"blogstreak/components"
	"blogstreak/internal/services"
	"blogstreak/models"
	"blogstreak/shared"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	markdownService := services.NewMarkdownService()
	blogService := services.NewBlogService(markdownService)
	metadataService := services.NewMetadataService(markdownService)
	streakService := services.NewStreakService()

	http.HandleFunc("/blogs/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")

		clientIP := r.Header.Get("X-FORWARDED-FOR")
		log.Printf("🌏 Request from IP %s", clientIP)

		blog, err := blogService.Get(slug)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong")
			return
		}

		prev, err := metadataService.GetPrevious(slug)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong")
			return
		}
		next, err := metadataService.GetNext(slug)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Something went wrong")
			return
		}

		component := components.BlogPage(blog, &models.Navigation{
			Previous: prev,
			Next:     next,
		})

		component.Render(context.Background(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-FORWARDED-FOR")
		log.Printf("🌏 Request from IP %s", clientIP)

		metadatas, err := metadataService.GetAll()
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "Something went wrong")
			return
		}

		count := strconv.Itoa(len(metadatas))

		component := components.HomePage(metadatas, count)
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/api/streaks", func(w http.ResponseWriter, r *http.Request) {
		amount, longest := streakService.Get()
		tag := strconv.Itoa(amount)
		eTag := r.Header.Get("If-None-Match")

		if eTag == tag {
			w.WriteHeader(http.StatusNotModified)
		}

		w.Header().Set("Cache-Control", "public, max-age=3600")

		oneHourExp := time.Now().UTC().Add(time.Hour * 1)
		w.Header().Set("Expires", oneHourExp.Format(http.TimeFormat))

		w.Header().Set("Etag", tag)

		component := components.Streaks(amount, longest)
		component.Render(context.Background(), w)
	})

	http.HandleFunc("/assets/background.gif", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/background.gif")
	})

	http.HandleFunc(shared.CssName, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "assets/styles.css")
	})

	fmt.Println("Server start litening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
