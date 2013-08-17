package component

import ()

type TemplateManager struct{}

func (qb QuerySupport) QueryDataForListTemplate(listTemplate *ListTemplate, paramMap map[string]string) map[string]interface{} {
	collection := listTemplate.DataProvider.Collection
	fixBsonQuery := listTemplate.DataProvider.FixBsonQuery
	queryParameters := listTemplate.QueryParameterGroup.QueryParameterLi
	for _, queryParameter := range queryParameters {
		if queryParameter.Restriction == "" {

		}
	}

	querySupport := QuerySupport{}
	query := fixBsonQuery
	pageNo := 1
	pageSize := 10
	return querySupport.Index(collection, query, pageNo, pageSize)
}

func (qb QuerySupport) buildQueryForQueryParameter(queryParameter QueryParameter, paramMap map[string]string) map[string]interface{} {
	queryParameterBuilder := QueryParameterBuilder{}
	return queryParameterBuilder.buildQuery(queryParameter, paramMap[queryParameter.Name])// queryParameter.ColumnName
}

func (qb QuerySupport) BuildQuery(listTemplate ListTemplate) {
}
