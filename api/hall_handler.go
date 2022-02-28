package api

import (
	"melee_game_dbproxy/db/collection"
	"melee_game_dbproxy/model"
	"time"

	"github.com/sirupsen/logrus"
)

type HallHandler struct{}

func (h *HallHandler) IsAccountLegal(req *IsAccountLegalRequest, resp *IsAccountLegalResponse) error {
	accountColl, err := collection.GetAccountCollection()
	if err != nil {
		GetErrorIsAccountLegalResponse(resp, DBInnerError)
		return nil
	}
	go logrus.Debug("Finding account by phone...\n")
	accounts, err := accountColl.FindItemsByKey([]*collection.MatchItem{
		{
			Key:      "phone",
			MatchVal: req.Phone,
		},
	})
	if err != nil {
		GetErrorIsAccountLegalResponse(resp, DBInnerError)
		return nil
	}
	if len(accounts) == 0 {
		go logrus.Debug("Fail to find account by phone number\n")
		GetErrorIsAccountLegalResponse(resp, PhoneNotExist) //查找不到账户
	} else if accounts[0].Password != req.Password {
		go logrus.Debug("Wrong password\n")
		GetErrorIsAccountLegalResponse(resp, WrongPassword) //账户密码错误
	} else {
		GetOkIsAccountLegalResponse(resp, accounts[0]) //返回账户对应的玩家id
	}
	return nil
}

func (h *HallHandler) SearchPlayerInfo(req *SearchPlayerInfoRequest, resp *SearchPlayerInfoResponse) error {
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		GetErrorSearchPlayerInfoResponse(resp, DBInnerError)
		return nil
	}
	go logrus.Debug("Finding player by id...\n")
	players, err := playerColl.FindItemsByKey([]*collection.MatchItem{
		{
			Key:      "player_id",
			MatchVal: req.PlayerId,
		},
	})
	if err != nil {
		GetErrorSearchPlayerInfoResponse(resp, DBInnerError)
		return nil
	}
	if len(players) == 0 {
		go logrus.Debug("Fail to find player by id\n")
		GetErrorSearchPlayerInfoResponse(resp, PlayerNotExist) //查找不到玩家
	} else {
		GetOkSearchPlayerInfoResponse(resp, players[0]) //返回玩家信息
	}
	return nil
}

func (h *HallHandler) UpdatePlayerInfo(req *UpdatePlayerInfoRequest, resp *UpdatePlayerInfoResponse) error {
	playerColl, err := collection.GetPlayerCollection()
	if err != nil {
		GetErrorUpdatePlayerInfoResponse(resp, DBInnerError, req.Info)
		return nil
	}
	err = playerColl.UpdateItemByKey(
		[]*collection.MatchItem{
			{
				Key:      "player_id",
				MatchVal: req.Info.PlayerId,
			},
		},
		&collection.Operation{
			Op: "$set",
			Items: []*collection.MatchItem{
				{
					Key:      "nickname",
					MatchVal: req.Info.NickName,
				},
				{
					Key:      "game_count",
					MatchVal: req.Info.GameCount,
				},
				{
					Key:      "kill_num",
					MatchVal: req.Info.KillNum,
				},
				{
					Key:      "max_kill",
					MatchVal: req.Info.MaxKill,
				},
				{
					Key:      "update_at",
					MatchVal: time.Now().UnixNano(),
				},
			},
		},
	)
	if err != nil {
		GetErrorUpdatePlayerInfoResponse(resp, DBInnerError, req.Info)
		return nil
	}
	GetOkUpdatePlayerInfoResponse(resp)
	return nil
}

func (h *HallHandler) AddSingleGameInfo(req *AddSingleGameInfoRequest, resp *AddSingleGameInfoResponse) error {
	singleGameColl, err := collection.GetSingleGameCollection()
	if err != nil {
		GetErrorAddSingleGameInfo(resp, DBInnerError)
		return nil
	}
	_, err = singleGameColl.InsertItem(&model.SingleGame{
		Players:   req.Info.Players,
		StartTime: req.Info.StartTime,
		EndTime:   req.Info.EndTime,
		CreateAt:  time.Now().UnixNano(),
	})
	if err != nil {
		GetErrorAddSingleGameInfo(resp, DBInnerError)
		return nil
	}
	GetOkAddSingleGameInfo(resp)
	return nil
}
