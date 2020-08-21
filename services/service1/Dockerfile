FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN pwd
RUN ls -alrt
RUN go build -i services/service1/main.go .
CMD ["/app/main"]