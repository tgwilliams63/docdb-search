package api_endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	document_connector "github.com/tgwilliams/simple-search-ui/internal/document-connector"
)

func GetFields(c *gin.Context) {
	DocConfig := GetDocConfig()
	fields := []document_connector.Field{}
	fieldName := c.Query("FieldName")
	fieldType := c.Query("Type")
	fieldSchema := c.Query("Schema")
	if fieldName == "" && fieldType == "" && fieldSchema == "" {
		fmt.Printf("The query string was nil\n")
		fields = DocConfig.GetFields()
	} else {
		var queryList [][]string
		if fieldName != "" {
			queryList = append(queryList, []string{"field_name",fieldName})
		}
		if fieldType != "" {
			queryList = append(queryList, []string{"type",fieldType})
		}
		if fieldSchema != "" {
			queryList = append(queryList, []string{"schema",fieldSchema})
		}
		fmt.Printf("The query string was not nil\n")
		fields = DocConfig.GetFieldsWithParams(queryList)
	}

	c.JSON(200, fields)
}
