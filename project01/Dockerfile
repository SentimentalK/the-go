FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .

RUN go build -o proj1 main.go

CMD ["./proj1"]