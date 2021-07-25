package service

import (
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/ilyavovnenko/shops_categories_ms/internal/response"
)

/*
* Responses
 */
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(getSuccessResponse(c, data))
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
func getSuccessResponse(c *fiber.Ctx, data interface{}) response.Response {
	return response.Response{
		Data: data,
		Meta: generateMeta(c),
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
	parsedURL, _ := url.Parse(generateSelfLink(c))
	currentPage := getCurrentPage(c)

	return response.Meta{
		Page:            0,
		Self:            parsedURL.String(),
		Next:            generateNextLink(c, parsedURL, currentPage),
		Previous:        generatePreviousLink(c, parsedURL, currentPage),
		TotalPages:      0,
		TotalItemsCount: 0,
	}
}

func generateSelfLink(c *fiber.Ctx) string {
	return c.BaseURL() + c.OriginalURL()
}

func getCurrentPage(c *fiber.Ctx) int {
	numPage := c.Query("page")
	currentPage, _ := strconv.Atoi(numPage)

	if currentPage == 0 {
		currentPage = 1
	}

	return currentPage
}

func generateNextLink(c *fiber.Ctx, parsedUrl *url.URL, currentPage int) string {
	// TODO: check if the currentPage not bigger then last page

	values, _ := url.ParseQuery(parsedUrl.RawQuery)
	values.Set("page", strconv.Itoa(currentPage+1))

	parsedUrl.RawQuery = values.Encode()

	return parsedUrl.String()
}

func generatePreviousLink(c *fiber.Ctx, parsedUrl *url.URL, currentPage int) string {
	prevPage := currentPage - 1

	if prevPage < 1 {
		return ""
	}

	values, _ := url.ParseQuery(parsedUrl.RawQuery)
	values.Set("page", strconv.Itoa(prevPage))

	parsedUrl.RawQuery = values.Encode()

	return parsedUrl.String()
}
