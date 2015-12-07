package main

import "fmt"

func BasicLoop(times int) {
	fmt.Print("Basic loop")
	for i := 0; i < times; i++ {
		fmt.Print(".")
	}
	fmt.Println("End")
}

func WhileLoop(times int) {
	i := 0
	fmt.Print("While loop")
	for i < times {
		if i%2 > 0 {
			fmt.Print("odd")
			i++
			// 使用continue關鍵字可以跳到下一個loop
			continue
		}
		fmt.Print(".")
		i++
	}
	fmt.Println("End")
}

func InfiniteLoop() {
	for {
		fmt.Println("This loop never stops until a 'break' statement.")
		break
	}
}

func ForEach() {
	array := [5]int{}
	// i 是index，s是內容
	// 若沒有要用i，可用_代替
	for i, s := range array {
		fmt.Print(i)
		fmt.Print(s)
	}
	fmt.Println(" End ForEach Loop")
}

func main() {
	BasicLoop(5)
	WhileLoop(10)
	InfiniteLoop()
	ForEach()
}
