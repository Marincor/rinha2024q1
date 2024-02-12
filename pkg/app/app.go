package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"api.default.marincor.com/adapters/database"
	"api.default.marincor.com/adapters/logging"
	"api.default.marincor.com/app/appinstance"
	constanterrors "api.default.marincor.com/app/errors"
	"api.default.marincor.com/config"
	"api.default.marincor.com/config/constants"
	"api.default.marincor.com/entity"
	"api.default.marincor.com/pkg/helpers"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ApplicationInit() {
	configs := config.New()
	ctx := context.Background()

	appinstance.Data = &appinstance.Application{
		Config: configs,
		Server: fiber.New(fiber.Config{
			ServerHeader: "Rinha",
			ErrorHandler: customErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			Prefork:      false,
		}),
	}

	appinstance.Data.DB = database.Connect(ctx)
}

func Setup() {
	err := appinstance.Data.Server.Listen(fmt.Sprintf(":%s", constants.Port))

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func customErrorHandler(ctx *fiber.Ctx, err error) error {

	if strings.Contains(err.Error(), constanterrors.ErrErrorTooLargeEOFBuffer.Error()) {
		helpers.CreateResponse(ctx, entity.ErrorResponse{
			StatusCode: constants.HTTPStatusBadRequest,
			Message:    constanterrors.ErrErrorTooLargeEOFBuffer.Error(),
		}, constants.HTTPStatusBadRequest)

		return nil
	}

	var code int = fiber.StatusInternalServerError
	var capturedError *fiber.Error
	message := "unknown error"

	if errors.As(err, &capturedError) {
		code = capturedError.Code
		if code == fiber.StatusNotFound {
			message = "route not found"
		}
	}

	var errorResponse *entity.ErrorResponse

	erro := json.Unmarshal([]byte(err.Error()), &errorResponse)
	if erro != nil {
		errorResponse = &entity.ErrorResponse{
			Message:     message,
			StatusCode:  code,
			Description: err.Error(),
		}
	}

	go logging.Log(
		&entity.LogDetails{
			Message:  message,
			Method:   ctx.Method(),
			Reason:   err.Error(),
			RemoteIP: ctx.IP(),
			Request: map[string]interface{}{
				"body":  string(ctx.BodyRaw()),
				"query": ctx.Queries(),
			},
			StatusCode: code,
			URLpath:    ctx.Path(),
		},
		"critical",
		nil,
	)

	helpers.CreateResponse(ctx, errorResponse, code) //nolint: wrapcheck

	return nil
}

func Log(ctx *fiber.Ctx) error {
	logSeverity := ctx.Locals(constants.LogSeverityKey)

	payload := new(entity.LogDetails)
	bytedata, _ := helpers.Marshal(ctx.Locals(constants.LogDataKey))
	helpers.Unmarshal(bytedata, &payload) //nolint: errcheck

	if logSeverity == nil {
		logSeverity = "debug"
	}

	body := map[string]interface{}{}
	helpers.Unmarshal(ctx.BodyRaw(), &body) //nolint: errcheck

	request := map[string]interface{}{
		"body":       body,
		"query":      ctx.Queries(),
		"url_params": ctx.AllParams(),
	}

	severity := fmt.Sprintf("%v", logSeverity)

	go logging.Log(&entity.LogDetails{
		Message:    payload.Message,
		StatusCode: payload.StatusCode,
		Reason:     payload.Reason,
		Response:   payload.Response,
		Request:    request,
		Method:     ctx.Method(),
		RemoteIP:   ctx.IP(),
		URLpath:    ctx.Path(),
	}, severity, nil)

	return nil
}
