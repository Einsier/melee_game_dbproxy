package api

import (
	"melee_game_dbproxy/configs"
	"net"
	"net/http"
	"net/rpc"

	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func (h *Handler) Start() {
	// 功能对象注册
	hallHandler := new(HallHandler)
	err := rpc.Register(hallHandler) //rpc.RegisterName("自定义服务名",encryption)
	if err != nil {
		panic(err.Error())
	}
	// HTTP注册
	rpc.HandleHTTP()

	// 端口监听
	listen, err := net.Listen("tcp", configs.TcpPort)
	if err != nil {
		panic(err.Error())
	}
	go logrus.Info("RPC service is listening on port " + configs.TcpPort)
	// 启动服务
	_ = http.Serve(listen, nil)
}
