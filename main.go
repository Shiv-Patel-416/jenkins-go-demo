package main

import "fmt"

func add(a, b int)int {
	return a+b
}

func main(){
	fmt.Println("Hello from go")
	fmt.Println("10+20 = ",add(10,20))
	fmt.Println("Triggered by a GitHub push!")
}
