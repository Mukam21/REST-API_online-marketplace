# REST API для онлайн-маркетплейса


Простое API для размещения объявлений с авторизацией пользователей. Реализовано на Go с использованием PostgreSQL, Gin и JWT.

Этот код реализует REST API для онлайн-маркетплейса со следующими ключевыми возможностями:

1. Система аутентификации
Регистрация новых пользователей:

Принимает логин и пароль

Проверяет минимальные требования (логин от 3 символов, пароль от 6 символов)

Хеширует пароль перед сохранением в БД

Защищает от дублирования логинов

Авторизация пользователей:

Проверяет соответствие логина и пароля

Генерирует JWT-токен для авторизованных запросов

Использует безопасное сравнение хешей паролей

2. Управление объявлениями
Создание объявлений:

Требует авторизации (валидный JWT-токен)

Валидирует входные данные (заголовок, описание, цена, URL изображения)

Автоматически связывает объявление с автором через JWT

Просмотр ленты объявлений:

Поддерживает пагинацию (page, limit)

Фильтрацию по цене (min_price, max_price)

Сортировку (sort_by=price|created_at, order=asc|desc)

Возвращает информацию об авторе каждого объявления

3. Безопасность
Все пароли хранятся в виде хешей (bcrypt)

Защищенные маршруты требуют JWT-токена

Автоматическая проверка валидности токена

Защита от SQL-инъекций через ORM (GORM)

4. Работа с базой данных
Автоматические миграции моделей при запуске

Связи между таблицами (User has many Ads)

Оптимизированные запросы с предзагрузкой данных

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

