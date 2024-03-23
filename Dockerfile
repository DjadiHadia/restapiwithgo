FROM golang:1.19.0

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd 

CMD ["./main"]
