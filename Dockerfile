FROM golang:1.8

WORKDIR /go/src/app
COPY . .

WORKDIR /go/src/app/api

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["api"]