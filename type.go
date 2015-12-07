package main

import "fmt"

// User Defined Type
type Hello string

// struct很像Ruby的struct可以自定義變數和method
// 但這就是一個type，建立在struct上進行修改的type
type Salute struct {
	name     string
	greeting string
}

func main() {
	var m Hello = "Hello!"
	fmt.Println(m)

	// var s = Salute{"John", "Hi!"}
	// var s = Salute{name: "John", greeting: "Hi!"}
	var s = Salute{name: "John"}
	s.name = "Bob"
	s.greeting = "Hi!"
	fmt.Println(s.name, s.greeting)
}
