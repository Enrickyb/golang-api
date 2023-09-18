package handler

import (
	"fmt"
	"net/http"

	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{"msg": msg, "Error code": code})
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s sucessfully", op), "data": data})
}

type ErrorResponse struct {
	Message   string `json:"msg"`
	ErrorCode string `json:"error_code"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"msg"`
	Data    schemas.OpeningResponse `json:"data"`
}
