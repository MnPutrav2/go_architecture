build:
	go build ./cmd/server

template:
	go run ./cmd/cli make:template name=$(name) type=$(type)