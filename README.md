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
---
### To install Go dependencies
```
go get github.com/labstack/echo/v4
go get github.com/gorilla/sessions
go get github.com/labstack/echo-contrib/session
go get github.com/a-h/templ
go install github.com/a-h/templ/cmd/templ@latest
go get -u gorm.io/gorm
go get github.com/wader/gormstore/v2
go get gorm.io/driver/postgres
```
---
### To install Tailwind (on linux x64)
```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
./tailwindcss init

For different OS and achitectures check: https://github.com/tailwindlabs/tailwindcss/releases
```
---
### To update your templates, run
```
templ generate
```
---
### To update tailwind, run
```
# Start a watcher
./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

# Compile and minify your CSS for production
./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
```
---