package controllers

import (
	"context"
	"go-backend-demo/models"
	"go-backend-demo/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	ctx         context.Context
	authService services.AuthService
}

func NewAuthController(ctx context.Context, authService services.AuthService) AuthController {
	return AuthController{ctx, authService}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var req *models.SignUpRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Request format error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	newUser, err := ac.authService.SignUp(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.BaseResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{
		StatusCode: 0,
		Message:    "success",
		Data:       newUser,
	})

}

func (ac *AuthController) SignIn(ctx *gin.Context) {
	var req *models.SignInRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Request format error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, models.BaseResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	user, err := ac.authService.SignIn(req)
	if err != nil {
		log.Printf("User sign failed: %v.\n", err.Error())
		ctx.JSON(http.StatusUnauthorized, models.BaseResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponse{
		StatusCode: 0,
		Message:    "success",
		Data:       user,
	})

}
