package api

import (
	"log"
	"net/rpc"
	"testing"
	"time"
)

var (
	address = "127.0.0.1:1234"
)

func TestIsAccountLegal(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialHttp error", err)
		return
	}

	req := &IsAccountLegalRequest{
		Phone:    "17306409322",
		Password: "123456",
	}
	resp := &IsAccountLegalResponse{}
	err = client.Call("HallHandler.IsAccountLegal", req, resp) //阻塞三秒

	if err != nil {
		log.Fatal("call error", err)
	}
	log.Println("获取的结果为", resp)
}

func TestSearchPlayerInfo(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialHttp error", err)
		return
	}

	req := &SearchPlayerInfoRequest{
		PlayerId: 0,
	}
	resp := &SearchPlayerInfoResponse{}
	err = client.Call("HallHandler.SearchPlayerInfo", req, resp) //阻塞三秒

	if err != nil {
		log.Fatal("call error", err)
	}
	log.Println("获取的结果为", resp)
}

func TestUpdatePlayerInfo(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialHttp error", err)
		return
	}

	req := &UpdatePlayerInfoRequest{
		Info: &PlayerInfo{
			PlayerId:  0,
			NickName:  "player0",
			GameCount: 1,
			KillNum:   5,
			MaxKill:   3,
		},
	}
	resp := &UpdatePlayerInfoResponse{}
	err = client.Call("HallHandler.UpdatePlayerInfo", req, resp) //阻塞三秒

	if err != nil {
		log.Fatal("call error", err)
	}
	log.Println("获取的结果为", resp)
}

func TestAddSingleGameInfo(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Fatal("dialHttp error", err)
		return
	}

	req := &AddSingleGameInfoRequest{
		Info: &SingleGameInfo{
			Players:   []int32{0, 1, 2},
			StartTime: time.Now().UnixNano(),
			EndTime:   time.Now().UnixNano(),
		},
	}
	resp := &AddSingleGameInfoResponse{}
	err = client.Call("HallHandler.AddSingleGameInfo", req, resp) //阻塞三秒

	if err != nil {
		log.Fatal("call error", err)
	}
	log.Println("获取的结果为", resp)
}
