package main

import "fmt"

type Salute struct {
  name     string
  greeting string
}

// function 的變數也要宣告type
func Say(salute Salute) {
  // fmt.Println(Message(salute.name, salute.greeting))
  // _底線可以宣告變數，但允許不使用
  message_1, _ := Message(salute.name, salute.greeting)
  fmt.Println(message_1)
}

// 可在首行末設定回傳的type
// 可設定複數回傳值
// func Message(name, greeting string) (string, string) {
//  return name + " " + greeting, "Another Hi!"
// }

// 可命名回傳值並在結尾直接return兩個
func Message(name, greeting string) (m1 string, m2 string) {
  m1 = name + " " + greeting
  m2 = "Another Hi!"
  return
}

// 變數結尾可加上...代表有多個變數會帶入
func MultipleMessage(name string, greeting ...string) {
  fmt.Println(len(greeting))
  fmt.Println(greeting[0])
}

// 可在變數內帶入function
func TryFuncVariable(name string, do func(string)) {
  do(name)
}

func AnotherPrint(name string) {
  fmt.Println(name)
}

// 把func的變數變成type
type Printer func(string) ()
func TryAnotherFuncVariable(do Printer) {
  do("success")
}

// nested functions
// 回傳值的function要先定義
func CreatePrinterFunc(custom string) Printer {
  return func(s string) {
    fmt.Println(s + custom)
  }
}


func main() {
  greeting := Salute{"John", "Hi!"}
  Say(greeting)
  MultipleMessage("Bella", "Come on")
  TryFuncVariable("Jessie", AnotherPrint)
  TryAnotherFuncVariable(AnotherPrint)
  TryFuncVariable("Fred", CreatePrinterFunc("!!!"))
}