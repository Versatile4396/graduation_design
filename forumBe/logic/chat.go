package logic

import (
	"context"
	"forum/models"
	"forum/pkg/mongodb"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func GetHistoryChatMsg(uid uint64) (results map[string][]models.Trainer, err error) {
	// 获取所有集合名字
	collectionNames, err := mongodb.MongoDBClient.Database(mongodb.MongoDBName).ListCollectionNames(context.TODO(), bson.D{})
	strUid := strconv.FormatUint(uid, 10)
	results = make(map[string][]models.Trainer)
	for _, collectName := range collectionNames {
		if strings.Contains(collectName, strUid) {
			collection := mongodb.MongoDBClient.Database(mongodb.MongoDBName).Collection(collectName)
			cursor, err := collection.Find(context.TODO(), bson.D{})
			if err != nil {
				continue
			}
			var res []models.Trainer
			err = cursor.All(context.TODO(), &res)
			if err != nil {
				continue
			}
			results[collectName] = res
		}
	}
	return results, nil
}
