# Phrase 1: Compile TailwindCSS
FROM node:20 AS tailwind

WORKDIR /app

COPY package.json tailwind.config.js assets/input.css ./
COPY ./components/ ./components

RUN npm i

RUN npx tailwindcss -i input.css -o styles.css


# Phrase 2: Compile Go
FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .
COPY --from=tailwind /app/styles.css ./assets/styles.css

RUN go mod download

RUN go run ./cmd/ci/ci.go

RUN GOOS=linux GOARCH=amd64 go build -o ./main ./cmd/server/server.go


# Phrase 3: lean image
FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/main ./main
COPY --from=build /app/blogs ./blogs
COPY --from=build /app/blogs_metadata.json ./blogs_metadata.json
COPY --from=build /app/streak.json ./streak.json
COPY --from=build /app/assets ./assets

EXPOSE 3000

CMD ["./main"]

