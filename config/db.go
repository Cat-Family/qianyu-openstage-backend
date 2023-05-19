package config

import (
	"context"
	"log"
	"qianyu/openstage/controllers"
	"qianyu/openstage/services"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Us          services.UserService
	Uc          controllers.UserController
	Ctx         context.Context
	Userc       *mongo.Collection
	Mongoclient *mongo.Client
	err         error
)

// connect to mysql
func ConnectMysql() {}

// connect to mongodb
func ConnectMongodb() {
	// load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Ctx = context.TODO()
	// uri := os.Getenv("MONGODB_URI")
	uri := "mongodb://localhost:27017"
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	Mongoclient, err = mongo.Connect(Ctx, clientOptions)
	log.Println("Connected to established!")

	Userc = Mongoclient.Database("qianyudb").Collection("users")
	Us = services.NewUserService(Userc, Ctx)
	Uc = controllers.New(Us)
}

// connect to redis
func ConnectRedis() {}
