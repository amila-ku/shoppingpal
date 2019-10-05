FROM golang:1.8

WORKDIR /go/src/github.com/amila-ku/shoppingpal
COPY . .

RUN go get -d -v ./...

WORKDIR /go/src/github.com/amila-ku/shoppingpal/api

RUN ls -lhtr; go build .

CMD ["api"]