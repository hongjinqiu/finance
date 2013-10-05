package model

import (
	"encoding/xml"
)

type DataSource struct {
	XMLName                 xml.Name     `xml:"datasource"`
	Id                      string       `xml:"id"`
	DisplayName             string       `xml:"displayName"`
	SystemId                string       `xml:"systemId"`
	CodeFieldName           string       `xml:"codeFieldName"`
	BusinessDateField       string       `xml:"businessDateField"`
	ModelType               string       `xml:"modelType"`
	InUsedDenyEdit          bool         `xml:"inUsedDenyEdit"`
	ActionNameSpace         string       `xml:"actionNameSpace"`
	ListUrl                 string       `xml:"listUrl"`
	BillTypeField           string       `xml:"billTypeField"`
	BillTypeParamDataSource string       `xml:"billTypeParamDataSource"`
	HasCheckField           string       `xml:"hasCheckField"`
	ListSortFields          string       `xml:"listSortFields"`
	MasterData              MasterData   `xml:"masterData"`
	DetailData              []DetailData `xml:"detailData"`
}

type MasterData struct {
	XMLName     xml.Name `xml:"masterData"`
	Id          string   `xml:"id"`
	DisplayName string   `xml:"displayName"`
	AllowCopy   bool     `xml:"allowCopy"`
	PrimaryKey  string   `xml:"primaryKey"`
	FixField    FixField `xml:"fixField"`
	BizField    BizField `xml:"bizField"`
}

type DetailData struct {
	XMLName       xml.Name `xml:"detailData"`
	Id            string   `xml:"id"`
	DisplayName   string   `xml:"displayName"`
	ParentId      string   `xml:"parentId"`
	AllowEmptyRow bool     `xml:"allowEmptyRow"`
	AllowCopy     bool     `xml:"allowCopy"`
	Readonly      bool     `xml:"readonly"`
	PrimaryKey    string   `xml:"primaryKey"`
	FixField      FixField `xml:"fixField"`
	BizField      BizField `xml:"bizField"`
}

type FixField struct {
	XMLName     xml.Name    `xml:"fixField"`
	PrimaryKey  PrimaryKey  `xml:"primaryKey"`
	CreateBy    CreateBy    `xml:"createBy"`
	CreateTime  CreateTime  `xml:"createTime"`
	CreateUnit  CreateUnit  `xml:"createUnit"`
	ModifyUnit  ModifyUnit  `xml:"modifyUnit"`
	ModifyTime  ModifyTime  `xml:"modifyTime"`
	BillStatus  BillStatus  `xml:"billStatus"`
	AttachCount AttachCount `xml:"attachCount"`
	Remark      Remark      `xml:"remark"`
}

type BizField struct {
	XMLName xml.Name `xml:"bizField"`
	Field   []Field  `xml:"field"`
}

type PrimaryKey struct {
	XMLName xml.Name `xml:"primaryKey"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type CreateBy struct {
	XMLName xml.Name `xml:"createBy"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type CreateTime struct {
	XMLName xml.Name `xml:"createTime"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type CreateUnit struct {
	XMLName xml.Name `xml:"createUnit"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type ModifyUnit struct {
	XMLName xml.Name `xml:"modifyUnit"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type ModifyTime struct {
	XMLName xml.Name `xml:"modifyTime"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type BillStatus struct {
	XMLName xml.Name `xml:"billStatus"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type AttachCount struct {
	XMLName xml.Name `xml:"attachCount"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}
type Remark struct {
	XMLName xml.Name `xml:"remark"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	Id      string   `xml:"id,attr"`
	Extends string   `xml:"extends,attr"`
	FieldGroup
}

type FieldGroup struct {
	FieldName         string `xml:"fieldName,attr"`
	DisplayName       string `xml:"displayName,attr"`
	FieldDataType     string `xml:"fieldDataType,attr"`
	FieldNumberType   string `xml:"fieldNumberType,attr"`
	FieldLength       string `xml:"fieldLength,attr"`
	DefaultValueExpr  string `xml:"defaultValueExpr,attr"`
	CheckInUsed       bool   `xml:"checkInUsed,attr"`
	FixHide           bool   `xml:"fixHide,attr"`
	FixReadOnly       bool   `xml:"fixReadOnly,attr"`
	AllowDuplicate    bool   `xml:"allowDuplicate,attr"`
	DenyEditInUsed    bool   `xml:"denyEditInUsed,attr"`
	AllowEmpty        bool   `xml:"allowEmpty,attr"`
	LimitOption       bool   `xml:"limitOption,attr"`
	LimitMax          string `xml:"limitMax,attr"`
	LimitMin          string `xml:"limitMin,attr"`
	ValidateExpr      string `xml:"validateExpr,attr"`
	ValidateMessage   string `xml:"validateMessage,attr"`
	Dictionary        string `xml:"dictionary,attr"`
	DictionaryWhere   string `xml:"dictionaryWhere,attr"`
	CalcValueExpr     string `xml:"calcValueExpr,attr"`
	Virtual           bool   `xml:"virtual,attr"`
	ZeroShowEmpty     bool   `xml:"zeroShowEmpty,attr"`
	LocalCurrencyency bool   `xml:"localCurrencyency,attr"`
	FieldInList       bool   `xml:"fieldInList,attr"`
	ListWhereField    bool   `xml:"listWhereField,attr"`
	FormatExpr        string `xml:"formatExpr,attr"`
	RelationDS        string `xml:"relationDS,attr"`
}
