package routes

import (
	"go-backend-demo/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controllers.AuthController
}

func NewAuthRoute(authController controllers.AuthController) AuthRoute {
	return AuthRoute{authController}
}

func (authRoute *AuthRoute) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/signUp", authRoute.authController.SignUpUser)
	router.POST("/signIn", authRoute.authController.SignIn)

}
