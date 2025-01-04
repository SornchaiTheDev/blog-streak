up: 
	air

tailwind:
	pnpm dlx tailwindcss -i ./assets/input.css -o ./assets/styles.css --watch

templ:
	templ generate -watch -proxy http://localhost:3000

clean:
	rm -rf ./tmp
	rm -rf ./assets/styles.css
