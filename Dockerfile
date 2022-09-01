FROM golang:alpine AS builder

RUN apk add gcc g++ ca-certificates --no-cache

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./get-latest

ENV PORT=8080

EXPOSE 8080

ENTRYPOINT ["/app/get-latest"]
