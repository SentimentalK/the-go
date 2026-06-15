FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go test ./...
RUN go build -o proj3 .

CMD ["./proj3"]
