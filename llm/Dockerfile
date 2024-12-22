# Базовый образ Python
FROM python:3.9-slim

# Установка системных зависимостей
RUN --mount=type=cache,target=/var/cache/apt apt update && apt install -y \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

# Установка рабочего каталога
WORKDIR /app

# Копирование файлов приложения
COPY . .

# Установка Python-зависимостей
RUN pip install --no-cache-dir -r requirements.txt

# Открываем порт
EXPOSE 7000

# Команда запуска сервера
CMD ["python", "server.py"]
