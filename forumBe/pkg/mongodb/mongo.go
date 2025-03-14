package mongodb

import (
	"context"
	"forum/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDBClient *mongo.Client
	AppMode       string
	HttpPort      string
	Db            string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassWord    string
	DbName        string

	MongoDBName string = ""
	MongoDBAddr string = ""
	MongoDBPwd  string = ""
	MongoDBPort string = "27017"
)

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost" + ":" + "27017")
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logger.Info("MongoDB Connect Failed")
	}
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logger.Info("MongoDB Connect Failed")
	}
	logger.Info("MongoDB Connect")
}
