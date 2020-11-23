package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")

	nBytes, _ := fmt.Println("Hello world..")
	fmt.Println(nBytes,)
	x:=10
	fmt.Println(x)
	if x > 5 {
		fmt.Println("Great..")
	}

	if(x+5 <10){
		fmt.Println("Lesser")
	} else{
		fmt.Println("LOt better")
	}

}