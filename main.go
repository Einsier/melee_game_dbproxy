package main

import (
	"melee_game_dbproxy/api"
	"melee_game_dbproxy/configs"
	"melee_game_dbproxy/db"
)

func main() {
	db.InitMongoConn(configs.MongoURI) //初始化数据库
	handler := &api.Handler{}
	go handler.Start() //启动 rpc 服务
}
