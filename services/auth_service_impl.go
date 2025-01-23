package services

import (
	"context"
	"errors"
	"go-backend-demo/config"
	"go-backend-demo/models"
	"go-backend-demo/utils"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

const authDbCollectionName = "users"

func NewAuthService(ctx context.Context, config config.Config, mongoClient *mongo.Client) AuthService {
	collection := mongoClient.Database(config.DBName).Collection(authDbCollectionName)
	return &AuthServiceImpl{collection, ctx}
}

func (authService *AuthServiceImpl) SignUp(user *models.SignUpRequest) (*models.SignResponse, error) {
	log.Println("Begin to sign up user")
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"

	if encodedPwd, err := utils.HashPassword(user.Password); err != nil {
		log.Printf("Hash password failed! %v\n", err.Error())
		return nil, err
	} else {
		user.Password = encodedPwd
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := authService.collection.Indexes().CreateOne(authService.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	res, err := authService.collection.InsertOne(authService.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	var newUser *models.SignResponse
	query := bson.M{"_id": res.InsertedID}

	err = authService.collection.FindOne(authService.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}

func (authService *AuthServiceImpl) SignIn(signInRequest *models.SignInRequest) (*models.SignResponse, error) {

	query := bson.M{"email": strings.ToLower(signInRequest.Email)}
	opt := options.FindOptions{}
	opt.SetLimit(int64(1))
	opt.SetSkip(int64(0))
	cur, err := authService.collection.Find(authService.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cur.Close(authService.ctx)

	var user *models.SignResponse

	if !cur.Next(authService.ctx) {
		return nil, errors.New("email or password error")
	}

	if err = cur.Decode(&user); err != nil {
		return nil, err
	}

	// Another way to query user. FindOne() will return error if user not exist
	// if err := authService.collection.FindOne(authService.ctx, query).Decode(&user); err != nil {
	// 	return nil, err
	// }

	if err := utils.VerifyPassword(user.Password, signInRequest.Password); err != nil {
		return nil, errors.New("email or password error")
	} else {
		return user, nil
	}
}
