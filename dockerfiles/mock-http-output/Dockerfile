FROM gcr.io/gcp-runtimes/go1-builder:1.9 as builder

WORKDIR /go/src/app
COPY main.go .

RUN /usr/local/go/bin/go build -o app .

# Application image.
FROM gcr.io/distroless/base:latest

COPY --from=builder /go/src/app/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]