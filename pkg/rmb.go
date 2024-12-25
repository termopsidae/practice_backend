package pkg

import (
	"fmt"
	"strconv"
)

func Yuan2Fen(yuan float64) int64 {
	// 将元转为字符串类型，保留两位小数
	yuanStr := fmt.Sprintf("%.2f", yuan)
	// 去掉小数点
	yuanStr = yuanStr[:len(yuanStr)-3] + yuanStr[len(yuanStr)-2:]
	// 将字符串转为int64类型
	fen, _ := strconv.ParseInt(yuanStr, 10, 64)
	return fen
}

func Fen2Yuan(fen int64) float64 {
	// 将分转为字符串类型
	fenStr := strconv.FormatInt(fen, 10)
	// 在字符串中插入小数点，使其变为元
	yuanStr := fenStr[:len(fenStr)-2] + "." + fenStr[len(fenStr)-2:]
	// 将字符串转为float64类型
	yuan, _ := strconv.ParseFloat(yuanStr, 64)
	return yuan
}

// Fen2YuanRetString
//
//	@Description: 分单位 转 app视图字符串 元单位 异常为0的数据 返回 "-"
func Fen2YuanRetString(fen int64, nullRetStr string) string {
	if fen <= 0 {
		return nullRetStr
	}
	str := fmt.Sprintf("%.2f", Fen2Yuan(fen))
	return str
}
