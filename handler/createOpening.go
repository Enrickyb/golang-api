package handler

import (
	"net/http"

	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		logger.ErrorF("Error validating request: %v", err)

		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		logger.ErrorF("Error creating opening: %v", err)

		return
	}

	sendSucess(ctx, "Create Opening", opening)

}
