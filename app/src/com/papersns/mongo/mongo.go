package mongo

import "github.com/robfig/revel"
import (
	"labix.org/v2/mgo"
)

var MongoDBFactory connectionFactory

type connectionFactory struct {}

func (c connectionFactory) GetConnection() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial(revel.Config.StringDefault("MONGODB_ADDRESS", "localhost:27017"))
	if err != nil {
		panic(err)
	}
	//	session.SetMode(mgo.Monotonic, true)
	db := session.DB(revel.Config.StringDefault("MONGODB_DATABASE_NAME", "finance"))
	return session, db
}

