package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	switch s {
	case "en":
		fmt.Println("Hello, world!")
	case "ru":
		fmt.Println("Привет, мир!")
	case "fr":
		fmt.Println("Bonjour le monde!")
	case "ki":
		fmt.Println("Haro ensi yoona!!")
	case "ewe":
		fmt.Println("Mido gbe na xexeame!")
	case "ch":
		fmt.Println("你好世界！")
	case "kr":
		fmt.Println("안녕하세요!")
	case "es":
		fmt.Println("Saluton mondo!")
	}

}
