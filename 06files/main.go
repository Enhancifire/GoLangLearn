package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "test.txt"

	readFile(fileName)

}

func readFile(fileName string) {
	byteData, err := ioutil.ReadFile(fileName)
	handleError(err)
	fmt.Println(string(byteData))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
