FROM golang:alpine as builder
WORKDIR /go/src/build
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0
COPY . .
RUN apk add git && \
    go get -u -v github.com/golang/dep/... && \
    dep ensure -v && \
    go build -v -ldflags="-s -w" -o echo-app main.go

FROM alpine:latest
WORKDIR /app
EXPOSE 8080

COPY --from=builder /go/src/build/echo-app ./echo-app
COPY templates ./templates

ENTRYPOINT ["/app/echo-app"]
