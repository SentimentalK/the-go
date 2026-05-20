FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY data ./data
COPY model ./model
COPY utils ./utils

RUN go build -o proj1 .

CMD ["./proj1"]
