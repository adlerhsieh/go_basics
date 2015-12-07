package main

import "fmt"

// method 基本上是綁在type上的，只能對該type使用，type要先宣告
type Greeting struct {
	method string
	name   string
}

// 綁method時，要在function名稱前面加上(帶入的變數+type)代表該func是一個method
func (g Greeting) Say() {
	fmt.Println(g.method + " " + g.name)
}

// 帶入變數時，僅帶入copy，如果要改變原本struct內的值，必須帶入pointer，也就是在func變數上加上*
func (g *Greeting) Rename(NewName string) {
	g.name = NewName
}

// interface所宣告的內容，只要原本帶入的type有相同的method，interface會直接對應到該method上
type Renamable interface {
	Rename(NewName string)
}

func RenameToFrog(r Renamable) {
	r.Rename("Frog")
}

// interface只要是空的，所有帶入的type都會帶入自己的method
// 例如要判斷一個變數的型別，就可以帶入一個空的interface，由於interface可以跟任何型別串接，所以可以直接帶入
// 但interface依然不能跟其他型別直接互動，例如 x + 1
func TypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("something else")
	}
}

func main() {
	var greet = Greeting{"Hello, ", "John"}
	// 執行時，直接用物件導向的method call的方式即可，不需要再用function call的方式
	greet.Say()

	greet.Rename("Bob")
	fmt.Println("New name: " + greet.name)

	// 帶入到interface時，因為原本的method是用*reference帶入，必須用&dereference才能修改到原本的值
	var greet_2 = Greeting{"Hi, ", "Joy"}
	RenameToFrog(&greet_2)
	fmt.Println("New name: " + greet_2.name)
	TypeSwitch(5)
}
