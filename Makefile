.PHONY: build

build:
		@go mod tidy && \
		go build -a -ldflags '-w -extldflags "-static"' -o ./dist/get-latest
