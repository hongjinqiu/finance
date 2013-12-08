package mongo

import "github.com/robfig/revel"
import (
	"labix.org/v2/mgo"
)

func GetInstance() ConnectionFactory {
	return ConnectionFactory{}
}

type ConnectionFactory struct {}

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
