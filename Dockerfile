FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/rubenv/sql-migrate/...@latest

CMD ["air"]
