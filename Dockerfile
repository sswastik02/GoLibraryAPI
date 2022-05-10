FROM golang:alpine3.15

RUN apk update && apk add --no-cache git

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o main .

EXPOSE 8000

CMD ["/app/main"]