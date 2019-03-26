FROM golang:alpine
RUN apk add git
RUN mkdir -p /go/src/github.com/MShoaei/eagle
WORKDIR /go/src/github.com/MShoaei/eagle
COPY . .
RUN go get -d -v ./...
