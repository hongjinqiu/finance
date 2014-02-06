package controllers

//import "github.com/robfig/revel"
import (
	"com/papersns/global"
	. "com/papersns/model"
	. "com/papersns/model/handler"
	"com/papersns/mongo"
	. "com/papersns/mongo"
	. "com/papersns/script"
	"fmt"
	"labix.org/v2/mgo"
	"strconv"
	"strings"
)

type FinanceService struct{}

func (o FinanceService) SaveData(sessionId int, dataSource DataSource, bo *map[string]interface{}) *[]DiffDataRow {
	strId := ""
	if (*bo)["_id"] != nil {
		strId = fmt.Sprint((*bo)["_id"])
	}
	modelTemplateFactory := ModelTemplateFactory{}
	modelTemplateFactory.ConvertDataType(dataSource, bo)
	// 主数据集和分录数据校验
	message := o.validateBO((dataSource), (*bo))
	if message != "" {
		panic(message)
	}
	_, db := global.GetConnection(sessionId)
	if strId == "" || strId == "0" {
		// 主数据集和分录id赋值,
		modelIterator := ModelIterator{}
		var result interface{} = ""
		modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			o.setDataId(db, dataSource, &fieldGroup, bo, data)
		})
		// 被用过帐
		usedCheck := UsedCheck{}
		result = ""
		diffDataRowLi := []DiffDataRow{}
		modelIterator.IterateDataBo(dataSource, bo, &result, func(fieldGroupLi []FieldGroup, data *map[string]interface{}, rowIndex int, result *interface{}) {
			diffDataRowLi = append(diffDataRowLi, DiffDataRow{
				FieldGroupLi: fieldGroupLi,
				DestBo:       bo,
				DestData:     data,
				SrcData:      nil,
				SrcBo:        nil,
			})
		})
		for i, _ := range diffDataRowLi {
			fieldGroupLi := diffDataRowLi[i].FieldGroupLi
			data := diffDataRowLi[i].DestData
			usedCheck.Insert(sessionId, fieldGroupLi, bo, data)
		}
		txnManager := TxnManager{db}
		txnId := global.GetTxnId(sessionId)
		txnManager.Insert(txnId, dataSource.Id, *bo)

		return &diffDataRowLi
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		panic(err)
	}
	modelIterator := ModelIterator{}
	srcBo := map[string]interface{}{}
	var result interface{} = ""
	err = db.C(dataSource.Id).Find(map[string]interface{}{"id": id}).One(&srcBo)
	if err != nil {
		panic(err)
	}

	modelTemplateFactory.ConvertDataType(dataSource, &srcBo)
	diffDataRowLi := []DiffDataRow{}
	modelIterator.IterateDiffBo(dataSource, bo, srcBo, &result, func(fieldGroupLi []FieldGroup, destData *map[string]interface{}, srcData map[string]interface{}, result *interface{}) {
		// 分录+id
		if destData != nil {
			dataStrId := fmt.Sprint((*destData)["id"])
			if dataStrId == "" || dataStrId == "0" {
				for i, _ := range fieldGroupLi {
					o.setDataId(db, dataSource, &fieldGroupLi[i], bo, destData)
				}
			}
		}
		diffDataRowLi = append(diffDataRowLi, DiffDataRow{
			FieldGroupLi: fieldGroupLi,
			DestBo:       bo,
			DestData:     destData,
			SrcData:      srcData,
			SrcBo:        srcBo,
		})
	})

	// 被用差异行处理
	usedCheck := UsedCheck{}
	for i, _ := range diffDataRowLi {
		fieldGroupLi := diffDataRowLi[i].FieldGroupLi
		destData := diffDataRowLi[i].DestData
		srcData := diffDataRowLi[i].SrcData
		usedCheck.Update(sessionId, fieldGroupLi, bo, destData, srcData)
	}
	txnManager := TxnManager{db}
	txnId := global.GetTxnId(sessionId)
	//	txnManager.Update(txnId int, collection string, doc map[string]interface{}) (map[string]interface{}, bool) {
	if _, updateResult := txnManager.Update(txnId, dataSource.Id, *bo); !updateResult {
		panic("更新失败")
	}
	return &diffDataRowLi
}

func (o FinanceService) setDataId(db *mgo.Database, dataSource DataSource, fieldGroup *FieldGroup, bo *map[string]interface{}, data *map[string]interface{}) {
	if fieldGroup.Id == "id" {
		if fieldGroup.IsMasterField() {
			masterSeqName := GetMasterSequenceName((dataSource))
			masterSeqId := mongo.GetSequenceNo(db, masterSeqName)
			(*data)["_id"] = masterSeqId
			(*data)["id"] = masterSeqId
			(*bo)["_id"] = masterSeqId
			(*bo)["id"] = masterSeqId
		} else {
			detailData, found := fieldGroup.GetDetailData()
			if found {
				detailSeqName := GetDetailSequenceName((dataSource), detailData)
				detailSeqId := mongo.GetSequenceNo(db, detailSeqName)
				(*data)["_id"] = detailSeqId
				(*data)["id"] = detailSeqId
			}
		}
	}
}

