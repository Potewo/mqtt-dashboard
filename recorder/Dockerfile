FROM golang:1.18-alpine


RUN apk upgrade --update && \
    apk --no-cache add git

# COPY . /go/src

WORKDIR /go/src/recorder

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
