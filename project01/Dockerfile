FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .
COPY data ./data

RUN go build -o proj1 main.go

CMD ["./proj1"]