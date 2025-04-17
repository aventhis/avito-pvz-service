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
#### ℹ️ Посмотреть все задачи:

```bash
task --list
```

> Taskfile позволяет единообразно запускать команды на всех этапах: от генерации до CI.

---


## 📁 Структура проекта

```
avito-pvz-service/
├── api/                  # OpenAPI схема и конфиг генерации
│   ├── swagger.yaml
│   └── config.yaml
├── internal/
│   └── api/              # DTO и интерфейсы, сгенерированные oapi-codegen
├── cmd/
│   └── server/           # Точка входа в HTTP-сервер
├── Taskfile.yml          # Автоматизация генерации кода
├── README.md             # Техническая документация
├── go.mod / go.sum       # Зависимости Go
└── .gitignore
```
