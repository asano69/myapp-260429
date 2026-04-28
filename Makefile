.PHONY: build
build:
	go build -o build ./cmd/myapp

.PHONY: watch
watch:
	air

.PHONY: serve
serve:
	myapp serve

