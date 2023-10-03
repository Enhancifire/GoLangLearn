package main

import "fmt"

func main() {
	var name string = "Faiz Saiyad"

	fmt.Println(name)

	var age int = 23
	var isMale bool = true

	fmt.Println(age)
	fmt.Println(isMale)

	var instantiatedInt int // No Value assigned

	fmt.Println(instantiatedInt)

	var autoInferencedType = "what type"

	fmt.Printf("Type of autoInferencedType: %T\n", autoInferencedType)

	walrusOperator := "Walrus"
	fmt.Println(walrusOperator)

	var unusedString = "unused"
	_ = unusedString

}
