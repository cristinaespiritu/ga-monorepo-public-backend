FROM golang:latest AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN pwd
RUN go build -i -v
CMD ["/app/main"]