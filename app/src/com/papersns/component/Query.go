package component

import (
	. "com/papersns/mongo"
	"encoding/json"
	"labix.org/v2/mgo/bson"
)

type QuerySupport struct{}

func (qb QuerySupport) Find(collection string, query string) (result map[string]interface{}, found bool) {
	session, db := MongoDBFactory.GetConnection()
	defer session.Close()

	c := db.C(collection)

	queryMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(query), &queryMap)
	if err != nil {
		panic(err)
	}

	result = make(map[string]interface{})
	err = c.Find(queryMap).One(&result)
	if err != nil {
		return result, false
	}

	return result, true
}

func (qb QuerySupport) Index(collection string, query map[string]interface{}, pageNo int, pageSize int) (result map[string]interface{}) {
	session, db := MongoDBFactory.GetConnection()
	defer session.Close()

	c := db.C(collection)

	items := []interface{}{}
	err := c.Find(query).Limit(pageSize).Skip((pageNo - 1) * pageSize).All(&items)
	if err != nil {
		panic(err)
	}

	totalResults, err := c.Find(query).Count()
	if err != nil {
		panic(err)
	}

	mapItems := []interface{}{}
	for _, item := range items {
		record := item.(bson.M)
		mapItem := map[string]interface{}{}
		for k,v := range record {
			mapItem[k] = v
		}
		mapItems = append(mapItems, mapItem)
	}
	return map[string]interface{}{
		"totalResults": totalResults,
		"items":        mapItems,
	}
}
