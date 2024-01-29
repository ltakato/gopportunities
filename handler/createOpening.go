package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltakato/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	// it's interesting way to fill incoming object from request
	// it also ignores extra not defined fields for example
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "createOpening", opening)
}
