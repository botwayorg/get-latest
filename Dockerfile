FROM alpine:latest

WORKDIR /app

COPY ./dist/get-latest .

EXPOSE 8080

ENTRYPOINT ["./get-latest"]

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
