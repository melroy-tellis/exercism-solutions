package flatten

func Flatten(nested interface{}) []interface{} {
    result := []interface{}{}
    if nested != nil {
        if list, ok := nested.([]interface{}); ok {
            for _, item := range(list) {
            	result = append(result, Flatten(item)...)    
            }
        } else {
            result = append(result, nested)
        }
        
    }
    return result
}
