package collection

import (
	"melee_game_dbproxy/configs"
	"melee_game_dbproxy/db"
	"melee_game_dbproxy/model"
	"testing"
	"time"
)

func TestInsertPlayer(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)

	playerColl, err := GetPlayerCollection()
	if err != nil {
		t.Error("插入时发生了错误", err)
		return
	}
	id, err := playerColl.InsertItem(&model.Player{
		PlayerId:  0,
		AccountId: "6211f6f0822762c86c7e0ed7",
		Nickname:  "player0",
		CreateAt:  time.Now().UnixNano(),
		UpdateAt:  time.Now().UnixNano(),
	})
	if err != nil {
		t.Error("插入时发生了错误", err)
		return
	}
	t.Log("获取的ObjectID为", id)
}

func TestFindPlayerByPlayerId(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	playerColl, _ := GetPlayerCollection()
	player, err := playerColl.FindItemsByKey([]*MatchItem{
		{
			Key:      "player_id",
			MatchVal: 560,
		},
	})
	if err != nil {
		t.Error("查找时发生错误", err)
	}
	t.Log("查找到的player为", player[0])
}
