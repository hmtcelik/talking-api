FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app/src

RUN go mod download talkyapi

RUN go build -o main .

CMD ["/app/src/main"]
