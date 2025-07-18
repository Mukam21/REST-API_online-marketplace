REST API для онлайн-маркетплейса


Простое API для размещения объявлений с авторизацией пользователей. Реализовано на Go с использованием PostgreSQL, Gin и JWT.

# Технологии
Язык: Go 1.24

Фреймворк: Gin

База данных: PostgreSQL

Аутентификация: JWT

Контейнеризация: Docker

Быстрый старт

Требования
Docker и Docker Compose

Go 1.24+ (если запускаете без Docker)

        git clone https://github.com/Mukam21/REST-API_online-marketplace.git
        cd REST-API_online-marketplace
        docker-compose up --build

API будет доступно на http://localhost:8080.

Тестирование API

Через Postman

Импортируйте коллекцию Postman или создайте запросы вручную:

# Регистрация пользователя

        POST http://localhost:8080/register
        Content-Type: application/json

        {
            "login": "user@example.com",
            "password": "securepassword123"
        }

# Авторизация

        POST http://localhost:8080/login
        Content-Type: application/json

        {
            "login": "user@example.com",
            "password": "securepassword123"
        }
    Сохраните токен из ответа

# Создание объявления

        POST http://localhost:8080/orders
        Authorization: Bearer <ваш_токен>
        Content-Type: application/json

        {
            "title": "Ноутбук",
            "description": "Новый",
            "price": 50000,
            "image_url": "https://example.com/laptop.jpg"
        }

# Получение ленты

        GET http://localhost:8080/orders?page=1&min_price=10000&sort_by=price

Через терминал (curl)

# Регистрация
        curl -X POST http://localhost:8080/register \
        -H "Content-Type: application/json" \
        -d '{"login":"test@example.com","password":"qwerty123"}'

# Авторизация
        curl -X POST http://localhost:8080/login \
        -H "Content-Type: application/json" \
        -d '{"login":"test@example.com","password":"qwerty123"}'

# Создание объявления (подставьте ваш токен)
        curl -X POST http://localhost:8080/orders \
        -H "Authorization: Bearer ваш_токен" \
        -H "Content-Type: application/json" \
        -d '{"title":"Телефон","price":25000,"image_url":"http://example.com/phone.jpg"}'

# Получение ленты
        curl "http://localhost:8080/orders?page=1"

