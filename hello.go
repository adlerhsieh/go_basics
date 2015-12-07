package main

import "fmt"

func main() {
	// 宣告變數的三種方式
	// var message string = "hello"
	// var message = "hello"
	message := "hello"

	const PI = 3.14
	const lang string = "Go"

	const (
		// iota 會儲存宣告的順序，A = 2
		A = iota
		// 不宣告內容的話會繼承上一個內容，B = iota
		B
	)

	var a, b, c = 1, 2, 3
	// fmt.Println(message, a, b, c)

	// 星號代表他是pointer
	// &號代表要把他指向message這個變數
	var greeting *string = &message
	// 讀取時greeting是抓取memory的位置，*greeting是抓取他的value
	// * 是reference、& 是dereference的概念
	fmt.Println(message, *greeting)

	fmt.Println(PI, lang, A, B)
	fmt.Println(a, b, c)
}
