package logic

import (
	"context"
	"forum/global"
	"forum/models"
	"forum/pkg/mongodb"
	"sort"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHistoryChatMsg(uid uint64) (results map[string][]models.Trainer, err error) {
	// 获取所有集合名字
	collectionNames, err := mongodb.MongoDBClient.Database(mongodb.MongoDBName).ListCollectionNames(context.TODO(), bson.D{})
	strUid := strconv.FormatUint(uid, 10)
	results = make(map[string][]models.Trainer)
	for _, collectName := range collectionNames {
		if strings.Contains(collectName, strUid) {
			collection := mongodb.MongoDBClient.Database(mongodb.MongoDBName).Collection(collectName)
			cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"startTime", 1}}).SetLimit(1))
			if err != nil {
				continue
			}
			var res []models.Trainer
			err = cursor.All(context.TODO(), &res)
			if err != nil {
				continue
			}
			chatUids := strings.Split(collectName, "->")
			if chatUids[0] == strUid {
				results[chatUids[1]] = append(results[chatUids[1]], res[0])
			} else {
				results[chatUids[0]] = append(results[chatUids[0]], res[0])
			}
		}
	}

	for key, val := range results {
		sort.Slice(val, func(i, j int) bool {
			return val[i].StartTime < val[j].StartTime
		})
		results[key] = val[:1]
	}
	return results, nil
}

type chatInstance struct {
	latestMsg string
	time      int64
	userInfo  models.UserInfo
}

func GetUserInfoChatMsg(results map[string][]models.Trainer) (chatList []chatInstance, err error) {
	uinfo := models.User{}
	for key, value := range results {
		global.Db.Model(models.User{}).Where("user_id", key).First(&uinfo)
		chatList = append(chatList, chatInstance{
			latestMsg: value[0].Content,
			time:      value[0].StartTime,
			userInfo: models.UserInfo{
				UserId:   uinfo.UserId,
				Avatar:   uinfo.Avatar,
				UserName: uinfo.UserName,
				Gender:   uinfo.Gender,
			},
		})
	}
	return chatList, nil
}
