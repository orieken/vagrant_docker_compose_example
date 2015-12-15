FROM golang:1.5

ADD . /go/src/github.com/orieken/go-rest

RUN go get github.com/julienschmidt/httprouter
#RUN go get gopkg.in/mgo.v2

RUN go install github.com/orieken/go-rest

EXPOSE 49200
CMD ["/go/bin/go-rest"]
