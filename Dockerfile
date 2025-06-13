FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /

COPY --from=builder /app/cmd/server/todo-server /todo-server

COPY config /app/config

EXPOSE 8000

ENTRYPOINT ["/todo-server"]
