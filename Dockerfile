FROM golang:latest

WORKDIR /go/src/shorterql
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["shorterql"]

