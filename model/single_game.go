package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SingleGame struct {
	ID        primitive.ObjectID `bson:"_id"`
	Players   []int32            `bson:"players"`    //参与玩家
	StartTime int64              `bson:"start_time"` //游戏开始时间
	EndTime   int64              `bson:"end_time"`   //游戏结束时间
	CreateAt  int64              `bson:"create_at"`  //创建时间
}
