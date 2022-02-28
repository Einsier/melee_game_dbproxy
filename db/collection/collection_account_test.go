package collection

import (
	"fmt"
	"log"
	"melee_game_dbproxy/configs"
	"melee_game_dbproxy/model"
	"os"
	"testing"
	"time"

	"melee_game_dbproxy/db"
)

func TestInsertAccount(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)

	accountColl, _ := GetAccountCollection()
	id, err := accountColl.InsertItem(&model.Account{
		AccountName: "test",
		Phone:       "17306409322",
		Password:    "123456",
		CreateAt:    time.Now().UnixNano(),
		UpdateAt:    time.Now().UnixNano(),
		Delete:      false,
	})
	if err != nil {
		log.Println("插入时发生了错误", err)
		return
	}
	log.Println("获取的ID为", id)
}

func TestFindAccountById(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	acc, err := accountColl.FindOneItemById("620cf57be1c94ffd2f98e309")
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	fmt.Println("获取的account为", acc)
}

func TestFindAccountByKey(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	acc, err := accountColl.FindItemsByKey([]*MatchItem{
		{
			Key:      "account_name",
			MatchVal: "test",
		},
	})
	if err != nil {
		t.Error("查找时发生了错误", err)
	}
	fmt.Println("获取的acc为", acc[0])
}

func TestDeleteAccountById(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	err := accountColl.DeleteItemById("620cf57be1c94ffd2f98e309")
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestDeleteAccountByKey(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	err := accountColl.DeleteItemByKey([]*MatchItem{})
	if err != nil {
		t.Error("删除时发生了错误", err)
	}
}

func TestUpdateAccountById(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	err := accountColl.UpdateItemById(
		"620cf57be1c94ffd2f98e309",
		&Operation{
			Op: "$set",
			Items: []*MatchItem{
				{
					Key:      "account_name",
					MatchVal: "chen",
				},
			},
		},
	)
	if err != nil {
		t.Error("更新时发生了错误", err)
	}
}

func TestUpdateAccountByKey(t *testing.T) {
	skipCI(t)
	db.InitMongoConn(configs.MongoURIForTest)
	accountColl, _ := GetAccountCollection()
	err := accountColl.UpdateItemByKey(
		[]*MatchItem{
			{
				Key:      "account_name",
				MatchVal: "test",
			},
		},
		&Operation{
			Op: "$set",
			Items: []*MatchItem{
				{
					Key:      "account_name",
					MatchVal: "chen",
				},
				{
					Key:      "update_at",
					MatchVal: time.Now().UnixNano(),
				},
			},
		},
	)
	if err != nil {
		t.Error("更新时发生了错误", err)
	}
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}
