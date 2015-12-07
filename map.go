package main

import "fmt"

func main() {
	// map都是reference，帶入function會直接修改原本的值

	// 基本宣告方式
	// var prefix_map map[string]string
	// prefix_map := make(map[string]string)
	// 用等於的方式可以insert也可以update
	// prefix_map["John"] = "Great"
	// prefix_map["Bob"] = "Good"
	// prefix_map["Mary"] = "Awesome"

	// 快速宣告
	// 第一個string代表key類別，第二個string代表value類別
	// 最後一個value的逗點要留著
	prefix_map := map[string]string{
		"John": "Great",
		"Bob":  "Good",
		"Mary": "Awesome",
	}

	delete(prefix_map, "Mary")

	// map會回傳兩個值，第一個是查詢結果，第二個是確認是否存在
	// 可直接用第二個值當做判斷
	if value, exists := prefix_map["John"]; exists {
		fmt.Println(value + " exists")
	}

	// fmt.Println(prefix_map["John"])

	for i, c := range prefix_map {
		fmt.Println(i)
		fmt.Println(c)
	}
}
