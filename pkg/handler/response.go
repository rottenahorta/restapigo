package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Msg string `json:"msg"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(statusCode, errorResponse{msg})
}