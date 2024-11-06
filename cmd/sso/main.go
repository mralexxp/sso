package main

import (
	"fmt"

	"github.com/mralexxp/sso/internal/config"
)

func main() {
	// TODO: инициализировать объект конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: Инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: Запустить gRPC сервер-приложения
}
