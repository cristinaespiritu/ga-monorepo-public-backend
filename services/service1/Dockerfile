FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o services/service1/main .
CMD ["/app/main"]