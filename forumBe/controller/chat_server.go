package controller

import (
	"context"
	"fmt"
	"forum/models"
	"forum/pkg/mongodb"
	"log"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

const month = 60 * 60 * 24 * 30 // 按照30天算一个月

func FindMany(database string, sendId string, id string, time int64, pageSize int) (results []models.Result, err error) {
	var resultsMe []models.Trainer
	var resultsYou []models.Trainer
	sendIdCollection := mongodb.MongoDBClient.Database(database).Collection(sendId)
	idCollection := mongodb.MongoDBClient.Database(database).Collection(id)
	fmt.Println(sendId, id, pageSize, "pageSize")
	// 如果不知道该使用什么context，可以通过context.TODO() 产生context
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	idTimeCursor, err := idCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		panic(err)
	}
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	if err != nil {
		panic(err)
	}
	err = idTimeCursor.All(context.TODO(), &resultsMe) // Id 发给对面的
	if err != nil {
		panic(err)
	}
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}
func AppendAndSort(resultsMe, resultsYou []models.Trainer) (results []models.Result, err error) {
	for _, r := range resultsMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := models.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			Content:   r.Content,
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultsYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := models.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			Content:   r.Content,
			From:      "you",
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}

func FirsFindtMsg(database string, sendId string, id string) (results []models.Result, err error) {
	// 首次查询(把对方发来的所有未读都取出来)
	var resultsMe []models.Trainer
	var resultsYou []models.Trainer
	sendIdCollection := mongodb.MongoDBClient.Database(database).Collection(sendId)
	idCollection := mongodb.MongoDBClient.Database(database).Collection(id)
	filter := bson.M{"read": bson.M{
		"$all": []uint{0},
	}}
	sendIdCursor, err := sendIdCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{
		"startTime", 1}}), options.Find().SetLimit(1))
	if sendIdCursor == nil {
		return
	}
	var unReads []models.Trainer
	err = sendIdCursor.All(context.TODO(), &unReads)
	if err != nil {
		log.Println("sendIdCursor err", err)
	}
	if len(unReads) > 0 {
		timeFilter := bson.M{
			"startTime": bson.M{
				"$gte": unReads[0].StartTime,
			},
		}
		sendIdTimeCursor, _ := sendIdCollection.Find(context.TODO(), timeFilter)
		idTimeCursor, _ := idCollection.Find(context.TODO(), timeFilter)
		err = sendIdTimeCursor.All(context.TODO(), &resultsYou)
		err = idTimeCursor.All(context.TODO(), &resultsMe)
		results, err = AppendAndSort(resultsMe, resultsYou)
	} else {
		results, err = FindMany(database, sendId, id, 9999999999, 10)
	}
	overTimeFilter := bson.D{
		{"$and", bson.A{
			bson.D{{"endTime", bson.M{"$lt": time.Now().Unix()}}},
			bson.D{{"read", bson.M{"$eq": 1}}},
		}},
	}
	_, _ = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
	_, _ = idCollection.DeleteMany(context.TODO(), overTimeFilter)
	// 将所有的维度设置为已读
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"$set": bson.M{"read": 1},
	})
	_, _ = sendIdCollection.UpdateMany(context.TODO(), filter, bson.M{
		"&set": bson.M{"ebdTime": time.Now().Unix() + int64(3*month)},
	})
	return
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := mongodb.MongoDBClient.Database(database).Collection(id)
	fmt.Println("id", id)
	comment := models.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}
