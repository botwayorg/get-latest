.PHONY: build

build:
		@go mod tidy && \
		CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./dist/get-latest
