package component

import (
	"encoding/xml"
)

type ListTemplate struct {
	XMLName      xml.Name     `xml:"list-template"`
	ListId       string       `xml:"list-id"`
	SelectorId   string       `xml:"selector-id"`
	Adapter      Adapter      `xml:"adapter"`
	Description  string       `xml:"description"`
	Scripts      string       `xml:"scripts"`
	ViewTemplate ViewTemplate `xml:"view-template"`
	Toolbar      Toolbar      `xml:"toolbar"`
	Security     Security     `xml:"security"`
	Pyscript         string `xml:"pyscript"`
	BeforeBuildQuery string `xml:"before-build-query"`
	AfterBuildQuery  string `xml:"after-build-query"`
	AfterQueryData   string `xml:"after-query-data"`
	DataProvider        DataProvider        `xml:"data-provider"`
	ColumnModel         ColumnModel         `xml:"column-model"`
	QueryParameterGroup QueryParameterGroup `xml:"query-parameters"`
}

type Adapter struct {
	XMLName xml.Name `xml:"adapter"`
	Name    string   `xml:"name,attr"`
}

type ViewTemplate struct {
	XMLName xml.Name `xml:"view-template"`
	View    string   `xml:"view,attr"`
}

type Toolbar struct {
	XMLName     xml.Name    `xml:"toolbar"`
	ButtonGroup ButtonGroup `xml:"button-group"`
	ButtonLi    []Button    `xml:",any"`

	Export         string `xml:"export,attr"`
	Exporter       string `xml:"exporter,attr"`
	ExportParam    string `xml:"exportParam,attr"`
	FreezedHeader  string `xml:"freezedHeader,attr"`
	ExportChart    string `xml:"exportChart,attr"`
	ExcelChart     string `xml:"excelChart,attr"`
	ExcelChartType string `xml:"excelChartType,attr"`
	ExportTitle    string `xml:"exportTitle,attr"`
	ExportSuffix   string `xml:"exportSuffix,attr"`
}

type Security struct {
	XMLName xml.Name `xml:"security"`

	FunctionId            string `xml:"functionId,attr"`
	Override              string `xml:"override,attr"`
	DEFAULT_RESOURCE_CODE string `xml:"DEFAULT_RESOURCE_CODE,attr"`
}

type DataProvider struct {
	XMLName xml.Name `xml:"data-provider"`

	Collection    string `xml:"collection"`
	FixBsonQuery  string `xml:"fix-bson-query"`
	Map           string `xml:"map"`
	Reduce        string `xml:"reduce"`
	Size          string `xml:"size,attr"`
	BsonIntercept string `xml:"bsonIntercept,attr"`
}

type ColumnModel struct {
	XMLName xml.Name `xml:"column-model"`

	CheckboxColumn CheckboxColumn `xml:"checkbox-column"`
	IdColumn       IdColumn       `xml:"id-column"`
	ColumnLi       []Column       `xml:",any"`

	ColumnModelAttributeGroup
}

type CheckboxColumn struct {
	XMLName xml.Name `xml:"checkbox-column"`

	ColumnAttribute ColumnAttribute `xml:"column-attribute"`
	Expression      string          `xml:"expression"`
	Hideable        string          `xml:"hideable,attr"`
}

