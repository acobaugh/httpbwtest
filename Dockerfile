FROM golang:1.12-alpine AS builder

ENV GO111MODULE=on

RUN mkdir -p /go/src/ \
 && apk update \
 && apk add git 

WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpbwtest


FROM scratch

COPY --from=0 /go/src/httpbwtest /httpbwtest

CMD ["/httpbwtest"]
