package component

import (
	"encoding/json"
//	"labix.org/v2/mgo/bson"
	. "com/papersns/mongo"
)

type QuerySupport struct{}

func (qb QuerySupport) BuildQuery(listTemplate ListTemplate) {

}

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
