package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson(t *testing.T) {
	test := Test{
		Name: "张三",
		Age:  18,
	}
	buf, _ := json.Marshal(test)
	// [123 34 110 97 109 101 34 58 34 229 188 160 228 184 137 34 44 34 97 103 101 34 58 49 56 125]
	fmt.Println("对象转换成字节：", buf)

	var tes Test
	err := json.Unmarshal(buf, &tes)
	if err != nil {
		fmt.Println("json unmarshal error：", err)
	}
	// {张三 18}
	fmt.Println("字节转换成对象：", tes)
}
