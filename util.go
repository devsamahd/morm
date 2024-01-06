package morm

import (
	"errors"
	"reflect"
	"strings"
)

// getModelType returns the reflect.Type of the model pointed to by the input pointer.
//
// Parameters:
//   - value: A pointer to the model instance.
//
// Returns:
//   - reflect.Type: The reflect.Type of the model.
//   - error: An error if the input is not a pointer.
//
// Example:
//   var user User
//   modelType, err := getModelType(&user)
//   if err != nil {
//     // Handle error
//   }
//
// This function is useful for obtaining the reflect.Type of a model from a pointer.
func getModelType(value interface{}) (reflect.Type, error) {
	resultType := reflect.TypeOf(value)

	if resultType.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}

	modelType := resultType.Elem()
	return modelType, nil
}

// getFirstStringAfterSplit returns the first string after splitting the input string by commas.
//
// Parameters:
//   - input: The input string to be split.
//
// Returns:
//   - string: The first string after splitting.
//
// Example:
//   input := "apple, orange, banana"
//   firstString := getFirstStringAfterSplit(input)
//
// This function is useful for extracting the first string from a comma-separated list.
func getFirstStringAfterSplit(input string) string {
	splitStrings := strings.Split(input, ",")
	if len(splitStrings) > 0 {
		return splitStrings[0]
	}
	return ""
}

// getTags retrieves specific tags associated with a field in a struct.
//
// Parameters:
//   - modelType: The reflect.Type of the struct.
//   - field: The field for which tags are to be retrieved.
//
// Returns:
//   - map[string]string: A map of tag names to tag values.
//   - error: An error if the field is not found.
//
// Example:
//   var user User
//   tags, err := getTags(reflect.TypeOf(user), "Name")
//   if err != nil {
//     // Handle error
//   }
//
// This function is useful for obtaining specific tags associated with a struct field.
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

// removeBrackets removes "[" and "]" from the input string.
//
// Parameters:
//   - input: The input string containing "[" and "]".
//
// Returns:
//   - string: The input string with "[" and "]" removed.
//
// Example:
//   input := "[apple]"
//   result := removeBrackets(input)
//
// This function is useful for cleaning up strings that contain square brackets.
func removeBrackets(input string) string {
	// Remove "[" and "]"
	result := strings.ReplaceAll(strings.ReplaceAll(input, "[", ""), "]", "")
	return result
}
