package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"timecassette_api/models"
	"timecassette_api/repository"
)

func CreateFragment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.FragmentCreateRequestModel
		responseModel := models.FragmentCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Name == "" || requestPayload.CassetteId == primitive.NilObjectID {
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

		createdID, message, result := repository.CreateFragment(requestPayload, ctx, cancel, userCredential.UserId)
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

func UpdateFragment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.FragmentUpdateRequestModel
		responseModel := models.FragmentUpdateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == "" || requestPayload.Name == "" || requestPayload.CassetteId == primitive.NilObjectID {
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

		message, result := repository.UpdateFragment(requestPayload, ctx, cancel, userCredential.UserId)
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

func DeleteFragment() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.FragmentDeleteRequestModel
		responseModel := models.FragmentCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

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
				Messages: []string{"Fragment id is required"},
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

		message, result := repository.DeleteFragment(requestPayload, ctx, cancel, userCredential.UserId)
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

func GetAllFragmentsByCassette() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.FragmentGetAllCassetteRequestModel
		responseModel := models.FragmentGetAllCassetteResponseModel{Meta: models.MetaBaseModel{Result: false}}

		queryParams := c.Request.URL.Query()
		cassetteId, err := primitive.ObjectIDFromHex(queryParams.Get("cassetteId"))
		if err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		} else {
			requestPayload.CassetteId = cassetteId
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

		//Find All
		statusCode := http.StatusOK

		list, message, ok := repository.FindFragmentsByCassette(ctx, cancel, userCredential.UserId, requestPayload.CassetteId)
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
