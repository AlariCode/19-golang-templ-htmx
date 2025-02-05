# Этап 1
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/main.go

# Этап 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/public ./public
EXPOSE 3000

CMD ["./main"]