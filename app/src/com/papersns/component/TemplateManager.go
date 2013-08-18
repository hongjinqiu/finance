package component

import (
	"encoding/json"
	"log"
)

type TemplateManager struct{}

func (o TemplateManager) QueryDataForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	queryMap := map[string]interface{}{}
	queryLi := []map[string]interface{}{}
	
	collection := listTemplate.DataProvider.Collection
	fixBsonQuery := listTemplate.DataProvider.FixBsonQuery
	
	fixBsonQueryMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(fixBsonQuery), &fixBsonQueryMap)
	if err != nil {
		panic(err)
	}

	queryLi = append(queryLi, fixBsonQueryMap)	
	
	queryParameters := listTemplate.QueryParameterGroup.QueryParameterLi
	queryParameterBuilder := QueryParameterBuilder{}
	for _, queryParameter := range queryParameters {
		if queryParameter.Editor != "" && queryParameter.Restriction != "" {
			name := queryParameter.Name
			if paramMap[name] != "" {
				queryParameterMap := queryParameterBuilder.buildQuery(queryParameter, paramMap[name])
				queryLi = append(queryLi, queryParameterMap)
			}
		}
	}

	querySupport := QuerySupport{}
	queryMap["$and"] = queryLi

	queryByte, err := json.MarshalIndent(queryMap, "", "\t")
	log.Println("QueryDataForListTemplate,query is:" + string(queryByte))

	return querySupport.Index(collection, queryMap, pageNo, pageSize)
}


