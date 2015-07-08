FROM golang:1.4.2

ADD . /go/src/github.com/orieken/go-rest

RUN go get github.com/julienschmidt/httprouter
RUN go get gopkg.in/mgo.v2

RUN go install github.com/orieken/go-rest

EXPOSE 7000
CMD ["/go/bin/go-rest"]
