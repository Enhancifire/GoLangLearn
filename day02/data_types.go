package main

import (
	"bufio"
	"fmt"
	"os"
	// "slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
