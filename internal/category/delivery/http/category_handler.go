package http

import (
	"context"
	"strconv"
	"time"

	config "github.com/ilyavovnenko/shops_categories_ms/configs"
	catUsecase "github.com/ilyavovnenko/shops_categories_ms/internal/category/usecase"
	"github.com/ilyavovnenko/shops_categories_ms/internal/response"
	responseService "github.com/ilyavovnenko/shops_categories_ms/internal/response/service"

	fiber "github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	defConf         config.Default
	CategoryUsecase *catUsecase.CategoryUsecase
	contextTimeout  time.Duration
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// It will initialize the categories/ resources endpoints
func New(app *fiber.App, usecase *catUsecase.CategoryUsecase, defConf config.Default, timeout time.Duration) {
	handler := &CategoryHandler{
		defConf:         defConf,
		CategoryUsecase: usecase,
		contextTimeout:  timeout,
	}

	app.Get("/categories", handler.GetAll)
	app.Get("/categories/:id", handler.GetByID)
}

func (ch *CategoryHandler) GetAll(ctx *fiber.Ctx) error {
	perPage := getCurrentPerPageParam(ctx, ch.defConf.PerPage)

	numPage := ctx.Query("page")
	page, _ := strconv.Atoi(numPage)

	c, cancel := context.WithTimeout(ctx.Context(), ch.contextTimeout)
	defer cancel()

	listAr, err := ch.CategoryUsecase.GetAll(c, page, perPage)
	if err != nil {
		appError := response.AppError{
			Code:    "Error",
			Message: err.Error(),
		}
		return responseService.InternalErrorResponse(ctx, []response.AppError{appError})
	}

	return responseService.SuccessResponse(ctx, listAr)
}

func (ch *CategoryHandler) GetByID(ctx *fiber.Ctx) error {
	idTemp, _ := strconv.Atoi(ctx.Params("id"))

	c, cancel := context.WithTimeout(ctx.Context(), ch.contextTimeout)
	defer cancel()

	category, err := ch.CategoryUsecase.GetByID(c, int64(idTemp))
	if err != nil {
		validationError := response.ValidationError{
			FailedField: "ID",
			Tag:         "error",
			Value:       err.Error(),
		}
		return responseService.BadRequestResponse(ctx, []response.ValidationError{validationError})
	}

	return responseService.SuccessResponse(ctx, category)
}

func getCurrentPerPageParam(ctx *fiber.Ctx, defPerPage uint16) uint16 {
	var perPage uint16

	num := ctx.Query("per_page")
	perPageNumber, _ := strconv.Atoi(num)

	if perPageNumber == 0 {
		perPage = defPerPage
	} else {
		perPage = uint16(perPageNumber)
	}

	return perPage
}
