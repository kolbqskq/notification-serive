# Notification Service
 
Микросервисная система уведомлений на Go. Принимает запросы через REST API, доставляет уведомления в Telegram и хранит историю доставки.

### Сервисы
 
| Сервис | Описание | Порт |
|--------|----------|------|
| **api-gateway** | REST API, rate limiting, публикация в Kafka | 8081 |
| **notification-worker** | Читает из Kafka, отправляет в Telegram, сохраняет историю | — |
| **history-service** | gRPC сервер, отдаёт историю уведомлений | 9090 |
 
### Стек технологий
 
- **Go**
- **PostgreSQL**
- **Docker**
- **[gin](https://github.com/gin-gonic/gin)**
- **[kafka-go](github.com/segmentio/kafka-go)**
- **[go-redis](https://github.com/redis/go-redis)**
- **[pgx](https://github.com/jackc/pgx)**
- **[grpc-go](https://github.com/grpc/grpc-go)**
- **[go-telegram](github.com/go-telegram/bot)**
- **[zerolog](https://github.com/rs/zerolog)**
- **[golang-migrate](https://github.com/golang-migrate/migrate)**
- **[swaggo/swag](https://github.com/swaggo/swag)**
 
## Быстрый старт
 
### Требования
 
- Docker
- Telegram бот api-key (получить у [@BotFather](https://t.me/BotFather))
- Telegram Chat ID (получить у [@userinfobot](https://t.me/userinfobot))
 
### Установка
 
**1. Клонируй репозиторий:**
```bash
git clone https://github.com/kolbqskq/notification-serive.git
cd notification-serive
```
 
**2. Создай `.env` файлы из примеров и заполни:**
```bash
cp .env.example .env
cp api-gateway/.env.example api-gateway/.env
cp notification-worker/.env.example notification-worker/.env
cp history-service/.env.example history-service/.env
```
 
**3. Запусти инфраструктуру:**
```bash
make env-up
```
 
**4. Примени миграции:**
```bash
make migrate-up
```
 
**5. Запусти сервисы:**
```bash
make services-up
```
 
Swagger UI:
```
http://localhost:8081/swagger/index.html
```
 
## API
 
### POST /api/v1/notifications
 
Отправить уведомление.
 
**Тело запроса:**
```json
{
  "user_id": "uuid",
  "message": "Текст уведомления"
}
```
 
**Ответ 202 Accepted:**
```json
{
  "id": "uuid",
  "status": "Статус уведомления"
}
```
 
**Коды ошибок:**
- `400` — невалидный запрос
- `429` — превышен rate limit (5 запросов в минуту с одного IP)
- `500` — ошибка сервера
 
### GET /api/v1/history
 
История уведомлений пользователя.
 
**Query параметры:**
- `user_id` — UUID пользователя (обязательный)
- `limit` — количество записей (по умолчанию 10, максимум 100)
- `offset` — смещение (по умолчанию 0)
 
### GET /api/v1/notifications/:id
- `id` — UUID уведомления
 
Статус конкретного уведомления.
 
## Makefile команды
 
```bash
# Инфраструктура
make env-up              # Запустить postgres, redis, kafka
make env-down            # Остановить инфраструктуру
make env-cleanup         # Очистить данные (осторожно!)
make env-port-forward    # Пробросить порт PostgreSQL наружу
make env-port-close      # Закрыть проброс порта
 
# Сервисы
make services-up         # Запустить все сервисы
make services-down       # Остановить сервисы
make services-rebuild    # Пересобрать образы
 
# Миграции
make migrate-up                 # Применить миграции
make migrate-down               # Откатить миграции
make migrate-create seq=name    # Создать новую миграцию
 
# Документация
make swagger-gen         # Сгенерировать Swagger документацию
 
# Утилиты
make ps                  # Статус контейнеров
```
 
## Структура проекта
 
```
notification-service/
├── api-gateway/              
│   ├── cmd/                  
│   ├── docs/                 
│   └── internal/
│       ├── core/            
│       └── features/         
├── notification-worker/      
│   ├── cmd/
│   └── internal/
│       ├── core/
│       └── features/
├── history-service/          
│   ├── cmd/
│   └── internal/
│       ├── core/
│       └── features/
├── proto/                   
├── migrations/               
├── Dockerfile               
├── docker-compose.yml
└── Makefile
```
## Примечание

Это учебный проект, для простоты отправка уведомлений в телеграм захардкожена уведомления приходят только на тот чат айди который указан в .env