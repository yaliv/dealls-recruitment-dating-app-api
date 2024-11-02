package jsonresponse

import (
	"github.com/gofiber/fiber/v2"
)

type (
	SuccessArgs struct {
		HttpStatus int
		Data       fiber.Map
	}

	ErrorArgs struct {
		HttpStatus int
		Error      ErrorProp
	}

	SuccessResponse struct {
		IsSuccess bool      `json:"is_success"`
		Data      fiber.Map `json:"data"`
	}

	ErrorResponse struct {
		IsSuccess bool      `json:"is_success"`
		Error     ErrorProp `json:"error"`
	}

	ErrorProp struct {
		Code    string    `json:"code"`
		Message string    `json:"message"`
		Details fiber.Map `json:"details"`
	}
)

func Success(c *fiber.Ctx, args *SuccessArgs) error {
	if args == nil {
		args = new(SuccessArgs)
	}
	if args.HttpStatus == 0 {
		args.HttpStatus = 200
	}
	if args.Data == nil {
		args.Data = make(fiber.Map)
	}

	res := SuccessResponse{
		IsSuccess: true,
		Data:      args.Data,
	}

	return c.Status(args.HttpStatus).JSON(res)
}

func Error(c *fiber.Ctx, args *ErrorArgs) error {
	if args == nil {
		args = new(ErrorArgs)
	}
	if args.HttpStatus == 0 {
		args.HttpStatus = 500
	}
	if args.Error.Code == "" {
		args.Error.Code = "UNSPECIFIED_ERROR"
	}
	if args.Error.Message == "" {
		args.Error.Message = "Error ini tidak dirinci."
	}
	if args.Error.Details == nil {
		args.Error.Details = make(fiber.Map)
	}

	res := ErrorResponse{
		IsSuccess: false,
		Error:     args.Error,
	}

	return c.Status(args.HttpStatus).JSON(res)
}

func ErrorPayloadSyntax(c *fiber.Ctx, err error) error {
	return Error(c, &ErrorArgs{
		HttpStatus: fiber.StatusBadRequest,
		Error: ErrorProp{
			Code:    "ERR_PAYLOAD_SYNTAX",
			Message: err.Error(),
		},
	})
}

func ErrorReadData(c *fiber.Ctx, err error) error {
	errArgs := &ErrorArgs{
		HttpStatus: fiber.StatusInternalServerError,
		Error: ErrorProp{
			Code:    "ERR_READ_DATA",
			Message: err.Error(),
		},
	}

	return Error(c, errArgs)
}

func ErrorWriteData(c *fiber.Ctx, err error) error {
	errArgs := &ErrorArgs{
		HttpStatus: fiber.StatusInternalServerError,
		Error: ErrorProp{
			Code:    "ERR_WRITE_DATA",
			Message: err.Error(),
		},
	}

	return Error(c, errArgs)
}
