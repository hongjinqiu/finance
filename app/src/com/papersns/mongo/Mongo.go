package mongo

import "github.com/robfig/revel"
import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
	"strings"
)

func GetInstance() ConnectionFactory {
	return ConnectionFactory{}
}

type ConnectionFactory struct{}

func (c ConnectionFactory) GetConnection() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial(revel.Config.StringDefault("MONGODB_ADDRESS", "localhost:27017"))
	if err != nil {
		panic(err)
	}
	//	session.SetMode(mgo.Monotonic, true)
	db := session.DB(revel.Config.StringDefault("MONGODB_DATABASE_NAME", "finance"))
	return session, db
}

func (c ConnectionFactory) GetSession() *mgo.Session {
	session, err := mgo.Dial(revel.Config.StringDefault("MONGODB_ADDRESS", "localhost:27017"))
	if err != nil {
		panic(err)
	}
	return session
}

func (c ConnectionFactory) GetDatabase(session *mgo.Session) *mgo.Database {
	//	session.SetMode(mgo.Monotonic, true)
	db := session.DB(revel.Config.StringDefault("MONGODB_DATABASE_NAME", "finance"))
	return db
}

func GetCollectionSequenceName(collection string) string {
	byte0 := collection[0]
	return strings.ToLower(string(byte0)) + collection[1:] + "Id"
}

func GetSequenceNo(db *mgo.Database, sequenceName string) int {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"c": 1}},
		ReturnNew: true,
	}
	doc := map[string]interface{}{}
	_, err := db.C("counters").Find(bson.M{"_id": sequenceName}).Apply(change, &doc)
	if err != nil {
		println("^^^^sequenceName is:" + sequenceName)
		panic(err)
	}
	idText := fmt.Sprint(doc["c"])
	id, err := strconv.Atoi(idText)
	if err != nil {
		panic(err)
	}
	return id
}
