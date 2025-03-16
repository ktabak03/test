# Используем официальный образ Go как базовый
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum (если есть) и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Используем минимальный образ для финального контейнера
FROM alpine:latest

# Копируем собранный бинарный файл из builder-стадии
COPY --from=builder /app/main /main

# Указываем команду, которая будет запущена при старте контейнера
ENTRYPOINT ["/main"]
