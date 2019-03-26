# FROM golang:alpine as go
# RUN apk add git
# RUN mkdir -p /go/src/github.com/MShoaei/eagle
# WORKDIR /go/src/github.com/MShoaei/eagle
# COPY . .
# RUN go get -d -v ./...
# RUN go install .

FROM eagle:init as go
COPY . .
RUN go install .

FROM nginx:alpine
COPY --from=go /go/src/github.com/MShoaei/eagle/keys /go/bin/keys
COPY --from=go /go/bin/eagle /go/bin/eagle
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./passwd /go/bin/
COPY entrypoint.sh /go/bin/
WORKDIR /go/bin
RUN chmod +x entrypoint.sh
CMD "./entrypoint.sh"
