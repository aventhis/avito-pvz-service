# avito-pvz-service

Сервис для управления пунктами выдачи заказов (ПВЗ), приёмками и товарами. Выполнено в рамках отбора на стажировку Go-разработчика весной 2025 в Авито.

---

## 🔧 Сборка и запуск

Для автоматизации задач используется [`task`](https://taskfile.dev).

### Установка `task`:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

### Установка зависимостей:

```bash
go mod tidy
```

### Запуск сервера:

```bash
go run ./cmd/server
```

### Запуск через Docker:

1. Создайте файл `.env` на основе шаблона `.env.example`:

```bash
cp .env.example .env
```

2. Запустите сервисы через docker-compose:

```bash
task docker-up
```

или напрямую:

```bash
docker-compose up -d
```

---

## ⚙️ Генерация кода из OpenAPI

Проект использует [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) для генерации:

- DTO (структур)
- интерфейса `ServerInterface`
- типов параметров запросов

### Установка генератора:

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

Генерация выполняется через `task`, подробнее — в разделе [🛠 Taskfile](#taskfile).

---

> Конфигурация находится в `api/config.yaml`, результат генерации — `internal/api/openapi.gen.go`.

---

## Taskfile

Для автоматизации рутинных задач используется [`task`](https://taskfile.dev).

### Доступные команды:

#### ✅ Генерация DTO и интерфейсов из OpenAPI схемы:

```bash
task generate-api
```

#### 🐳 Docker команды:

```bash
task docker-up      # Запуск сервисов
task docker-down    # Остановка сервисов
task docker-logs    # Просмотр логов приложения
task docker-rebuild # Пересборка и перезапуск сервисов
```

#### ℹ️ Посмотреть все задачи:

```bash
task --list
```

> Taskfile позволяет единообразно запускать команды на всех этапах: от генерации до CI.

---

## 🌐 Переменные окружения

Для запуска приложения требуются следующие переменные окружения:

| Переменная    | Описание                       | Пример значения      |
|---------------|--------------------------------|----------------------|
| DB_HOST       | Хост БД                        | postgres             |
| DB_PORT       | Порт БД                        | 5432                 |
| DB_USER       | Пользователь БД                | postgres             |
| DB_PASSWORD   | Пароль пользователя БД         | postgres             |
| DB_NAME       | Имя базы данных                | avito_pvz            |
| DB_SSL_MODE   | Режим SSL для соединения с БД  | disable              |
| APP_ENV       | Окружение приложения           | development          |
| APP_PORT      | Порт HTTP API                  | 8080                 |
| GRPC_PORT     | Порт gRPC API                  | 3000                 |
| METRICS_PORT  | Порт метрик Prometheus         | 9000                 |

Скопируйте файл `.env.example` в `.env` и настройте значения под своё окружение.

---

## 📁 Структура проекта

```
avito-pvz-service/
├── api/                  # OpenAPI схема и конфиг генерации
│   ├── swagger.yaml
│   └── config.yaml
├── cmd/
│   └── server/           # Точка входа в HTTP-сервер
├── deployments/          # Файлы для деплоя (CI/CD, Kubernetes и т.д.)
├── internal/
│   ├── api/              # DTO и интерфейсы, сгенерированные oapi-codegen
│   ├── config/           # Конфигурация приложения
│   ├── db/               # Работа с базой данных
│   ├── grpc/             # gRPC сервис
│   ├── middlewares/      # Middleware для HTTP-запросов
│   ├── models/           # Модели данных
│   ├── repositories/     # Репозитории для работы с данными
│   ├── services/         # Бизнес-логика
│   └── tests/            # Тесты
├── migrations/           # Миграции базы данных
├── Dockerfile            # Сборка Docker-образа
├── docker-compose.yml    # Конфигурация Docker Compose
├── .env.example          # Шаблон переменных окружения
├── Taskfile.yml          # Автоматизация задач
├── README.md             # Техническая документация
├── go.mod / go.sum       # Зависимости Go
└── .gitignore
```
