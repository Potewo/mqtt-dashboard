FROM golang:1.18-alpine

WORKDIR /go/src

RUN apk upgrade --update && \
    apk --no-cache add git

COPY . .
WORKDIR /gor/src/server

RUN go install github.com/cosmtrek/air@latest

EXPOSE 1323

CMD ["air", "-c", ".air.toml"]
