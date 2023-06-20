package json

import (
	"encoding/json"
	"fmt"
)

func Demo() {

	type Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("转换为JSON时出错:", err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)

}

func Demo1() {
	type Address struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
	}

	type Person struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	type Person1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		// 嵌套的匿名结构体, 在序列化时, 其包含的字段会被抽离并和其所属的外层结构体合并
		Address
	}

	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("转换为JSON时出错:", err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	person1 := Person1{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}

	jsonData, err = json.Marshal(person1)
	if err != nil {
		fmt.Println("转换为JSON时出错:", err)
		return
	}
	jsonString = string(jsonData)
	fmt.Println(jsonString)
}
