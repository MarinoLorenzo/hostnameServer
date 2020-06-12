FROM golang:1.13-alpine

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o main .

CMD ["./main"]