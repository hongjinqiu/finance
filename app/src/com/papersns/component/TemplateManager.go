package component

import (
	"encoding/json"
)

type TemplateManager struct{}

func (o TemplateManager) QueryDataForListTemplate(listTemplate *ListTemplate, paramMap map[string]string, pageNo int, pageSize int) map[string]interface{} {
	queryMap := map[string]interface{}{}
	
	collection := listTemplate.DataProvider.Collection
	fixBsonQuery := listTemplate.DataProvider.FixBsonQuery
	
	fixBsonQueryMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(fixBsonQuery), &fixBsonQueryMap)
	if err != nil {
		panic(err)
	}
	
	for k,v := range fixBsonQueryMap {
		queryMap[k] = v
	}
	
	queryParameters := listTemplate.QueryParameterGroup.QueryParameterLi
	queryParameterBuilder := QueryParameterBuilder{}
	for _, queryParameter := range queryParameters {
		if queryParameter.Editor != "" && queryParameter.Restriction != "" {
			name := queryParameterBuilder.GetQueryName(queryParameter)
			if paramMap[name] != "" {
				queryParameterMap := queryParameterBuilder.buildQuery(queryParameter, paramMap[name])
				for k,v := range queryParameterMap {
					queryMap[k] = v
				}
			}
		}
	}

	querySupport := QuerySupport{}
	return querySupport.Index(collection, queryMap, pageNo, pageSize)
}


