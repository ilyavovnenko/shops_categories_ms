package http

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	config "github.com/ilyavovnenko/shops_categories_ms/configs"

	attrUsecase "github.com/ilyavovnenko/shops_categories_ms/internal/attribute/usecase"
	"github.com/ilyavovnenko/shops_categories_ms/internal/response"
	responseService "github.com/ilyavovnenko/shops_categories_ms/internal/response/service"
)

type AttributeHandler struct {
	defConf         config.Default
	CategoryUsecase *attrUsecase.AttributeUsecase
	contextTimeout  time.Duration
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// It will initialize the attributes/ resources endpoints
func New(app *fiber.App, usecase *attrUsecase.AttributeUsecase, defConf config.Default, timeout time.Duration) {
	handler := &AttributeHandler{
		defConf:         defConf,
		CategoryUsecase: usecase,
		contextTimeout:  timeout,
	}

	app.Get("/attributes", handler.GetAllAttributes)
	app.Get("/attributes/:id", handler.GetAttributeByID)
}

func (ah *AttributeHandler) GetAllAttributes(ctx *fiber.Ctx) error {
	perPage := getCurrentPerPageParam(ctx, ah.defConf.PerPage)

	numPage := ctx.Query("page")
	page, _ := strconv.Atoi(numPage)

	c, cancel := context.WithTimeout(ctx.Context(), ah.contextTimeout)
	defer cancel()

	listAr, err := ah.CategoryUsecase.GetAllAttributes(c, page, perPage)
	if err != nil {
		appError := response.AppError{
			Code:    "Error",
			Message: err.Error(),
		}
		return responseService.InternalErrorResponse(ctx, []response.AppError{appError})
	}

	return responseService.SuccessResponse(ctx, listAr)
}

func (ah *AttributeHandler) GetAttributeByID(ctx *fiber.Ctx) error {
	idTemp, _ := strconv.Atoi(ctx.Params("id"))

	c, cancel := context.WithTimeout(ctx.Context(), ah.contextTimeout)
	defer cancel()

	category, err := ah.CategoryUsecase.GetAttributeByID(c, int64(idTemp))
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
