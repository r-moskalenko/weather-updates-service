FROM golang:1.24.0

WORKDIR /go/src/app

COPY . .

RUN go build -o main .

CMD ["./main"]
