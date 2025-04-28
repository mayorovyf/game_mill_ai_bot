#!/bin/bash

# сборка и запуск бота

PROJECT_DIR="."

cd "$PROJECT_DIR" || { echo "❌ Не удалось перейти в директорию проекта"; exit 1; }

echo "🔨 Сборка..."
go build cmd/main.go
if [ $? -ne 0 ]; then
  echo "❌  Сборка завершилась с ошибкой"
  exit 1
fi

echo "✅  Сборка успешна"

echo "🚀 Запуск..."
./main