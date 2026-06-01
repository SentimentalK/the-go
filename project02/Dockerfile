FROM golang:1.23-alpine

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY data ./data
COPY model ./model
COPY utils ./utils
COPY business ./business
COPY persistence ./persistence
COPY presentation ./presentation

RUN go build -o proj2 .

CMD ["./proj2"]
