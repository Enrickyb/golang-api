package handler

import (
	"net/http"

	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		logger.ErrorF("Error finding opening: %v", err)

		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		logger.ErrorF("Error deleting opening: %v", err)

		return
	}

	sendSucess(ctx, "Delete", opening)

}
