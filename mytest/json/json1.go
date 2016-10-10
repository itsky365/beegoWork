package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//bolB, _ := json.Marshal(true)
	//fmt.Println(string(bolB))
	//intB, _ := json.Marshal(1)
	//fmt.Println(string(intB))
	//fltB, _ := json.Marshal(2.34)
	//fmt.Println(string(fltB))
	//strB, _ := json.Marshal("gopher")
	//fmt.Println(string(strB))
	//// 这里是将切片和字典编码为JSON数组或对象
	//slcD := []string{"apple", "peach", "pear"}
	//slcB, _ := json.Marshal(slcD)
	//fmt.Println(string(slcB))
	//mapD := map[string]int{"apple": 5, "lettuce": 7}
	//mapB, _ := json.Marshal(mapD)
	//fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// 你可以使用tag来自定义编码后JSON键的名称
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// 现在我们看看解码JSON数据为Go数值
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// 我们需要提供一个变量来存储解码后的JSON数据,这里
	// 的`map[string]interface{}`将以Key-Value的方式
	// 保存解码后的数据,Value可以为任意数据类型
	var dat map[string]interface{}
	// 解码过程,并检测相关可能存在的错误
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["strs"])
    //for k, v := range dat["strs"] {
    //    fmt.Println(k, v)
    //}

    num := dat["num"].(float64)
    fmt.Println(num)
    // 访问嵌套的数据需要一些类型转换
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
}
