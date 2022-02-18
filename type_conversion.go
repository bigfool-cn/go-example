package go_example

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// 类型转换

// int 转 string
func int2str(val int) string {
	return strconv.Itoa(val)
}

// string 转 int
func str2int(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("str conversion int error: %v", err))
	}
	return val
}

// float64 转 string
func float2str(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

// str 转 float64
func str2float(str string) float64 {
	// bitSize-32-float32
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("str conversion float error: %v", err))
	}
	return val
}

// str 转 []byte
func str2byte(str string) []byte {
	return []byte(str)
}

// []byte 转 str
func byte2str(bty []byte) string {
	return string(bty)
}

// map 转 json
func map2json(mp map[string]interface{}) string {
	val, err := json.Marshal(mp)
	if err != nil {
		panic(fmt.Sprintf("map conversion json error: %v", err))
	}
	return string(val)
}

// json 转 map
func json2map(jsn string) map[string]interface{} {
	var val interface{}
	err := json.Unmarshal([]byte(jsn), &val)
	if err != nil {
		panic(fmt.Sprintf("json conversion map error: %v", err))
	}
	return val.(map[string]interface{})
}
