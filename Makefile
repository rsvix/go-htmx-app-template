run:
	@temple generate
	@./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
	@go run cmd/server.go