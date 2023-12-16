# Контейнер для компиляции
FROM golang:latest AS builder

# Копирование исходного кода в контейнер
COPY . /app

# Установка рабочей директории
WORKDIR /app/cmd

# Компиляция приложения
RUN go build -o /app/cmd/main .

# Контейнер для запуска
FROM golang:latest

# Копирование скомпилированного приложения из предыдущего контейнера
COPY --from=builder /app/cmd/main /app/cmd/main

# Установка рабочей директории
WORKDIR /app/cmd

# Команда запуска приложения
CMD ["./main"]