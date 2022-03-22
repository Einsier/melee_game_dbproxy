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
	DBAddr     string
)

func initParameter() {
	flag.StringVar(&DBUser, "DBUser", "", "User name of database")
	flag.StringVar(&DBPassword, "DBPassword", "", "Password of database user")
	flag.StringVar(&DBAddr, "DBAddr", "", "address(IP:Port) of database")

	flag.Parse()
	if DBUser != "" {
		configs.MongoURI = "mongodb://" + DBUser + ":" + DBPassword + "@" + DBAddr
	} else {
		configs.MongoURI = "mongodb://" + DBAddr
	}
}

func main() {
	initParameter()
	db.InitMongoConn(configs.MongoURI) //初始化数据库
	handler := &api.Handler{}
	handler.Start() //启动 rpc 服务
}
