package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
