FROM golang:1.19.0

WORKDIR /home/hadia/myproject/RentalCar

COPY . .
RUN go mod tidy