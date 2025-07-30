package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"timecassette_api/models"
	"timecassette_api/repository"
)

func SearchEmail() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.SearchEmailRequestModel
		responseModel := models.SearchEmailResonseModel{Meta: models.MetaBaseModel{Result: false}}

		queryParams := c.Request.URL.Query()
		requestPayload.Email = queryParams.Get("branchId")
		if requestPayload.Email == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{},
				Result:   true,
			}
			c.JSON(http.StatusOK, responseModel)
			return
		}
		//Find All
		statusCode := http.StatusOK

		list, ok := repository.SearchEmails(ctx, cancel, requestPayload.Email)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{},
			Result:   ok,
		}
		if !ok {
			statusCode = http.StatusServiceUnavailable
		} else {
			responseModel.Data = list
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
	}
}
