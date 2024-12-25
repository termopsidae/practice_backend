package pkg

import "encoding/json"

func Contains[T uint | string | int64 | int](arr []T, data T) bool {
	for _, n := range arr {
		if n == data {
			return true
		}
	}
	return false
}

// ArrStringFromJsonNoErr
//
//	@Description:格式数据转化
//	@param jsonStr 案例="["官方认证","精品"]"
//	@return []string 异常则返回空的数组
func ArrStringFromJsonNoErr(jsonStr string) []string {
	arr := make([]string, 0)
	_ = json.Unmarshal([]byte(jsonStr), &arr)
	return arr
}
