FROM golang

ADD . /go/src/github.com/MagnusTiberius/weatherservice

ENTRYPOINT /go/bin/outyet

EXPOSE 8080