type ColumnAttribute struct {
	XMLName xml.Name `xml:"column-attribute"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type IdColumn struct {
	XMLName xml.Name `xml:"id-column"`

	ColumnAttributeGroup
}

type ColumnModelAttributeGroup struct {
	AutoLoad              string `xml:"autoLoad,attr"`
	SummaryLoad           string `xml:"summaryLoad,attr"`
	SummaryStat           string `xml:"summaryStat,attr"`
	Summation             string `xml:"summation,attr"`
	GroupSummation        string `xml:"groupSummation,attr"`
	GroupMerge            string `xml:"groupMerge,attr"`
	ShowGroupFilter       string `xml:"showGroupFilter,attr"`
	ShowAggregationFilter string `xml:"showAggregationFilter,attr"`
	AutoRowHeight         string `xml:"autoRowHeight,attr"`
	Nowrap                string `xml:"nowrap,attr"`
	Rownumber             string `xml:"rownumber,attr"`
	SelectionMode         string `xml:"selectionMode,attr"`
	ShowClearBtn          string `xml:"showClearBtn,attr"`
	SelectionSupport      string `xml:"selectionSupport,attr"`
	GroupField            string `xml:"groupField,attr"`
	BsonOrderBy           string `xml:"bsonOrderBy,attr"`
	SaveUrl               string `xml:"saveUrl,attr"`
	DeleteUrl             string `xml:"deleteUrl,attr"`
	StoreIntercept        string `xml:"storeIntercept,attr"`
	RecordIntercept       string `xml:"recordIntercept,attr"`
	SelectionTemplate     string `xml:"selectionTemplate,attr"`
	SelectionTitle        string `xml:"selectionTitle,attr"`
}

type ColumnAttributeGroup struct {
	Name             string `xml:"name,attr"`
	Bson             string `xml:"bson,attr"`
	Text             string `xml:"text,attr"`
	Align            string `xml:"align,attr"`
	Graggable        string `xml:"graggable,attr"`
	Groupable        string `xml:"groupable,attr"`
	Hideable         string `xml:"hideable,attr"`
	Editable         string `xml:"editable,attr"`
	MenuDisabled     string `xml:"menuDisabled,attr"`
	Sortable         string `xml:"sortable,attr"`
	Comparable       string `xml:"comparable,attr"`
	Locked           string `xml:"locked,attr"`
	Width            string `xml:"width,attr"`
	ExcelWidth       string `xml:"excelWidth,attr"`
	Renderer         string `xml:"renderer,attr"`
	RendererTemplate string `xml:"rendererTemplate,attr"`
	SummaryText      string `xml:"summaryText,attr"`
	SummaryType      string `xml:"summaryType,attr"`
	Cycle            string `xml:"cycle,attr"`
	Exported         string `xml:"exported,attr"`
}

type Column struct {
	XMLName xml.Name `xml:""` // 有可能是string-column,number-column,date-column,boolean-column,dictionary-column,virtual-column,
	Name    string   `xml:"name,attr"`
	ColumnAttribute
	Editor
	ColumnAttributeGroup

	Format         string `xml:"format,attr"`
	DisplayPattern string `xml:"displayPattern,attr"`
	DbPattern      string `xml:"dbPattern,attr"`
	BooleanColumnAttributeGroup

	Dictionary string `xml:"dictionary,attr"`
	Complex    string `xml:"complex,attr"`

	Buttons Buttons `xml:"buttons"`
}

type Editor struct {
	XMLName         xml.Name        `xml:"editor"`
	EditorAttribute EditorAttribute `xml:"editor_attribute"`
	Name            string          `xml:"name,attr"`
}

type EditorAttribute struct {
	XMLName xml.Name `xml:"editor_attribute"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type BooleanColumnAttributeGroup struct {
	TrueText      string `xml:"trueText,attr"`
	FalseText     string `xml:"falseText,attr"`
	UndefinedText string `xml:"undefinedText,attr"`
}

type QueryParameterGroup struct {
	XMLName xml.Name `xml:"query-parameters"`

	FixedParameterLi []FixedParameter `xml:"fixed-parameter"`
	QueryParameterLi []QueryParameter `xml:"query-parameter"`

	FormColumns      string `xml:"formColumns,attr"`
	EnableEnterParam string `xml:"enableEnterParam,attr"`
}

type FixedParameter struct {
	XMLName xml.Name `xml:"fixed-parameter"`

	QueryParamAttributeGroup
}

type QueryParameter struct {
	XMLName xml.Name `xml:"query-parameter"`

	ParameterAttributeLi []ParameterAttribute `xml:"parameter-attribute"`
	QueryParamAttributeGroup
}

type ParameterAttribute struct {
	XMLName xml.Name `xml:"parameter-attribute"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type QueryParamAttributeGroup struct {
	Name        string `xml:"name,attr"`
	Text        string `xml:"text,attr"`
	ColumnName  string `xml:"columnName,attr"`
	EnterParam  string `xml:"enterParam,attr"`
	Editor      string `xml:"editor,attr"`
	Restriction string `xml:"restriction,attr"`
	ColSpan     string `xml:"colSpan,attr"`
	RowSpan     string `xml:"rowSpan,attr"`
	Value       string `xml:"value,attr"`
	OtherName   string `xml:"otherName,attr"`
	Having      string `xml:"having,attr"`
	Required    string `xml:"required,attr"`
	UseIn       string `xml:"use-in,attr"`
}

type ButtonGroup struct {
	XMLName  xml.Name `xml:"button-group"`
	ButtonLi []Button `xml:",any"`
}

type Buttons struct {
	XMLName xml.Name `xml:"buttons"`

	ButtonLi []Button `xml:"button"`
}

type Button struct {
	XMLName         xml.Name        `xml:""` // 有可能是button,也有可能是split-button
	Expression      string          `xml:"expression"`
	ButtonAttribute ButtonAttribute `xml:"button-attribute"`
	ButtonAttributeGroup
}

type ButtonAttribute struct {
	XMLName xml.Name `xml:"button-attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type ButtonAttributeGroup struct {
	Xtype      string `xml:"xtype,attr"`
	Name       string `xml:"name,attr"`
	Text       string `xml:"text,attr"`
	IconCls    string `xml:"iconCls,attr"`
	IconAlign  string `xml:"iconAlign,attr"`
	Disabled   string `xml:"disabled,attr"`
	Hidden     string `xml:"hidden,attr"`
	ArrowAlign string `xml:"arrowAlign,attr"`
	Scale      string `xml:"scale,attr"`
	Rowspan    string `xml:"rowspan,attr"`
	Handler    string `xml:"handler,attr"`
	Mode       string `xml:"mode,attr"`
}
