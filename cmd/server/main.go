package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"
	// "strconv"

	"github.com/aventhis/avito-pvz-service/internal/config"
	"github.com/aventhis/avito-pvz-service/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Подключаемся к базе данных
	database, err := db.New(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer database.Close()

	// Применяем миграции
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Ошибка выполнения миграций: %v", err)
	}

	// Создаем Echo-сервер
	e := echo.New()

	// Добавляем middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Добавляем тестовый маршрут
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Запускаем сервер
	serverPort := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("HTTP-сервер запущен на порту %s", serverPort)
	if err := e.Start(serverPort); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска HTTP-сервера: %v", err)
	}
}
