FROM golang:1.16.7-alpine

ENV GO111MODULE=on

RUN apk add build-base

WORKDIR /bin/app

ADD . /bin/app

RUN go mod tidy -v
RUN go mod download

RUN go build -o main .

CMD ["/bin/app/main"]