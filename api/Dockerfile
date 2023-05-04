FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o api cmd/server/main.go

EXPOSE 8080

CMD ["./api"]