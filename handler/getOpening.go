package handler

import (
	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, 500, err.Error())
		logger.ErrorF("Error finding openings: %v", err)

		return
	}

	sendSucess(ctx, "Get Openings", openings)

}
