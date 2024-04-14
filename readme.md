### Web server
```
go get github.com/labstack/echo/v4
go get github.com/gorilla/sessions

go get github.com/a-h/templ
go install github.com/a-h/templ/cmd/templ@latest

go get -u gorm.io/gorm
go get gorm.io/driver/postgres

https://github.com/wader/gormstore


https://github.com/antonlindstrom/pgstore

https://medium.com/@fadhlimulyana20/building-rest-api-in-go-with-echo-gorm-and-postgresql-6734cae2b9cf

https://github.com/polaris1119/echo-login-example/blob/master/main.go

https://github.com/TomDoesTech/GOTTH/blob/main/cmd/main.go



https://www.alexedwards.net/blog/organising-database-access
https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4
https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
https://github.com/labstack/echo/issues/2075#issuecomment-1016819041

curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss

./tailwindcss -i ./web_server/static/css/input.css -o ./web_server/static/css/style.css --watch

./tailwindcss -i ./web_server/static/css/input.css -o ./web_server/static/css/style.min.css --minify

```

## PROD
```
templ generate
./tailwindcss -i ./web_server/static/css/input.css -o ./web_server/static/css/style.min.css --minify
go run web_server/cmd/main.go
```