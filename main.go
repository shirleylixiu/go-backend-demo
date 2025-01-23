package main

import (
	"context"
	"go-backend-demo/config"
	"go-backend-demo/controllers"
	_ "go-backend-demo/controllers"
	"go-backend-demo/routes"
	"go-backend-demo/services"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx            context.Context
	server         *gin.Engine
	authService    services.AuthService
	authController controllers.AuthController
	authRoute      routes.AuthRoute
)

func main() {

	log.Printf("Starting....\n")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	log.Printf("config load success: %v\n", config)
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoClient, err := mongo.Connect(ctx, mongoconn)

	// init service, controller, router
	authService = services.NewAuthService(ctx, config, mongoClient)
	authController = controllers.NewAuthController(ctx, authService)
	authRoute = routes.NewAuthRoute(authController)

	startGinServer(config)

}

func startGinServer(config config.Config) {
	server = gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.Origin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
	})
	authRoute.AuthRoute(router)
	log.Fatal(server.Run(":" + config.Port))
}
