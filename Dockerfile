FROM golang:1.22 as proto-builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags musl -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o /main cmd/main.go

EXPOSE 8080
CMD ["/main"]