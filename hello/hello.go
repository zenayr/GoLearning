package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
// func main() {
// 	fmt.Println(quote.Go())
// }
