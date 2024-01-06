package morm

import (
	"errors"
	"reflect"
	"strings"
)

func getModelType(value interface{}) (reflect.Type, error) {
	resultType := reflect.TypeOf(value)

	if resultType.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}

	modelType := resultType.Elem()
	return modelType, nil
}

func getFirstStringAfterSplit(input string) string {
	splitStrings := strings.Split(input, ",")
	if len(splitStrings) > 0 {
		return splitStrings[0]
	}
	return ""
}

func getTags(modelType reflect.Type, field string) (map[string]string, error) {
	fieldTags := make(map[string]string)

	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)

		if structField.Name == field {
			localFieldTag := structField.Tag.Get("localField")
			foreignFieldTag := structField.Tag.Get("foreignField")
			justOneTag := structField.Tag.Get("justOne")
			countTag := structField.Tag.Get("count")
			jsonTag := structField.Tag.Get("json")

			fieldTags["localField"] = localFieldTag
			fieldTags["foreignField"] = foreignFieldTag
			fieldTags["justOne"] = justOneTag
			fieldTags["count"] = countTag
			fieldTags["json"] = jsonTag

			return fieldTags, nil
		}
	}

	return nil, errors.New("field not found")
}

func removeBrackets(input string) string {
	// Remove "[" and "]"
	result := strings.ReplaceAll(strings.ReplaceAll(input, "[", ""), "]", "")
	return result
}
