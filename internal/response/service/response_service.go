package service

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ilyavovnenko/shops_categories_ms/internal/response"
)

/*
* Responses
 */
func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(getSuccessResponse(c, message, data))
}

func InternalErrorResponse(c *fiber.Ctx, errors []response.AppError) error {
	return c.Status(fiber.StatusInternalServerError).JSON(getErrorResponse(c, "Internal error happens", errors))
}

func BadRequestResponse(c *fiber.Ctx, errors []response.ValidationError) error {
	return c.Status(fiber.StatusBadRequest).JSON(getValidationErrorResponse(c, "Validation was finished with error", errors))
}

/*
* Preparation block
 */
func getSuccessResponse(c *fiber.Ctx, message string, data interface{}) response.Response {
	return response.Response{
		Data:    data,
		Message: message,
		Meta:    generateMeta(c),
	}
}

func getErrorResponse(c *fiber.Ctx, message string, errors []response.AppError) response.Response {
	return response.Response{
		Data:       nil,
		Errors:     errors,
		Message:    message,
		Meta:       generateMeta(c),
		Validation: []response.ValidationError{},
	}
}

func getValidationErrorResponse(c *fiber.Ctx, message string, errors []response.ValidationError) response.Response {
	return response.Response{
		Validation: errors,
		Message:    message,
		Meta:       generateMeta(c),
	}
}

/*
* META block
 */
func generateMeta(c *fiber.Ctx) response.Meta {
	return response.Meta{
		Page:            0, // todo: add functionality for getting page and another staff
		PerPage:         0,
		Self:            generateSelfLink(c),
		TotalPages:      0,
		TotalItemsCount: 0,
	}
}

func generateSelfLink(c *fiber.Ctx) string {
	return c.BaseURL() + c.OriginalURL()
}
