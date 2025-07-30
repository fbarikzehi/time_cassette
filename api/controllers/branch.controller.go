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

func CreateBranch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchCreateRequestModel
		responseModel := models.BranchCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Name == "" || requestPayload.HandlerUserEmail == "" || requestPayload.FragmentId == primitive.NilObjectID {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Name and Email is required"},
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

		handlerUser, ok := repository.FindUserByEmail(ctx, cancel, requestPayload.HandlerUserEmail)
		if !ok {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Email not found"},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
			return
		}

		//Create
		statusCode := http.StatusOK
		createdID, message, result := repository.CreateBranch(requestPayload, ctx, cancel, userCredential.UserId, handlerUser.ID)
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

func UpdateBranch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchUpdateRequestModel
		responseModel := models.BranchUpdateResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == "" || requestPayload.Name == "" || requestPayload.FragmentId == primitive.NilObjectID {
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

		message, result := repository.UpdateBranch(requestPayload, ctx, cancel, userCredential.UserId)
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

func ConfirmBranch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchConfirmRequestModel
		responseModel := models.BranchConfirmResponseModel{Meta: models.MetaBaseModel{Result: false}}

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

		//Update
		statusCode := http.StatusOK

		message, result := repository.BranchConfirm(requestPayload, ctx, cancel, userCredential.UserId)
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

func DeleteBranchRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchDeleteRequestModel
		responseModel := models.BranchCreateResponseModel{Meta: models.MetaBaseModel{Result: false}}

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

		message, result := repository.DeleteBranchRequest(requestPayload, ctx, cancel, userCredential.UserId)
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

func DeleteBranchConfirm() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchDeleteConfirmationRequestModel
		responseModel := models.BranchDeleteConfirmationResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Id == "" || requestPayload.SecretCode == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Branch SecretCode is required"},
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

		//Confirm Delete
		statusCode := http.StatusOK

		message, result := repository.DeleteBranchConfirmation(requestPayload, ctx, cancel, userCredential.UserId)
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

func GetAllBranchesByFragment() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.BranchGetAllByFragmentRequestModel
		responseModel := models.BranchGetAllResponseModel{Meta: models.MetaBaseModel{Result: false}}

		queryParams := c.Request.URL.Query()
		fragmentId, err := primitive.ObjectIDFromHex(queryParams.Get("fragmentId"))
		if err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		} else {
			requestPayload.FragmentId = fragmentId
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

		list, message, ok := repository.FindBranchesByFragment(ctx, cancel, userCredential.UserId, requestPayload.FragmentId)
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
