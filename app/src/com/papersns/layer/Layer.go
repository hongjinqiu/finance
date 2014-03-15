package layer

import (
	"com/papersns/dictionary"
	"labix.org/v2/mgo"
	"com/papersns/mongo"
)

func GetInstance() LayerManager {
	return LayerManager{}
}

type LayerManager struct {
}

func (o LayerManager) GetLayer(code string) map[string]interface{} {
	connectionFactory := mongo.GetInstance()
	session, db := connectionFactory.GetConnection()
	defer session.Close()

	return o.GetLayerBySession(db, code)
}

func (o LayerManager) GetLayerBySession(db *mgo.Database, code string) map[string]interface{} {
	dictionaryManager := dictionary.GetInstance()
	result := dictionaryManager.GetDictionaryBySession(db, code)
	if result == nil {
		programDictionaryManager := dictionary.GetProgramDictionaryInstance()
		result = programDictionaryManager.GetProgramDictionaryBySession(db, code)
	}
	return result
}

