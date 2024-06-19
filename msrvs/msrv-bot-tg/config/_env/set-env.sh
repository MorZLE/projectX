#!/bin/sh

ENV_FILE="msrvs/msrv-bot-tg/config/_secret.env"

if [ -f "$ENV_FILE" ]; then
    export $(cat $ENV_FILE | xargs)
    echo "Переменные окружения из файла .env успешно применены."
else
    echo "Файл .env не найден."
fi