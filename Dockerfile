FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod ./
# Once we have required third party modules...
# COPY go.mod go.sum ./
# RUN go mod download

ADD config ./config
ADD main ./main

RUN go build -o go-playground ./main

CMD ["./go-playground"]
