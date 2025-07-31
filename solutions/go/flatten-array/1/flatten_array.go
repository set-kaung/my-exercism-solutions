package flatten

func Flatten(nested interface{}) []interface{} {
	return recurse(nested)
}

func recurse(nested interface{}) []interface{} {
	result := []any{}
	if nested == nil {
		return result
	}
	if x, ok := nested.([]interface{}); ok {
		for _, k := range x {
			result = append(result, recurse(k)...)
		}
	} else {
		result = append(result, nested)
	}
	return result
}
