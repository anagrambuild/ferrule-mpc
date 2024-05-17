FROM golang:1.22.2-alpine3.18 as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR /
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/ferrule ./bin
FROM scratch
WORKDIR /
COPY --from=builder /go/bin/ferrule /ferrule
EXPOSE 7946
CMD ["/ferrule"]