.PHONY: build
build:
	go build -v ./cmd/blog
.PHONY: run
run: 
	./blog

.DEFAULT_GOAL := build	
