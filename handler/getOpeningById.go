package handler

import (
	"net/http"

	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningByIdHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		logger.ErrorF("Error finding opening: %v", err)

		return
	}

	sendSucess(ctx, "Get Opening", opening)

}
