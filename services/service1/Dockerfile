FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN cd services/service1/
RUN go build -o main .
CMD ["/app/main"]