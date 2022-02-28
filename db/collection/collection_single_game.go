package collection

import (
	"melee_game_dbproxy/model"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type singleGameCollection struct {
	*BaseCollection
}

var singleGameColl *singleGameCollection

func GetSingleGameCollection() (*singleGameCollection, error) {
	if singleGameColl == nil {
		singleGameColl = &singleGameCollection{
			BaseCollection: NewBaseCollection("single_game"),
		}
	}
	return singleGameColl, nil
}

// InsertItem 新增对局信息
func (singleGameColl *singleGameCollection) InsertItem(item interface{}) (string, error) {
	go logrus.Debug("Inserting item...")
	singleGame := item.(*model.SingleGame)
	singleGame.ID = primitive.NewObjectID()
	objectId, err := singleGameColl.BaseCollection.InsertItem(singleGame)
	if err != nil {
		return "", err
	}
	return objectId, nil
}
