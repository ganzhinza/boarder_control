package main

import "fmt"

func main(){
	var s string
	fmt.Scanln(&s)
	switch s {
	case "en":
		fmt.Println("Hello, world!")
	case "ru":
		fmt.Println("Привет, мир!")	
	}
	
}