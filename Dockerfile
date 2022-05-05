FROM golang:alpine AS builder

RUN mkdir app

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o ./get-latest

FROM alpine:latest

RUN apk add gcc g++ ca-certificates --no-cache

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/get-latest"]

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/get-latest /app/get-latest
