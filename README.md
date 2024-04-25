# go-htmx-app-template
Because you don't need a bloated javascript framework :joy:

### Template of a web application using:
```text
Golang
Echo framework
Templ
Gorm
HTMX
Tailwind CSS
Postgresql
```

---
### To install Go dependencies
```bash
go mod download
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
# Or
go get github.com/labstack/echo/v4
go get github.com/gorilla/sessions
go get github.com/labstack/echo-contrib/session
go get github.com/a-h/templ
go get -u gorm.io/gorm
go get github.com/wader/gormstore/v2
go get gorm.io/driver/postgres
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```
---
### To install Tailwind (on linux x64)
```bash
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
./tailwindcss init

# For different OS/Architecture check: 
# https://github.com/tailwindlabs/tailwindcss/releases
```
---
### To update your templates
```bash
templ generate
```
---
### To update css files
```bash
# Start a watcher
./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

# Compile and minify your CSS for production
./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
```

### To update and run
```bash
make
# or
templ generate
./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
go run cmd/main.go

# You can also configure air for hot reloading if you prefer:
# https://github.com/cosmtrek/air
```

```
---
### To build and run the dockerfile
```bash
docker build -t go-htmx-app-template .
docker run -p 8080:8080 go-htmx-app-template:latest
```
### To build and run the docker compose
```bash
docker compose up
# or
docker-compose -f docker-compose-complete.yml up
# then
docker compose down
.
# Clean docker cache
docker system prune
``` 