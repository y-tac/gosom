FROM golang:1.10.3-alpine3.8 AS build-env
ENV GOPATH /go
WORKDIR /go/src/github.com/y-tac/gosom
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server -v

FROM busybox
COPY front/dist  /front/dist
COPY config.json /config.json
COPY --from=build-env /go/src/github.com/y-tac/gosom/server /server

ENTRYPOINT ["/server"]
