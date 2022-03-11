package db

import (
	"context"
	"melee_game_dbproxy/configs"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConn struct {
	database *mongo.Database
	client   *mongo.Client
}

var mongoConn *MongoConn

// InitMongoConn 初始化数据库连接，获取 client
func InitMongoConn(MongoURI string) {
	go logrus.Debug("Initializing MongoDB connection...")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(MongoURI)
	clientOptions.SetMaxPoolSize(configs.MongoPoolSize) // 连接池配置
	// 连接到MongoDB
	var err error
	cli, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		go logrus.Error(err)
	}
	mongoConn = &MongoConn{
		database: cli.Database(configs.DBName),
		client:   cli,
	}

	// 检查连接
	err = mongoConn.client.Ping(context.TODO(), nil)
	if err != nil {
		go logrus.Error(err)
	}
	go logrus.Info("Connect to " + MongoURI + "/" + configs.DBName)
}

// GetCollection 获取指定集合
func GetCollection(collectionName string, options *options.CollectionOptions) *mongo.Collection {
	return mongoConn.database.Collection(collectionName, options)
}