func (o FinanceService) validateBO(dataSource DataSource, bo map[string]interface{}) string {
	messageLi := []string{}
	modelIterator := ModelIterator{}
	detailIndex := map[string]int{}
	for _, item := range dataSource.DetailDataLi {
		detailIndex[item.Id] = 0
	}
	var result interface{} = messageLi
	modelIterator.IterateAllFieldBo(dataSource, &bo, &result, func(fieldGroup FieldGroup, data *map[string]interface{}, rowIndex int, messageLi *interface{}) {
		stringLi := (*messageLi).([]string)
		fieldMessageLi := o.validateFieldGroup(fieldGroup, *data)
		if fieldGroup.IsMasterField() {
			for _, item := range fieldMessageLi {
				stringLi = append(stringLi, item)
			}
		} else {
			detailData, _ := fieldGroup.GetDetailData()
			detailIndex[detailData.Id]++
			for _, item := range fieldMessageLi {
				stringLi = append(stringLi, "分录:"+detailData.DisplayName+"序号为"+strconv.Itoa(detailIndex[detailData.Id])+"的数据,"+item)
			}
		}
	})
	return strings.Join(messageLi, "<br />")
}

func (o FinanceService) validateFieldGroup(fieldGroup FieldGroup, data map[string]interface{}) []string {
	messageLi := []string{}

	if fieldGroup.AllowEmpty != "true" {
		value := data[fieldGroup.Id]
		if value != nil {
			strValue := fmt.Sprint(value)
			if strValue == "" {
				messageLi = append(messageLi, fieldGroup.DisplayName+"不允许空值")
				return messageLi
			}
		} else {
			messageLi = append(messageLi, fieldGroup.DisplayName+"不允许空值")
			return messageLi
		}
	}
	fieldValue := fmt.Sprint(data[fieldGroup.Id])
	if fieldGroup.ValidateExpr != "" {
		// python and golang validate, TODO miss golang validate
		expressionParser := ExpressionParser{}
		if !expressionParser.Validate(fieldValue, fieldGroup.ValidateExpr) {
			messageLi = append(messageLi, fieldGroup.DisplayName+fieldGroup.ValidateMessage)
			return messageLi
		}
	}
	isDataTypeNumber := false
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "DECIMAL"
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "FLOAT"
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "INT"
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "LONGINT"
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "MONEY"
	isDataTypeNumber = isDataTypeNumber || fieldGroup.FieldDataType == "SMALLINT"
	isUnLimit := fieldGroup.LimitOption != "" && fieldGroup.LimitOption != "unLimit"
	if isDataTypeNumber && isUnLimit {
		fieldValueFloat, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			panic(err)
		}
		if fieldGroup.LimitOption == "limitMax" {
			maxValue, err := strconv.ParseFloat(fieldGroup.LimitMax, 64)
			if err != nil {
				panic(err)
			}
			if maxValue < fieldValueFloat {
				messageLi = append(messageLi, fieldGroup.DisplayName+"超出最大值"+fieldGroup.LimitMax)
			}
		} else if fieldGroup.LimitOption == "limitMin" {
			minValue, err := strconv.ParseFloat(fieldGroup.LimitMin, 64)
			if err != nil {
				panic(err)
			}
			if fieldValueFloat < minValue {
				messageLi = append(messageLi, fieldGroup.DisplayName+"小于最小值"+fieldGroup.LimitMin)
			}
		} else if fieldGroup.LimitOption == "limitRange" {
			minValue, err := strconv.ParseFloat(fieldGroup.LimitMin, 64)
			if err != nil {
				panic(err)
			}
			maxValue, err := strconv.ParseFloat(fieldGroup.LimitMax, 64)
			if err != nil {
				panic(err)
			}
			if fieldValueFloat < minValue || maxValue < fieldValueFloat {
				messageLi = append(messageLi, fieldGroup.DisplayName+"超出范围("+fieldGroup.LimitMin+"~"+fieldGroup.LimitMax+")")
			}
		}
	} else {
		isDataTypeString := false
		isDataTypeString = isDataTypeString || fieldGroup.FieldDataType == "STRING"
		isDataTypeString = isDataTypeString || fieldGroup.FieldDataType == "REMARK"
		isFieldLengthLimit := fieldGroup.FieldLength != ""
		if isDataTypeString && isFieldLengthLimit {
			limit, err := strconv.Atoi(fmt.Sprint(fieldGroup.FieldLength))
			if err != nil {
				panic(err)
			}
			if len(fieldValue) > limit {
				messageLi = append(messageLi, fieldGroup.DisplayName+"长度超出最大值"+fieldGroup.FieldLength)
			}
		}
	}

	return messageLi
}
