package main

import (
	"api.default.marincor.com/app/appinstance"
	"api.default.marincor.com/config/constants"
	"api.default.marincor.com/handler/clientes"
	"api.default.marincor.com/handler/health"
	"api.default.marincor.com/middleware"
	"api.default.marincor.com/pkg/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func route() *fiber.App {
	appinstance.Data.Server.Use(logger.New())
	appinstance.Data.Server.Use(recover.New())
	appinstance.Data.Server.Use(favicon.New())
	appinstance.Data.Server.Use(cors.New(cors.Config{
		AllowMethods: constants.AllowedMethods,
		AllowHeaders: constants.AllowedHeaders,
	}))
	appinstance.Data.Server.Use(middleware.ValidateContentType())
	appinstance.Data.Server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	apiGroup := appinstance.Data.Server.Group("")
	apiGroup.Get("/health", health.Handle().Check, app.Log)

	clienteHandler := clientes.Handle()
	clientesGroup := apiGroup.Group("clientes")
	clientesGroup.Post("/:id/transacoes", clienteHandler.Create, app.Log)
	clientesGroup.Get("/:id/extrato", clienteHandler.Balance, app.Log)

	return appinstance.Data.Server
}
