package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi! Enter your age: ")
	reader := bufio.NewReader(os.Stdin)

	age, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	intAge, err := strconv.ParseInt(strings.TrimSpace(age), 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Your age is: ", intAge)
}
