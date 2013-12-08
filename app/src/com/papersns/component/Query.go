package component

import (
	"com/papersns/mongo"
	"encoding/json"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

type QuerySupport struct{}

func (qb QuerySupport) FindByMap(collection string, query map[string]interface{}) (result map[string]interface{}, found bool) {
	mongoDBFactory := mongo.GetInstance()
	session := mongoDBFactory.GetSession()
	defer session.Close()

	return qb.FindByMapWithSession(session, collection, query)
}

func (qb QuerySupport) FindByMapWithSession(session *mgo.Session, collection string, query map[string]interface{}) (result map[string]interface{}, found bool) {
	mongoDBFactory := mongo.GetInstance()
	db := mongoDBFactory.GetDatabase(session)
	c := db.C(collection)

	result = make(map[string]interface{})
	err := c.Find(query).One(&result)
	if err != nil {
		return result, false
	}

	return result, true
}

func (qb QuerySupport) Find(collection string, query string) (result map[string]interface{}, found bool) {
	queryMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(query), &queryMap)
	if err != nil {
		panic(err)
	}

	return qb.FindByMap(collection, queryMap)
}

func (qb QuerySupport) Index(collection string, query map[string]interface{}, pageNo int, pageSize int, orderBy string) (result map[string]interface{}) {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()

	c := db.C(collection)

	items := []interface{}{}
	var err error
	if orderBy != "" {
		fieldLi := strings.Split(orderBy, ",")
		err = c.Find(query).Sort(fieldLi...).Limit(pageSize).Skip((pageNo - 1) * pageSize).All(&items)
	} else {
		err = c.Find(query).Limit(pageSize).Skip((pageNo - 1) * pageSize).All(&items)
	}
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

func (qb QuerySupport) MapReduceAll(collection string, query map[string]interface{}, mapReduce mgo.MapReduce) (result []map[string]interface{}) {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	
	result = []map[string]interface{}{}
	_, err := db.C(collection).Find(query).MapReduce(&mapReduce, &result)
	if err != nil {
		panic(err)
	}
	
	return result 
}

func (qb QuerySupport) MapReduce(collection string, query map[string]interface{}, mapReduce mgo.MapReduce, pageNo int, pageSize int) (result []map[string]interface{}) {
	mongoDBFactory := mongo.GetInstance()
	session, db := mongoDBFactory.GetConnection()
	defer session.Close()
	
	result = []map[string]interface{}{}
	_, err := db.C(collection).Find(query).Limit(pageSize).Skip((pageNo - 1) * pageSize).MapReduce(&mapReduce, &result)
	if err != nil {
		panic(err)
	}
	
	return result 
}
