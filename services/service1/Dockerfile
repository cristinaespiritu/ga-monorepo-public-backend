FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN pwd
RUN cd services/service1/
RUN pwd
RUN go build -o main .
CMD ["/app/main"]