package views

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func CreatedResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

// Custom error responses

func InvalidParams(c *fiber.Ctx) error {
	return ErrorResponse(c, 400, "Invalid parameters : The request was invalid and was unable to be processed")
}

func InternalServerError(c *fiber.Ctx, err error) error {
	log.Println(err.Error())
	return ErrorResponse(c, 500, "Internal Server Error: Something went wrong on the server side.")
}

func RecordNotFound(c *fiber.Ctx) error {
	return ErrorResponse(c, 404, "Not Found: The requested resource was not found.")
}

func Unauthorized(c *fiber.Ctx) error {
	return ErrorResponse(c, 401, "Unauthorized: You don't have permission to access this resource.")
}

func Forbidden(c *fiber.Ctx) error {
	return ErrorResponse(c, 403, "Forbidden: You don't have permission to access this resource.")
}

func BadRequest(c *fiber.Ctx) error {
	return ErrorResponse(c, 400, "Bad request : The request was invalid and was unable to be processed")
}

func Conflict(c *fiber.Ctx) error {
	return ErrorResponse(c, 409, "Conflict: The request could not be completed due to a conflict with the current state of the target resource.")
}
