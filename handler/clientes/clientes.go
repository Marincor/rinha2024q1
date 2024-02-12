package clientes

import (
	"api.default.marincor.com/app/usecases/clientes"
	"api.default.marincor.com/config/constants"
	"api.default.marincor.com/entity"
	"api.default.marincor.com/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase   *clientes.Usecase
	validator *Validator
}

func Handle() *Handler {
	return &Handler{
		usecase:   clientes.New(),
		validator: NewValidator(),
	}
}

func (handler *Handler) Create(ctx *fiber.Ctx) error {
	var (
		request entity.TransacaoRequest
		params  entity.PathParams
	)

	if err := ctx.ParamsParser(&params); err != nil {
		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
			Message:    "error to get cliente id in transaction",
			Reason:     err.Error(),
			StatusCode: constants.HTTPStatusInternalServerError,
		})
		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

		helpers.CreateResponse(ctx, &entity.ErrorResponse{
			Message:     "error to get cliente id in transaction",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		}, constants.HTTPStatusBadRequest)

		return ctx.Next()
	}

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
			Message:    "error to get body in transaction",
			Reason:     err.Error(),
			StatusCode: constants.HTTPStatusInternalServerError,
		})
		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

		helpers.CreateResponse(ctx, &entity.ErrorResponse{
			Message:     "error to get body in transaction",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		}, constants.HTTPStatusBadRequest)

		return ctx.Next()
	}

	if statusCode, err := handler.validator.ValidateCreateTransaction(params.ID, &request); err != nil {
		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
			Message:    "error to create transaction",
			Reason:     err.Error(),
			StatusCode: statusCode,
		})
		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

		helpers.CreateResponse(ctx, &entity.ErrorResponse{
			Message:     "error to create transaction",
			Description: err.Error(),
			StatusCode:  statusCode,
		}, statusCode)

		return ctx.Next()
	}

	response, statusCode, err := handler.usecase.CreateTransaction(params.ID, request)
	if err != nil {
		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
			Message:    "error to create transaction",
			Reason:     err.Error(),
			StatusCode: statusCode,
		})
		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

		helpers.CreateResponse(ctx, &entity.ErrorResponse{
			Message:     "error to create transaction",
			Description: err.Error(),
			StatusCode:  statusCode,
		}, statusCode)

		return ctx.Next()
	}

	ctx.Locals(constants.LogDataKey, &entity.LogDetails{
		Message:    "transaction successfully created",
		StatusCode: constants.HTTPStatusOK,
		Response:   response,
	})
	ctx.Locals(constants.LogSeverityKey, constants.SeverityInfo)

	helpers.CreateResponse(ctx, response, constants.HTTPStatusOK)

	return ctx.Next()
}
