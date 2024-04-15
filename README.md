# go-htmx-app-template
### Template of a web application using:
```
Golang
Echo framework
Templ
Gorm
HTMX
Tailwind CSS
Postgresql
```
### Why Echo?
Because Echo returns error in its handlers
https://www.youtube.com/watch?v=g-gsmike7qc

### Why HTMX?
Because in 99.9% of the cases, you don't need a bloated javascript framework, lol

---
### To install Go dependencies
```
go mod download
```
### Or
```
go get github.com/labstack/echo/v4
go get github.com/gorilla/sessions
go get github.com/labstack/echo-contrib/session
go get github.com/a-h/templ
go get -u gorm.io/gorm
go get github.com/wader/gormstore/v2
go get gorm.io/driver/postgres

go install github.com/a-h/templ/cmd/templ@latest
```
---
### To install Tailwind (on linux x64)
```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
./tailwindcss init

For different OS and architectures check: https://github.com/tailwindlabs/tailwindcss/releases
```
---
### To update your templates
```
templ generate
```
---
### To update css files
```
# Start a watcher
./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

# Compile and minify your CSS for production
./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
```
---
### To build and run the dockerfile
```
docker build -t go-htmx-app-template .
docker run -p 8080:8080 go-htmx-app-template:latest
```
### To build and run the docker compose
```
docker compose up
docker compose down
.
# Clean docker cache
docker system prune
``` 