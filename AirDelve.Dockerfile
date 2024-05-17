FROM golang:1.22.2-alpine3.18
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /node
EXPOSE 2345
ENTRYPOINT ["air"]