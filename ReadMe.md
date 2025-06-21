Full-backend Rental Platform
Комплексная платформа аренды жилья (аналог Airbnb) на Go с микросервисной архитектурой, включающая управление недвижимостью, систему бронирования, real-time чат и интеграцию с платежными системами.

🚀 Технологический стек
Go 1.23.5 - основной язык разработки
Chi Router v5.0.12 - HTTP маршрутизация
PostgreSQL 17 - основная база данных
Kafka - асинхронная обработка платежей
JWT - аутентификация и авторизация
Docker & Docker Compose - контейнеризация
Google Wire - dependency injection
WebSocket - real-time коммуникация go.mod:3-30
📋 Основные функции
🔐 Аутентификация и безопасность
JWT-токены с HMAC-SHA256 подписью
SHA-256 хеширование паролей
OTP верификация для email и сброса паролей
Middleware для защиты эндпоинтов
🏠 Управление недвижимостью
CRUD операции для объявлений
Загрузка и управление изображениями
Система фильтрации по локации, цене, удобствам
Статистика для владельцев
📅 Система бронирования
Создание и управление бронированиями
Check-in/check-out функциональность
Проверка доступности дат
Отслеживание статусов бронирований
💳 Платежная система
Интеграция с Yookassa через Kafka
Асинхронная обработка платежей
Поддержка различных методов оплаты
💬 Система коммуникации
WebSocket чат в реальном времени
Система сообщений между пользователями
Connection manager для управления соединениями
🛠 Быстрый старт
Предварительные требования
Docker и Docker Compose
Make
Go 1.23+ (для локальной разработки)
Установка и запуск
Клонируйте репозиторий:
git clone https://github.com/imperatorofdwelling/Full-backend.git  
cd Full-backend
Настройте переменные окружения:
cp .env.example .env.local
# Отредактируйте .env.local с вашими настройками
Запустите приложение:
# Для локальной разработки
make docker-local

# Для среды разработки
make docker-dev

# Для staging
make docker-stage
cd.yml:8-42

📚 API Документация
API полностью документирован через Swagger и доступен по адресу:

http://localhost/api/v1/swagger/  
Основные эндпоинты:
Аутентификация: /api/v1/login, /api/v1/registration
Недвижимость: /api/v1/stays/*
Бронирование: /api/v1/reservation/*
Чат: /api/v1/chat/*
Локации: /api/v1/locations/*
Всего реализовано 80+ API эндпоинтов с полной документацией.

🏗 Архитектура
Приложение следует принципам чистой архитектуры с четким разделением слоев:

├── internal/  
│   ├── api/           # HTTP handlers и middleware  
│   ├── domain/        # Бизнес-логика и интерфейсы  
│   ├── repo/          # Слой доступа к данным  
│   ├── service/       # Бизнес-сервисы  
│   └── di/            # Dependency injection  
├── cmd/               # Точки входа приложения  
├── docs/              # Swagger документация  
└── docker/            # Docker конфигурации  
🧪 Тестирование
Проект включает комплексное тестирование:

# Запуск всех тестов
make test

# Генерация моков
make mock

# Генерация Wire DI
make wire

# Генерация Swagger документации
make swag
🚀 Развертывание
CI/CD Pipeline
Автоматическое развертывание настроено через GitHub Actions:

CI этап: Генерация кода, тесты
CD этап: Развертывание контейнеров
Проверка: Верификация через Swagger API
Окружения
local - локальная разработка
dev - среда разработки
stage - предпродакшн
prod - продакшн
📝 Основные команды Make
make run-local      # Запуск локально  
make build          # Сборка приложения  
make docker-local   # Docker для локальной разработки  
make docker-dev     # Docker для разработки  
make migrate-up     # Применить миграции БД  
make migrate-down   # Откатить миграции БД
🤝 Вклад в проект
Форкните репозиторий
Создайте feature ветку (git checkout -b feature/amazing-feature)
Зафиксируйте изменения (git commit -m 'Add amazing feature')
Отправьте в ветку (git push origin feature/amazing-feature)
Откройте Pull Request
📄 Лицензия
Этот проект лицензирован под Apache 2.0 License - см. файл LICENSE для деталей.

📞 Контакты
API Support: support@swagger.io
Документация: Swagger UI
Issues: GitHub Issues