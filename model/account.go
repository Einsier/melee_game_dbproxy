package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID          primitive.ObjectID `bson:"_id"`
	PlayerId    int32              `bson:"player_id"`    //玩家id
	AccountName string             `bson:"account_name"` //账户名称
	Phone       string             `bson:"phone"`        //手机号
	Password    string             `bson:"password"`     //密码
	RecentLogin int64              `bson:"recent_login"` //最近登录时间
	CreateAt    int64              `bson:"create_at"`    //创建时间
	UpdateAt    int64              `bson:"update_at"`    //更新时间
	Delete      bool               `bson:"delete"`       //是否注销
}
