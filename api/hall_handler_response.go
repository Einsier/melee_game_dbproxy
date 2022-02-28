package api

import "melee_game_dbproxy/model"

func GetErrorIsAccountLegalResponse(resp *IsAccountLegalResponse, error ErrorType) {
	resp.Ok = false
	resp.Error = error
}

func GetOkIsAccountLegalResponse(resp *IsAccountLegalResponse, account *model.Account) {
	resp.Ok = true
	resp.PlayerId = account.PlayerId
}

func GetErrorSearchPlayerInfoResponse(resp *SearchPlayerInfoResponse, error ErrorType) {
	resp.Ok = false
	resp.Error = error
}

func GetOkSearchPlayerInfoResponse(resp *SearchPlayerInfoResponse, player *model.Player) {
	resp.Ok = true
	resp.Info = &PlayerInfo{
		PlayerId:  player.PlayerId,
		NickName:  player.Nickname,
		GameCount: player.GameCount,
		KillNum:   player.KillNum,
		MaxKill:   player.MaxKill,
	}
}

func GetErrorUpdatePlayerInfoResponse(resp *UpdatePlayerInfoResponse, error ErrorType, playInfo *PlayerInfo) {
	resp.Ok = false
	resp.Error = error
	resp.Info = playInfo
}

func GetOkUpdatePlayerInfoResponse(resp *UpdatePlayerInfoResponse) {
	resp.Ok = true
}

func GetErrorAddSingleGameInfo(resp *AddSingleGameInfoResponse, error ErrorType) {
	resp.Ok = false
	resp.Error = error
}

func GetOkAddSingleGameInfo(resp *AddSingleGameInfoResponse) {
	resp.Ok = true
}
