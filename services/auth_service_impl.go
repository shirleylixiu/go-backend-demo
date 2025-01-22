package services

import (
	"context"
	"go-backend-demo/models"
	"go-backend-demo/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthService(ctx context.Context, collection *mongo.Collection) AuthService {
	return &AuthServiceImpl{collection, ctx}
}

func (uc *AuthServiceImpl) SignUp(user *models.SignUpRequest) (*models.SignResponse, error) {
	log.Printf("Begin to signUp\n")
	mockedPwd, _ := utils.HashPassword("1234")
	mockedRes := &models.SignResponse{
		ID:       primitive.NewObjectID(),
		Name:     "Mocked User",
		Email:    "mockedUser@someemail.com",
		Password: mockedPwd,
	}
	return mockedRes, nil
}

func (uc *AuthServiceImpl) SignIn(*models.SignInRequest) (*models.SignResponse, error) {
	log.Printf("Begin to signIn\n")
	mockedPwd, _ := utils.HashPassword("1234")
	mockedRes := &models.SignResponse{
		ID:       primitive.NewObjectID(),
		Name:     "Mocked User",
		Email:    "mockedUser@someemail.com",
		Password: mockedPwd,
	}
	return mockedRes, nil
}
