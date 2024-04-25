run:
	@templ generate
	@./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
	@templ generate
	@go run cmd/main.go