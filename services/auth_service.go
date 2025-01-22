package services

import "go-backend-demo/models"

type AuthService interface {
	SignUp(*models.SignUpRequest) (*models.SignResponse, error)
	SignIn(*models.SignInRequest) (*models.SignResponse, error)
}
