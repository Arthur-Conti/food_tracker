package controllers

import "github.com/gin-gonic/gin"

type SuccessResponseBody struct {
	StatusCode int `json:"status_code"`
	Data       any `json:"data"`
}

type FailResponseBody struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(code, SuccessResponseBody{
		StatusCode: code,
		Data:       data,
	})
}

func FailResponse(c *gin.Context, code int, message string) {
	c.JSON(code, FailResponseBody{
		StatusCode: code,
		Message:    message,
	})
}
