package handler

import (
	"net/http"

	"github.com/Enrickyb/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create a new opening
// @Description Create a new opening
// @Accept  json
// @Produce  json
// @Param opening body CreateOpeningRequest true "Create Opening"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]

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
