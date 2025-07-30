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

func CreateTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeCreateRequestModel
		responseModel := models.TimeCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.BranchId == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Branch is required"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}
		if time.Time.IsZero(requestPayload.StartDateTime) {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Start date and time is required"},
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
		createdID, message, result := repository.CreateTimeInstant(requestPayload, ctx, cancel, userCredential.UserId)
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

func UpdateTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeUpdateRequestModel
		responseModel := models.TimeUpdateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == primitive.NilObjectID || requestPayload.BranchId == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Data is required"},
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

		message, result := repository.UpdateTime(requestPayload, ctx, cancel, userCredential.UserId)
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

func UpdateTimeDescription() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeUpdateDescriptionRequestModel
		responseModel := models.TimeUpdateDescriptionResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Data is required"},
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

		message, result := repository.UpdateTimeDescription(requestPayload, ctx, cancel, userCredential.UserId)
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

func UpdateStartTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeStartRequestModel
		responseModel := models.TimeStartResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Data is required"},
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

		//Update Start time
		statusCode := http.StatusOK

		message, result := repository.UpdateStartTime(requestPayload, ctx, cancel, userCredential.UserId)
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

func UpdateEndTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeEndRequestModel
		responseModel := models.TimeEndResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Data is required"},
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

		//Update End time
		statusCode := http.StatusOK
		message, result := repository.UpdateEndTime(ctx, cancel, userCredential.UserId, requestPayload)
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

func DeleteTime() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeDeleteRequestModel
		responseModel := models.TimeDeleteResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Branch id is required"},
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

		message, result := repository.DeleteTime(requestPayload, ctx, cancel, userCredential.UserId)
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

func DeleteAllTime() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeDeleteAllRequestModel
		responseModel := models.TimeDeleteAllResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.BranchId == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Branch is required"},
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

		//Delete All Times By Branch
		statusCode := http.StatusOK

		message, result := repository.DeleteAllTimesByBranch(ctx, cancel, userCredential.UserId, requestPayload.BranchId)
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

func GetAllTimesByBranch() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.TimeGetAllByBranchRequestModel
		responseModel := models.TimeGetAllByBranchResponseModel{Meta: models.MetaBaseModel{Result: false}}

		queryParams := c.Request.URL.Query()
		branchId, err := primitive.ObjectIDFromHex(queryParams.Get("branchId"))
		if err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		} else {
			requestPayload.BranchId = branchId
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

		branchTimes, message, ok := repository.FindTimesByBranch(ctx, cancel, userCredential.UserId, requestPayload.BranchId)
		responseModel.Meta = models.MetaBaseModel{
			Messages: []string{message},
			Result:   ok,
		}
		if !ok {
			statusCode = http.StatusServiceUnavailable
		} else {
			responseModel.Data = branchTimes
			statusCode = http.StatusOK
		}
		c.JSON(statusCode, responseModel)
		return
	}
}
