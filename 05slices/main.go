package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	errorChecker(err)

	fmt.Printf("Value is %v", text)

	faiz := User{"Faiz", 21}

	fmt.Println(faiz.getName())
}

type User struct {
	Name string
	Age  int
}

func (u *User) getName() string {
	return u.Name
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
