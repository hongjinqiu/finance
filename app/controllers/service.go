package controllers

//import "github.com/robfig/revel"
import (
	"com/papersns/mongo"
	. "com/papersns/component"
	. "com/papersns/model"
	"com/papersns/model/handler"
	. "com/papersns/model/handler"
	"fmt"
	"strings"
	"strconv"
)

type FinanceService struct{}

func (o FinanceService) SaveData(dataSource *DataSource, bo *map[string]interface{}) {
	id := fmt.Sprint((*bo)["_id"])
	// 主数据集和分录数据校验
	message := o.validateBO((*dataSource), (*bo))
	if message != "" {
		panic(message)
	}
	if id == "" {
		// 主数据集和分录id赋值,
		mongoDBFactory := mongo.GetInstance()
		session, db := mongoDBFactory.GetConnection()
		defer session.Close()

		modelIterator := ModelIterator{}
		var result interface{} = ""
		modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup *FieldGroup, data *map[string]interface{}, result *interface{}) {
			if fieldGroup.Id == "id" {
				if fieldGroup.IsMasterField() {
					masterSeqName := GetMasterSequenceName((*dataSource))
					masterSeqId := GetSequenceNo(db, masterSeqName)
					(*data)["_id"] = masterSeqId
					(*data)["id"] = masterSeqId
					(*bo)["_id"] = masterSeqId
					(*bo)["id"] = masterSeqId
				} else {
					detailData, found := fieldGroup.GetDetailData()
					if found {
						detailSeqName := GetDetailSequenceName((*dataSource), detailData)
						detailSeqId := GetSequenceNo(db, detailSeqName)
						(*data)["_id"] = detailSeqId
						(*data)["id"] = detailSeqId
					}
				}
			}
		})
		// 被用过帐
		usedCheck := UsedCheck{}
		result = ""
		modelIterator.IterateAllFieldBo(dataSource, bo, &result, func(fieldGroup *FieldGroup, data *map[string]interface{}, result *interface{}) {
			if fieldGroup.IsRelationField() {
				usedCheck.Execute(db, fieldGroup, data, handler.ADD)
			}
		})
	} else {
		// 分录差异行处理,用简单地查找出删除的数据,然后更新或删除的操作,赋值一个DiffDataType到map里面,
		// 分录+id
		// 根据diffDataType被用过帐
	}
}

func (o FinanceService) validateBO(dataSource DataSource, bo map[string]interface{}) string {
	messageLi := []string{}
	
	data := bo["A"].(map[string]interface{})
	fieldMessageLi := o.validateFixField(dataSource.MasterData.FixField, data)
	for _, fieldMessage := range fieldMessageLi {
		messageLi = append(messageLi, fieldMessage)
	}
	fieldMessageLi = o.validateBizField(dataSource.MasterData.BizField, data)
	for _, fieldMessage := range fieldMessageLi {
		messageLi = append(messageLi, fieldMessage)
	}
	
	for i, detailData := range dataSource.DetailDataLi {
		dataLi := bo[detailData.Id].([]interface{})
		for _, dataItem := range dataLi {
			data := dataItem.(map[string]interface{})
			fieldMessageLi = o.validateFixField(detailData.FixField, data)
			for _, fieldMessage := range fieldMessageLi {
				messageLi = append(messageLi, "分录:" + detailData.DisplayName + "序号为" + strconv.Itoa(i) + "的数据," + fieldMessage)
			}
			fieldMessageLi = o.validateBizField(detailData.BizField, data)
			for _, fieldMessage := range fieldMessageLi {
				messageLi = append(messageLi, "分录:" + detailData.DisplayName + "序号为" + strconv.Itoa(i) + "的数据," + fieldMessage)
			}
		}
	}
	
	return strings.Join(messageLi, "<br />")
}

func (o FinanceService) validateFixField(fixField FixField, data map[string]interface{}) []string {
	messageLi := []string{}
	
	fixFieldLi := []FieldGroup{}
	fixFieldLi = append(fixFieldLi, fixField.CreateBy.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.CreateTime.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.ModifyBy.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.ModifyTime.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.BillStatus.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.AttachCount.FieldGroup)
	fixFieldLi = append(fixFieldLi, fixField.Remark.FieldGroup)
	
	for _, field := range fixFieldLi {
		fieldMessageLi := o.validateFieldGroup(field, data)
		for _, fieldMessage := range fieldMessageLi {
			messageLi = append(messageLi, fieldMessage)
		}
	}
	return messageLi
}

func (o FinanceService) validateBizField(bizField BizField, data map[string]interface{}) []string {
	messageLi := []string{}
	
	for _, field := range bizField.FieldLi {
		fieldMessageLi := o.validateFieldGroup(field.FieldGroup, data)
		for _, fieldMessage := range fieldMessageLi {
			messageLi = append(messageLi, fieldMessage)
		}
	}
	return messageLi
}

func (o FinanceService) validateFieldGroup(fieldGroup FieldGroup, data map[string]interface{}) []string {
	messageLi := []string{}

	if fieldGroup.AllowEmpty != "true" {
		value := data[fieldGroup.Id]
		if value != nil {
			strValue := fmt.Sprint(value)
			if strValue == "" {
				messageLi = append(messageLi, fieldGroup.DisplayName + "不允许空值")
				return messageLi
			}
		} else {
			messageLi = append(messageLi, fieldGroup.DisplayName + "不允许空值")
			return messageLi
		}
	}
	fieldValue := fmt.Sprint(data[fieldGroup.Id])
	if fieldGroup.ValidateExpr != "" {
		// python and golang validate, TODO miss golang validate
		expressionParser := ExpressionParser{}
		if !expressionParser.Validate(fieldValue, fieldGroup.ValidateExpr) {
			messageLi = append(messageLi, fieldGroup.DisplayName + fieldGroup.ValidateMessage)
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
				messageLi = append(messageLi, fieldGroup.DisplayName + "超出最大值" + fieldGroup.LimitMax)
			}
		} else if fieldGroup.LimitOption == "limitMin" {
			minValue, err := strconv.ParseFloat(fieldGroup.LimitMin, 64)
			if err != nil {
				panic(err)
			}
			if fieldValueFloat < minValue {
				messageLi = append(messageLi, fieldGroup.DisplayName + "小于最小值" + fieldGroup.LimitMin)
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
				messageLi = append(messageLi, fieldGroup.DisplayName + "超出范围(" + fieldGroup.LimitMin + "~" + fieldGroup.LimitMax + ")")
			}
		}
	}

	return messageLi
}

func (o FinanceService) convertDataType(fieldGroup FieldGroup, data *map[string]interface{}) string {
	return ""
}
