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
	InUsedDenyEdit          string       `xml:"inUsedDenyEdit"`
	ActionNameSpace         string       `xml:"actionNameSpace"`
	ListUrl                 string       `xml:"listUrl"`
	BillTypeField           string       `xml:"billTypeField"`
	BillTypeParamDataSource string       `xml:"billTypeParamDataSource"`
	HasCheckField           string       `xml:"hasCheckField"`
	ListSortFields          string       `xml:"listSortFields"`
	MasterData              MasterData   `xml:"masterData"`
	DetailDataLi            []DetailData `xml:"detailData"`
}

type MasterData struct {
	XMLName     xml.Name    `xml:"masterData"`
	Id          string      `xml:"id"`
	DisplayName string      `xml:"displayName"`
	AllowCopy   string      `xml:"allowCopy"`
	PrimaryKey  string      `xml:"primaryKey"`
	FixField    FixField    `xml:"fixField"`
	BizField    BizField    `xml:"bizField"`
	Parent      interface{} `xml:"-"`
}

type DetailData struct {
	XMLName       xml.Name    `xml:"detailData"`
	Id            string      `xml:"id"`
	DisplayName   string      `xml:"displayName"`
	ParentId      string      `xml:"parentId"`
	AllowEmptyRow string      `xml:"allowEmptyRow"`
	AllowCopy     string      `xml:"allowCopy"`
	Readonly      string      `xml:"readonly"`
	PrimaryKey    string      `xml:"primaryKey"`
	FixField      FixField    `xml:"fixField"`
	BizField      BizField    `xml:"bizField"`
	Parent        interface{} `xml:"-"`
}

type FixField struct {
	XMLName    xml.Name   `xml:"fixField"`
	PrimaryKey PrimaryKey `xml:"primaryKey"`
	CreateBy   CreateBy   `xml:"createBy"`
	CreateTime CreateTime `xml:"createTime"`
	//	CreateUnit  CreateUnit  `xml:"createUnit"`
	//	ModifyUnit  ModifyUnit  `xml:"modifyUnit"`
	ModifyBy    ModifyBy    `xml:"modifyBy"`
	ModifyTime  ModifyTime  `xml:"modifyTime"`
	BillStatus  BillStatus  `xml:"billStatus"`
	AttachCount AttachCount `xml:"attachCount"`
	Remark      Remark      `xml:"remark"`
	Parent      interface{} `xml:"-"`
}

type BizField struct {
	XMLName xml.Name    `xml:"bizField"`
	FieldLi []Field     `xml:"field"`
	Parent  interface{} `xml:"-"`
}

type PrimaryKey struct {
	XMLName xml.Name `xml:"primaryKey"`
	FieldGroup
}
type CreateBy struct {
	XMLName xml.Name `xml:"createBy"`
	FieldGroup
}
type CreateTime struct {
	XMLName xml.Name `xml:"createTime"`
	FieldGroup
}
type CreateUnit struct {
	XMLName xml.Name `xml:"createUnit"`
	FieldGroup
}
type ModifyBy struct {
	XMLName xml.Name `xml:"modifyBy"`
	FieldGroup
}
type ModifyUnit struct {
	XMLName xml.Name `xml:"modifyUnit"`
	FieldGroup
}
type ModifyTime struct {
	XMLName xml.Name `xml:"modifyTime"`
	FieldGroup
}
type BillStatus struct {
	XMLName xml.Name `xml:"billStatus"`
	FieldGroup
}
type AttachCount struct {
	XMLName xml.Name `xml:"attachCount"`
	FieldGroup
}
type Remark struct {
	XMLName xml.Name `xml:"remark"`
	FieldGroup
}

type Fields struct {
	XMLName xml.Name `xml:"fields"`
	FieldLi []Field  `xml:"field"`
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	FieldGroup
}

type FieldGroup struct {
	Id      string `xml:"id,attr"`
	Extends string `xml:"extends,attr"`
	//	FieldName         string     `xml:"fieldName"`
	DisplayName       string           `xml:"displayName"`
	FieldDataType     string           `xml:"fieldDataType"`
	FieldNumberType   string           `xml:"fieldNumberType"`
	FieldLength       string           `xml:"fieldLength"`
	DefaultValueExpr  DefaultValueExpr `xml:"defaultValueExpr"`
	CheckInUsed       string           `xml:"checkInUsed"`
	FixHide           string           `xml:"fixHide"`
	FixReadOnly       string           `xml:"fixReadOnly"`
	AllowDuplicate    string           `xml:"allowDuplicate"`
	DenyEditInUsed    string           `xml:"denyEditInUsed"`
	AllowEmpty        string           `xml:"allowEmpty"`
	LimitOption       string           `xml:"limitOption"`
	LimitMax          string           `xml:"limitMax"`
	LimitMin          string           `xml:"limitMin"`
	ValidateExpr      string           `xml:"validateExpr"`
	ValidateMessage   string           `xml:"validateMessage"`
	Dictionary        string           `xml:"dictionary"`
	DictionaryWhere   string           `xml:"dictionaryWhere"`
	CalcValueExpr     string           `xml:"calcValueExpr"`
	Virtual           string           `xml:"virtual"`
	ZeroShowEmpty     string           `xml:"zeroShowEmpty"`
	LocalCurrencyency string           `xml:"localCurrencyency"`
	FieldInList       string           `xml:"fieldInList"`
	ListWhereField    string           `xml:"listWhereField"`
	FormatExpr        string           `xml:"formatExpr"`
	RelationDS        RelationDS       `xml:"relationDS"`
	Parent            interface{}      `xml:"-"`
}

type DefaultValueExpr struct {
	XMLName xml.Name `xml:"defaultValueExpr"`
	Mode    string   `xml:"mode,attr"`
	Content string   `xml:",chardata"`
}

type RelationDS struct {
	XMLName        xml.Name       `xml:"relationDS"`
	RelationItemLi []RelationItem `xml:"relationItem"`
	Parent         interface{}    `xml:"-"`
}

type RelationItem struct {
	XMLName           xml.Name    `xml:"relationItem"`
	Id                string      `xml:"id"`
	RelationExpr      string      `xml:"relationExpr"`
	RelationModelId   string      `xml:"relationModelId"`
	RelationDataSetId string      `xml:"relationDataSetId"`
	DisplayField      string      `xml:"displayField"`
	Parent            interface{} `xml:"-"`
}
