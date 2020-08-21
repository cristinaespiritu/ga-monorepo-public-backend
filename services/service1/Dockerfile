FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN pwd
RUN ls -alrt
RUN go build -o main .
CMD ["/app/main"]