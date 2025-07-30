package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"timecassette_api/models"
	"timecassette_api/repository"
	"timecassette_api/utils"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.LoginRequestModel
		var user repository.User
		responseModel := models.LoginResponseModel{Meta: models.MetaBaseModel{Result: false}}

		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Email == "" || requestPayload.Password == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Required:Email and Password"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		var usersCollection *mongo.Collection = repository.OpenCollection(repository.Client, "users")

		filter := bson.M{"email": requestPayload.Email}
		err := usersCollection.FindOne(ctx, filter).Decode(&user)
		if err == mongo.ErrNoDocuments {
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Not found:You must register first."},
				Result:   false,
			}
			c.JSON(http.StatusUnauthorized, responseModel)
		} else {
			userCredential, ok := repository.FindSigninUserCredentialByUserId(ctx, cancel, user.ID)
			if !ok {
				defer cancel()
				responseModel.Meta = models.MetaBaseModel{
					Messages: []string{"Credentials Not found:You must register first."},
					Result:   false,
				}
				c.JSON(http.StatusUnauthorized, responseModel)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(userCredential.HashedPassword), []byte(requestPayload.Password))
			if err != nil {
				responseModel.Meta = models.MetaBaseModel{
					Messages: []string{"Password is incorrect"},
					Result:   false,
				}
				c.JSON(http.StatusBadRequest, responseModel)
			} else {

				//Generate New Token
				var userForCredential repository.User
				expiresAt := userCredential.TokenExpire
				token := userCredential.Token
				if token == "" || userCredential.TokenExpire.Unix() <= time.Now().Unix() {
					tokeExist := true
					for tokeExist == true {
						token, err = utils.GenerateToken(requestPayload.Email)
						if err == nil {
							usersCollection.FindOne(ctx, bson.M{"token": token}).Decode(&userForCredential)
							defer cancel()
							if userForCredential.ID == primitive.NilObjectID {
								tokeExist = false
							}
						}
					}

					expiresAt = time.Now().Add(time.Hour * 8)
					ok := repository.UpdateUserCredentialToken(ctx, cancel, user.ID, token, expiresAt)
					if !ok {
						defer cancel()
						responseModel.Meta = models.MetaBaseModel{
							Messages: []string{"Credentials Not found:You must register first."},
							Result:   false,
						}
						c.JSON(http.StatusUnauthorized, responseModel)
						return
					}
				}
				responseModel.Meta = models.MetaBaseModel{
					Messages: []string{"Welcome to Time Cassette"},
					Result:   true,
				}
				responseModel.Data = models.LoginResponseDataModel{Token: token, ExpireAt: expiresAt, ReturnUrl: userCredential.ReturnUrl}
				c.JSON(http.StatusOK, responseModel)
			}

		}
		defer cancel()
		return
	}
}
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var requestPayload models.SignupRequestModel
		var user repository.User
		responseModel := models.SignupResponseModel{Meta: models.MetaBaseModel{Result: false}}

		//validations
		if err := c.BindJSON(&requestPayload); err != nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"something wrong with the data"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if requestPayload.Email == "" || requestPayload.Password == "" || requestPayload.ConfirmPassword == "" {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Email,Password and ConfirmPassword are required"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if strings.Contains("", requestPayload.Email) || strings.Contains("", requestPayload.Password) || strings.Contains("", requestPayload.ConfirmPassword) {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Email,Password and Confirm Password is required and empty spaces are not allowed"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		if strings.Compare(requestPayload.Password, requestPayload.ConfirmPassword) != 0 {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Password and Confirm Password Confirm not matched"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
			return
		}

		var usersCollection *mongo.Collection = repository.OpenCollection(repository.Client, "users")
		if usersCollection == nil {
			defer cancel()
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Sorry! Server side error.We will fix this ASAP"},
				Result:   false,
			}
			c.JSON(http.StatusServiceUnavailable, responseModel)
			return
		}

		//signup process
		filter := bson.M{"email": requestPayload.Email}
		err := usersCollection.FindOne(ctx, filter).Decode(&user)
		defer cancel()
		if err == mongo.ErrNoDocuments {
			//generate unique token
			token := ""
			tokeExist := true
			for tokeExist == true {
				token, err = utils.GenerateToken(requestPayload.Email)
				if err == nil {
					usersCollection.FindOne(ctx, bson.M{"token": token}).Decode(&user)
					defer cancel()
					if user.ID == primitive.NilObjectID {
						tokeExist = false
					}
				}
			}

			user.ID = primitive.NewObjectID()
			user.Email = requestPayload.Email
			user.CreateDateTime = time.Now()
			user.UpdateDateTime = time.Now()
			insertOneResult, err := usersCollection.InsertOne(ctx, user)
			if err != nil || insertOneResult.InsertedID == nil {
				defer cancel()
				responseModel.Meta = models.MetaBaseModel{
					Messages: []string{"Sorry! Server side error.We will fix this ASAP"},
					Result:   false,
				}
				c.JSON(http.StatusServiceUnavailable, responseModel)
			} else {
				oid, _ := insertOneResult.InsertedID.(primitive.ObjectID)

				var userCredential repository.UserCredential
				userCredential.ID = primitive.NewObjectID()
				userCredential.UserId = oid
				userCredential.Username = requestPayload.Email
				userCredential.HashedPassword = utils.HashPassword(requestPayload.Password)
				userCredential.Role = "USER"
				userCredential.Token = token
				userCredential.TokenExpire = time.Now().Add(time.Hour + 8)
				userCredential.CreateDateTime = time.Now()
				userCredential.UpdateDateTime = time.Now()
				ok := repository.CreateUserCredential(ctx, cancel, userCredential)
				if !ok {
					filter := bson.M{"_id": oid}
					_, err := usersCollection.DeleteOne(ctx, filter)
					if err == nil {
						defer cancel()
						responseModel.Meta = models.MetaBaseModel{
							Messages: []string{"Sorry! User credentials error happened. Please try again"},
							Result:   false,
						}
						c.JSON(http.StatusServiceUnavailable, responseModel)
					} else {
						defer cancel()
						responseModel.Meta = models.MetaBaseModel{
							Messages: []string{"Sorry! User credentials server side error.We will fix this ASAP "},
							Result:   false,
						}
						c.JSON(http.StatusInternalServerError, responseModel)
					}
					return
				}

				responseModel.Meta = models.MetaBaseModel{
					Messages: []string{"Congratulations! You have signed up successfully. Now You can login to Time Cassette"},
					Result:   true,
				}

				c.JSON(http.StatusOK, responseModel)
			}

		} else {
			responseModel.Meta = models.MetaBaseModel{
				Messages: []string{"Email signed up already"},
				Result:   false,
			}
			c.JSON(http.StatusBadRequest, responseModel)
		}
		defer cancel()
		return
	}
}
