package utils

import "github.com/labstack/echo/v4"

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(ctx echo.Context, status int, err error) {
	er := HttpError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
