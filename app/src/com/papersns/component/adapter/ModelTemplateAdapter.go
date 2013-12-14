package adapter

import (
	. "com/papersns/component"
	. "com/papersns/model"
	"strings"
)

type ModelTemplateAdapter struct{}

func (o ModelTemplateAdapter) ApplyModelAdapter(listTemplate *ListTemplate) {
	if listTemplate.DataSourceModelId != "" {
		modelTemplateFactory := ModelTemplateFactory{}
		dataSource := modelTemplateFactory.GetDataSource(listTemplate.DataSourceModelId)
		o.applyDataProvider(dataSource, listTemplate)
		o.applyColumnModel(dataSource, listTemplate)
	}
}

func (o ModelTemplateAdapter) applyDataProvider(dataSource DataSource, listTemplate *ListTemplate) {
	if listTemplate.DataProvider.Collection == "" {
		listTemplate.DataProvider.Collection = dataSource.Id
	}
}

/*
	<element name="fieldDataType" minOccurs="0" maxOccurs="1">
	STRING,SMALLINT,INT,LONGINT,BOOLEAN,FLOAT,MONEY,DECIMAL,REMARK,BLOB
	<!-- 字段数值类型 -->
	<element name="fieldNumberType" minOccurs="0" maxOccurs="1">
	UNDEFINE,MONEY,PRICE,EXCHANGERATE,PERCENT,QUANTITY,UNITCOST,YEAR,YEARMONTH,DATE,TIME
*/

func (o ModelTemplateAdapter) applyColumnModel(dataSource DataSource, listTemplate *ListTemplate) {
	modelIterator := ModelIterator{}
	var result interface{} = ""
	// string-column,number-column,date-column,boolean-column,dictionary-column,select-column
	// 还有嵌套的问题,先整完,再考虑递归,
	for i, _ := range listTemplate.ColumnModel.ColumnLi {
		column := listTemplate.ColumnModel.ColumnLi[i]
		if column.XMLName.Local == "auto-column" {
			modelIterator.IterateAllField(&dataSource, result, func(fieldGroup *FieldGroup, result *interface{}){
				if column.Name == fieldGroup.Id {
					if column.Text == "" {
						column.Text = fieldGroup.DisplayName
					}
					if column.Hideable == "" {
						column.Hideable = fieldGroup.FixHide
					}
					isIntField := false
					intArray := []string{"SMALLINT", "INT", "LONGINT"}
					for item := range intArray {
						if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower(item) {
							isIntField = true
							break
						}
					}
					isFloatField := false
					floatArray := []string{"FLOAT", "MONEY", "DECIMAL"}
					for item := range floatArray {
						if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower(item) {
							isFloatField = true
							break
						}
					}
					isDateType := false
					dateArray := []string{"YEAR","YEARMONTH","DATE","TIME","DATETIME"}
					for item := range dateArray {
						if strings.ToLower(fieldGroup.FieldDataType) == strings.ToLower(item) {
							isDateType = true
							break
						}
					}
					if fieldGroup.IsRelationField() {
						column.XMLName.Local = "select-column"
					} else if strings.ToLower(fieldGroup.FieldDataType) == "string" {
						column.XMLName.Local = "string-column"
					} else if (isIntField || isFloatField) && !isDateType && fieldGroup.Dictionary == "" {
						column.XMLName.Local = "number-column"
//						fieldNumberType
						//number-column,fields:name,text,currencyField,isMoney,isUnitPrice,isCost,isPercent,
						if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("MONEY") {
							column.IsMoney = "true"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("PRICE") {
							column.IsUnitPrice = "true"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("UNITCOST") {
							column.IsCost = "true"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("PERCENT") {
							column.IsPercent = "true"
						}
					} else if (isIntField || isFloatField) && isDateType && fieldGroup.Dictionary == "" {
/*
需要从数据库中查找
<attribute name="displayPattern">
	<annotation>
		<documentation>日期格式化</documentation>
	</annotation>
	<simpleType>
		<restriction base="string">
			<enumeration value="yyyy"></enumeration>
			<enumeration value="yyyy-MM"></enumeration>
			<enumeration value="HH:mm:ss"></enumeration>
			<enumeration value="HH:mm"></enumeration>
			<enumeration value="yyyy-MM-dd"></enumeration>
			<enumeration value="yyyy-MM-dd HH:mm"></enumeration>
			<enumeration value="yyyy-MM-dd HH:mm:ss"></enumeration>
		</restriction>
	</simpleType>
</attribute>
<attribute name="dbPattern">
	<annotation>
		<documentation>数据库中存储的日期格式</documentation>
	</annotation>
	<simpleType>
		<restriction base="string">
			<enumeration value="yyyy"></enumeration>
			<enumeration value="yyyyMM"></enumeration>
			<enumeration value="HHmmss"></enumeration>
			<enumeration value="HHmm"></enumeration>
			<enumeration value="yyyyMMdd"></enumeration>
			<enumeration value="yyyyMMddHHmm"></enumeration>
			<enumeration value="yyyyMMddHHmmss"></enumeration>
		</restriction>
	</simpleType>
</attribute>
*/
						column.XMLName.Local = "date-column"
						if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATE") {
							column.DisplayPattern = "yyyy"//需要从业务中查找
							column.DbPattern = "yyyy"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("DATETIME") {
							column.DisplayPattern = "yyyy-MM-dd HH:mm:ss"//需要从业务中查找
							column.DbPattern = "yyyyMMddHHmmss"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEAR") {
							column.IsMoney = "true"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("YEARMONTH") {
							column.IsMoney = "true"
						} else if strings.ToLower(fieldGroup.FieldNumberType) == strings.ToLower("TIME") {
							column.IsMoney = "true"
						}
					}
					// fixHide
					// string-column,conditions is:fieldDataType=string,Dictionary=""
						// string-column,fields:column-attribute,editor,column-model,column-attributes
						// extend field:name,text,
					// number-column,conditions is:fieldDataType=int,float,fieldNumberType!=year,yearmonth,date,time,datetime,&& dictionary!=""不管币别,币别在列表配置出来,
						// number-column,fields:name,text,currencyField,isMoney,isUnitPrice,isCost,isPercent,
						// extend field:name,text
						// according field:fieldDataType,
					// date-column,fieldDataType=int,float,fieldNumberType=year,yearmonth,date,time,datetime,
						// date-column,fields:name,text,(displayPattern,dbPattern),需要自己从业务中读取,
					// boolean-column,fieldDataType=boolean,没用了,用字典列,干掉,
						// boolean-column,fields:name,text,
					// dictionary-column,fieldDataType=string,Dictionary!="",
						// dictionary-column,fields:name,text,dictionary
					// select-column,isRelationDS,只生成text,其它要自己生成
						// select-column,fields:displayField,valueField,selectorName,selectionMode,queryFunc,不管,自己写,
					// html,直接取text,但是listTemplate不渲染,
				}
			})
		} else {
			// 查找auto的field,还有extend,
		}
	}
}

