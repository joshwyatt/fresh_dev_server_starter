FROM golang:1.10.3-alpine3.8

RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN apk add git
RUN apk add dep
RUN go get github.com/pilu/fresh

ADD . .

RUN dep ensure

EXPOSE 8000
CMD ["fresh"]
