FROM golang:1.19-alpine

WORKDIR /app/auth

RUN apk update && apk add git
RUN apk add musl-dev gcc libc-dev

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]

EXPOSE 8080