FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Установка goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

ARG USER=appuser

RUN adduser -D -g '' $USER

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ./migrations ./migrations

RUN chown -R $USER /app

USER $USER

EXPOSE 8080

# Выполнение миграций перед запуском приложения
CMD ["sh", "-c", "goose -dir ./migrations postgres $POSTGRES_ADDR up && ./main"]