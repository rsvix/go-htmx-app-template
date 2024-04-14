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

### Go dependencies
```
go get github.com/labstack/echo/v4
go get github.com/gorilla/sessions
go get github.com/a-h/templ
go install github.com/a-h/templ/cmd/templ@latest
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
```
### Install Tailwind
```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
```