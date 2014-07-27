package controllers

import "github.com/robfig/revel"
import (
	. "com/papersns/common"
	. "com/papersns/model"
	"com/papersns/mongo"
	"fmt"
	"math/rand"
	"time"
	"io"
	"crypto/sha1"
//	. "com/papersns/taobao"
)

type Test struct {
	*revel.Controller
}

func (c Test) Index() revel.Result {
	time.Sleep(time.Millisecond * 500)
	return c.RenderText("index0")
}

func (c Test) Index1() revel.Result {
	for i := 0; i <= 21474836; i++ {
		if i == rand.Int() {
			println(i)
		}
	}
	return c.RenderText("index1")
}

func (c Test) Index2() revel.Result {
	time.Sleep(time.Second)
	return c.RenderText("index2")
}

func (c Test) StringTest() revel.Result {
	return c.RenderText(fmt.Sprint(nil))
}

func (c Test) Sha1Test() revel.Result {
	hash := sha1.New()
	_, err := io.WriteString(hash, "qwertyuiop")
	if err != nil {
		panic(err)
	}
	text := fmt.Sprintf("%x", hash.Sum(nil))
	return c.RenderText(text)
}

func (c Test) AddRemoveTestData() revel.Result {
	userIdLi := []int{10, 20}
	modelTemplateFactory := ModelTemplateFactory{}
	dataSourceInfoLi := modelTemplateFactory.GetDataSourceInfoLi()
	//	 GetCollectionName
	// 初始化数据只有 billType,BillTypeParameter,CurrencyType,
	connectionFactory := mongo.GetInstance()
	_, db := connectionFactory.GetConnection()
	modelIterator := ModelIterator{}
	for _, item := range dataSourceInfoLi {
		collectionName := modelTemplateFactory.GetCollectionName(item.DataSource)
		_, err := db.C(collectionName).RemoveAll(map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		if collectionName == "BillType" || collectionName == "BillTypeParameter" || collectionName == "CurrencyType" {

		} else {
			for _, userId := range userIdLi {
				for i := 0; i < 25; i += 1 {
					masterData := map[string]interface{}{}
					bo := map[string]interface{}{
						"A": masterData,
					}
					var result interface{} = ""
					modelIterator.IterateAllField(&item.DataSource, &result, func(fieldGroup *FieldGroup, result *interface{}) {
						if fieldGroup.IsMasterField() {
							c.initData(userId, *fieldGroup, &masterData)
						} else {
							dataSetId := fieldGroup.GetDataSetId()
							if bo[dataSetId] == nil {
								bo[dataSetId] = []interface{}{}
							}
							dataSetDataLi := bo[dataSetId].([]interface{})
							bo[dataSetId] = dataSetDataLi

							if len(dataSetDataLi) == 0 {
								dataSetDataLi = append(dataSetDataLi, map[string]interface{}{})
							}
							dataSetData0 := dataSetDataLi[0].(map[string]interface{})
							dataSetDataLi[0] = dataSetData0

							c.initData(userId, *fieldGroup, &dataSetData0)
						}
					})
					collectionSequenceName := mongo.GetCollectionSequenceName(collectionName)
					id := mongo.GetSequenceNo(db, collectionSequenceName)
					bo["_id"] = id
					bo["id"] = id
					masterData["id"] = id
					masterData["createBy"] = userId
					masterData["createUnit"] = userId
					masterData["createTime"] = DateUtil{}.GetCurrentYyyyMMddHHmmss()
					masterData["modifyBy"] = userId
					masterData["modifyUnit"] = userId
					masterData["modifyTime"] = DateUtil{}.GetCurrentYyyyMMddHHmmss()
					modelIterator.IterateDataBo(item.DataSource, &bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
						if !fieldGroupLi[0].IsMasterField() {
							id := mongo.GetSequenceNo(db, collectionSequenceName)
							(*data)["id"] = id
							(*data)["createBy"] = userId
							(*data)["createUnit"] = userId
							(*data)["createTime"] = DateUtil{}.GetCurrentYyyyMMddHHmmss()
							(*data)["modifyBy"] = userId
							(*data)["modifyUnit"] = userId
							(*data)["modifyTime"] = DateUtil{}.GetCurrentYyyyMMddHHmmss()
						}
					})
					modelTemplateFactory.ConvertDataType(item.DataSource, &bo)
					db.C(collectionName).Insert(bo)
				}
			}
		}
	}
	return c.RenderText("success")
}

func (c Test) initData(userId int, fieldGroup FieldGroup, masterData *map[string]interface{}) {
	dateUtil := DateUtil{}
	if fieldGroup.FieldDataType == "STRING" {
		(*masterData)[fieldGroup.Id] = "user" + fmt.Sprint(userId) + "_data_" + fmt.Sprint(rand.Int())
	} else if fieldGroup.FieldDataType == "FLOAT" {
		(*masterData)[fieldGroup.Id] = fmt.Sprint(rand.Float64())[:10]
	} else if fieldGroup.FieldNumberType == "YEAR" {
		(*masterData)[fieldGroup.Id] = dateUtil.GetCurrentYyyyMMdd() / (100 * 100)
	} else if fieldGroup.FieldNumberType == "YEARMONTH" {
		(*masterData)[fieldGroup.Id] = dateUtil.GetCurrentYyyyMMdd() / 100
	} else if fieldGroup.FieldNumberType == "DATE" {
		(*masterData)[fieldGroup.Id] = dateUtil.GetCurrentYyyyMMdd()
	} else if fieldGroup.FieldNumberType == "TIME" {
		(*masterData)[fieldGroup.Id] = 180605
	} else if fieldGroup.FieldNumberType == "DATETIME" {
		(*masterData)[fieldGroup.Id] = dateUtil.GetCurrentYyyyMMddHHmmss()
	} else { // int
		if fieldGroup.Id == "billTypeId" {
			(*masterData)[fieldGroup.Id] = 1
		} else if fieldGroup.Id == "billTypeParameterId" {
			(*masterData)[fieldGroup.Id] = 1
		} else if fieldGroup.Id == "currencyTypeId" {
			(*masterData)[fieldGroup.Id] = 1
		} else {
			(*masterData)[fieldGroup.Id] = rand.Int()
		}
	}
}

//func (c Test) GetTokenTest() revel.Result {
//	taobaoInterface := TaobaoInterface{}
//	text := taobaoInterface.GetTokenTest(map[string]string{})
//	return c.RenderText(text)
//}
