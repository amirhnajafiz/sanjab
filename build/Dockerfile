FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

FROM scratch:latest

WORKDIR /app

COPY --from=builder /main .

CMD ./main
