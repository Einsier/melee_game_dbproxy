package api

import (
	"net"
	"net/http"
	"net/rpc"
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
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	// 启动服务
	_ = http.Serve(listen, nil)
}
