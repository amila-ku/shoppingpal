FROM golang:1.8

WORKDIR /go/src/app
COPY ./api/api .

RUN ls -lhtr

CMD ["/go/src/app/api"]