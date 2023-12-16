FROM golang:latest AS builder

# Копирование исходного кода в контейнер
COPY . /app

# Установка рабочей директории
WORKDIR /app/cmd

# Компиляция приложения
RUN go build -o /cmd/main .

# Контейнер для запуска
FROM golang:latest

# Копирование скомпилированного приложения из предыдущего контейнера
COPY --from=builder /cmd/main /cmd/main

# Базовый образ
FROM golang:latest
# Копирование исходного кода в контейнер
COPY . /app

# Установка рабочей директории
WORKDIR /app/cmd

# Компиляция приложения
RUN go build -o /cmd/main .

# Команда запуска приложения
CMD ["./app/main"]