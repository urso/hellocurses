FROM golang:latest

COPY . /go/src/github.com/urso/hellocurses
WORKDIR /go/src/github.com/urso/hellocurses

RUN apt-get update && apt-get install -y ncurses-dev
RUN go build

CMD ./hellocurses
