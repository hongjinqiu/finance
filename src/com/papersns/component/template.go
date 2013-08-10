package component

import (
	"encoding/xml"
)

type ListTemplate struct {
	XMLName        xml.Name       `xml:"list-template"`
	Adapter        Adapter        `xml:"adapter"`
	Description    string         `xml:"description"`
	Scripts        string         `xml:"scripts"`
	ViewTemplate   string         `xml:"view-template"`
	Toolbar        Toolbar        `xml:"toolbar"`
	Security       Security       `xml:"security"`
	DataProvider   DataProvider   `xml:"data-provider"`
	ColumnModel    ColumnModel    `xml:"column-model"`
	QueryParameter QueryParameter `xml:"query-parameter"`
}

type Adapter struct {
	XMLName xml.Name `xml:"adapter"`
	Name string `xml:"name,attr"`
}

type Toolbar struct {
	XMLName xml.Name `xml:"toolbar"`
	Button Button `xml:"button"`
}

type Security struct {
	XMLName xml.Name `xml:"security"`
}

type DataProvider struct {
	XMLName xml.Name `xml:"data-provider"`
}

type ColumnModel struct {
	XMLName xml.Name `xml:"column-model"`
}

type QueryParameter struct {
	XMLName xml.Name `xml:"query-parameter"`
}

/*
<attribute name="xtype" default="splitbutton" />
<attribute name="name" />
<attribute name="text" use="required" />
<attribute name="iconCls" />
<attribute name="iconAlign" default="left"></attribute>
<attribute name="disabled" type="boolean" default="true" />
<attribute name="hidden" type="boolean" default="true" />
<attribute name="arrowAlign" default="bottom"></attribute>
<attribute name="scale" default="small">
	<simpleType>
		<restriction base="string">
			<enumeration value="small" />
			<enumeration value="medium" />
			<enumeration value="large" />
		</restriction>
	</simpleType>
</attribute>
<attribute name="rowspan" type="int" default="2" />
<attribute name="handler" />
<attribute name="mode">
	<simpleType>
		<restriction base="string">
			<enumeration value="fn" />
			<enumeration value="url" />
			<enumeration value="url^" />
		</restriction>
	</simpleType>
</attribute>
*/
type Button struct {
	XMLName xml.Name `xml:"button"`
	
}
