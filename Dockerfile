FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Финальный образ
FROM alpine:latest

WORKDIR /app

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates tzdata

# Копируем скомпилированное приложение
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations

# Установим часовой пояс Москвы
ENV TZ=Europe/Moscow

# Открываем порты для HTTP и gRPC
EXPOSE 8080 3000 9000

# Запускаем приложение
CMD ["./server"] 