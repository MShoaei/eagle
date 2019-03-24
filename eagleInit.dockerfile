FROM golang:alpine
RUN apk add git
RUN mkdir -p /go/src/github.com/MShoaei/eagle
WORKDIR /go/src/github.com/MShoaei/eagle
COPY . .
RUN go get -d -v ./...
# RUN go get -d github.com/99designs/gqlgen
# RUN go get -d github.com/vektah/gqlparser
# RUN go get -d github.com/gorilla/websocket
# RUN go get -d github.com/hashicorp/golang-lru
# RUN go get -d github.com/agnivade/levenshtein
# RUN go get -d github.com/dgrijalva/jwt-go
# RUN go get -d github.com/jinzhu/gorm
# RUN go get -d github.com/jinzhu/inflection
# RUN go get -d github.com/lib/pq
# RUN go get -d github.com/gofrs/uuid
# RUN go get -d golang.org/x/crypto/bcrypt
# RUN go get -d github.com/kataras/muxie

# RUN go install .