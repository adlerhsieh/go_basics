package main

import "fmt"

// 變數前面加上...稱為variadic parameter，代表可以帶任何數量的變數進去
func foo(is ...int) {
	for i := 0; i < len(is); i++ {
		fmt.Println(is[i])
	}
}

func main() {
	// array size 是固定的，宣告以後就無法修改
	// array 帶入function不會修改到原本的值

	// slice是array的wrapper或pointer，points to array
	// size可以修改
	// 用make宣告

	// 括號裡面代數字就變成array
	// var s []int
	// 變數順序type、length、capacity
	// s = make([]int, 3)
	// 簡易宣告
	s := []int{0, 1, 2, 3}

	// 重新指定slice，第一個數字是開始的index，第二個數字是結束的index，並不會包含到結束的index
	// 開頭index可以省略，同等於0開始，例如：s[:4]
	// 結尾index省略則同等於包含到結尾
	s = s[2:4]

	s = append(s, 4)

	// 複製自己
	// 數列後方的...代表將數列切成數個變數，就像是Ruby的*
	// 以下方式同等於append(s,2,3,4)
	s = append(s, s...)

	// 不同的slice組合，後方slice要加上...
	d := []int{5, 6, 7, 8}
	s = append(s, d...)

	// 刪除index 2的內容
	s = append(s[:2], s[3:]...)

	fmt.Println(s)
	foo([]int{1, 2, 3, 4}...)
}
