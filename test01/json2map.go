
package main

import (
	"encoding/json"
	"fmt"
)

//检查map里面是否存在某个key，返回字符串
func mapExist(m map[string]interface{}, key string) string {
	if _, ok := m[key]; ok {
		return fmt.Sprintf("%v",m[key])
	} else {
		return ""
	}
}

//解析json字符串成 map
func jsonStringToMap(jsonStr string) (m map[string]interface{}, err error) {
	a := map[string]interface{}{}
	unmarsha1Err := json.Unmarshal([]byte(jsonStr), &a)
	if unmarsha1Err != nil {
		return nil, unmarsha1Err
	}
	fmt.Println(unmarsha1Err, "转换结果", a)
	return a, nil
}
func main() {
	str := "{\"Name\":\"alpha\",\"Age\":15}"
	m, err := jsonStringToMap(str)
	if err != nil {
		fmt.Println("jsonStringToMap出错", err)
	}

	name := mapExist(m, "Name")
	if name != "" {
		fmt.Println("key=name的值是:", name)
	} else {
		fmt.Println("key=name的值不存在！！！")
	}

	age := mapExist(m, "Age")
	if age != "" {
		fmt.Println("key=Age的值是:", age)
	} else {
		fmt.Println("key=Age的值不存在!!!")
	}

}
