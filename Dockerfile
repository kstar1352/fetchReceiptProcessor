# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . ./

RUN go build -v -o /receipt-processor

EXPOSE 3000


CMD ["/receipt-processor"]