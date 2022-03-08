package main

import (
	"flag"
	"melee_game_dbproxy/api"
	"melee_game_dbproxy/configs"
	"melee_game_dbproxy/db"
)

var (
	//DBName string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
)

func initParameter() {
	flag.StringVar(&DBUser, "DBUser", "", "User name of database")
	flag.StringVar(&DBPassword, "DBPassword", "", "Password of database user")
	flag.StringVar(&DBHost, "DBHost", "", "IP address of database")
	flag.StringVar(&DBPort, "DBPort", "", "Port to connect to database")

	flag.Parse()

	configs.MongoURI = "mongodb://" + DBUser + ":" + DBPassword + "@" + DBHost + ":" + DBPort + "/" + configs.DBName
}

func main() {
	initParameter()
	db.InitMongoConn(configs.MongoURI) //初始化数据库
	handler := &api.Handler{}
	go handler.Start() //启动 rpc 服务
}
