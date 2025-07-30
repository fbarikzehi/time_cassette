package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"timecassette_api/models"
	"timecassette_api/repository"
)

func CreateCassette() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.CassetteCreateRequestModel
		responseModel := models.CassetteCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Name == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Name is required"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		userCredential, ok := repository.FindUserCredentialByToken(c, ctx, cancel)
		if !ok {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"User dose not exist"},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
			return
		}

		//Create
		statusCode := http.StatusOK

		createdID, message, result := repository.CreateCassette(requestPayload, ctx, cancel, userCredential.UserId)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{message},
			Result:   result,
		}
		if !result {
			statusCode = http.StatusServiceUnavailable
		} else {
			responseModel.Data = models.IdBaseModel{Id: createdID}
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
		return
	}
}

func UpdateCassette() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.CassetteUpdateRequestModel
		responseModel := models.CassetteUpdateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == "" || requestPayload.Name == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Name is required"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		userCredential, ok := repository.FindUserCredentialByToken(c, ctx, cancel)
		if !ok {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"User dose not exist"},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
			return
		}

		//Update
		statusCode := http.StatusOK

		message, result := repository.UpdateCassette(requestPayload, ctx, cancel, userCredential.UserId)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{message},
			Result:   result,
		}
		if !result {
			statusCode = http.StatusServiceUnavailable
		} else {
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
		return
	}
}

func DeleteCassette() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.CassetteDeleteRequestModel
		responseModel := models.CassetteCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Cassette id is required"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		userCredential, ok := repository.FindUserCredentialByToken(c, ctx, cancel)
		if !ok {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"User dose not exist"},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
			return
		}

		//Delete
		statusCode := http.StatusOK

		message, result := repository.DeleteCassette(requestPayload, ctx, cancel, userCredential.UserId)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{message},
			Result:   result,
		}
		if !result {
			statusCode = http.StatusServiceUnavailable
		} else {
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
		return
	}
}

func GetAllCassettesByUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		responseModel := models.CassetteGetAllResponseModel{Meta: models.MetaBaseModel{Result: false}}

		userCredential, ok := repository.FindUserCredentialByToken(c, ctx, cancel)
		if !ok {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"User dose not exist"},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
			return
		}

		//Find All
		statusCode := http.StatusOK

		list, message, ok := repository.FindCassettes(ctx, cancel, userCredential.UserId)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{message},
			Result:   ok,
		}
		if !ok {
			statusCode = http.StatusServiceUnavailable
		} else {
			responseModel.Data = list
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
		return
	}
}
