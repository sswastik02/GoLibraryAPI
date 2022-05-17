FROM golang:alpine3.15

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# copying only go.mod and go.sum ensures that the download is cached even when the rest of the code is changed

COPY . .

RUN go build -o main .

EXPOSE 8000

ENTRYPOINT ["/app/main","-runserver"]

# when using command line arguments entrypoint should be used over CMD