package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
	"fmt"
)

func GetMasterSequenceName(dataSource DataSource) string {
	byte0 := dataSource.Id[0]
	return string(byte0) + dataSource.Id[1:] + "Id"
}

func GetDetailSequenceName(dataSource DataSource, detailData DetailData) string {
	byte0 := dataSource.Id[0]
	return string(byte0) + dataSource.Id[1:] + detailData.Id + "Id"
}

func GetSequenceNo(db *mgo.Database, sequenceName string) int {
	change := mgo.Change{
        Update: bson.M{"$inc": bson.M{"c": 1}},
        ReturnNew: true,
	}
	doc := map[string]interface{}{}
	_, err := db.C("counters").Find(bson.M{"_id": sequenceName}).Apply(change, &doc)
	if err != nil {
		panic(err)
	}
	idText := fmt.Sprint(doc["c"])
	id, err := strconv.Atoi(idText)
	if err != nil {
		panic(err)
	}
	return id
}
