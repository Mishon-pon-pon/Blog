.PHONY: build
build:
	go build -v ./cmd/blog
.PHONY: run
run: 
	./blog
.PHONY: test
test:
	go test -v -race -timeout 30s ./app/...

.DEFAULT_GOAL := build	
