package appinstance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	DB     *pgxpool.Pool
	Server *fiber.App
}

var Data *Application
