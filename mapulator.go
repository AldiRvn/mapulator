package mapulator

import (
	"strings"
)

// Godoc Set
func Set(target map[string]interface{}, path string, value interface{}) map[string]interface{} {
	separator := "."

	var temp map[string]interface{}
	listPath := strings.Split(path, separator)
	for i, pathItem := range listPath {
		if len(listPath) == 1 {
			target[pathItem] = value
			break
		}

		if temp == nil {
			if target[pathItem] == nil {
				target[pathItem] = map[string]interface{}{}
			}
			temp = target[pathItem].(map[string]interface{})
		} else if i+1 == len(listPath) {
			temp[pathItem] = value
		} else {
			if temp[pathItem] == nil {
				temp[pathItem] = map[string]interface{}{}
			}
			temp = temp[pathItem].(map[string]interface{})
		}
		_ = i
		// fmt.Println(i, pathItem, temp)
	}

	return target
}

// Godoc SetNew
func SetNew(path string, value interface{}) map[string]interface{} {
	return Set(map[string]interface{}{}, path, value)
}
