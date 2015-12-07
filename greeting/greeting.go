package greeting

import "fmt"

type Greeter struct {
	// 跨pkg的assignment若要印出來則需宣要為大寫，代表exported
	// 沒大寫就是unexported
	Name string
}

func Greet(name string, isFormal bool) {
	// fmt.Println(name, "This is a greeting.")
	// 在if和;中間可以加上一個initialization statement
	// 代表這個變數只會存在於branch當中
	if prefix := GetPrefix(name); isFormal {
		fmt.Println("Also, this is a formal greeting. " + prefix + name)
	}
}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "John":
		prefix = ParseTitle(name)
		// 用fallthrough的話代表下面一個case也會一併執行
		// fallthrough
	case "Mary", "Amy":
		prefix = "Ms."
	default:
		prefix = "Dude"
	}
	return
}

func ParseTitle(name string) string {
	// switch 也可以不帶變數，直接當做多重的if..else使用
	switch {
	case name == "John":
		return "Holy Dr."
	default:
		return "Dr."
	}
}
