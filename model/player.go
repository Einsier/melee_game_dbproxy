package model

type Player struct {
	PlayerId  int32  `bson:"player_id"`  //玩家id
	AccountId string `bson:"account_id"` //账户id
	Nickname  string `bson:"nickname"`   //玩家昵称
	GameCount int32  `bson:"game_count"` //参与游戏总局数
	KillNum   int32  `bson:"kill_num"`   //总击杀数
	MaxKill   int32  `bson:"max_kill"`   //最高单局击杀数
	CreateAt  int64  `bson:"create_at"`  //创建时间
	UpdateAt  int64  `bson:"update_at"`  //更新时间
}
