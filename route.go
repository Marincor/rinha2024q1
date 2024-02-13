package main

import (
	"api.default.marincor.com/app/appinstance"
	"api.default.marincor.com/handler/clientes"
	"api.default.marincor.com/handler/health"
	"api.default.marincor.com/pkg/app"
	"github.com/gofiber/fiber/v2"
)

func route() *fiber.App {
	apiGroup := appinstance.Data.Server.Group("")
	apiGroup.Get("/health", health.Handle().Check, app.Log)

	clienteHandler := clientes.Handle()
	clientesGroup := apiGroup.Group("clientes")
	clientesGroup.Post("/:id/transacoes", clienteHandler.Create, app.Log)
	clientesGroup.Get("/:id/extrato", clienteHandler.Balance, app.Log)

	return appinstance.Data.Server
}
