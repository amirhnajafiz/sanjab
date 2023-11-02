FROM go:1.20-alpine

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

CMD ./main
