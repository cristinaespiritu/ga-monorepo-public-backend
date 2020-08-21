FROM golang:latest AS builder
RUN mkdir /app
WORKDIR /app
COPY ./services/service1 .
RUN go build -i -v
CMD ["/app/main"]