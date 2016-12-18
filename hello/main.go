package main

import (
	"fmt"

	"example.com/common"
	"example.com/hello/handlers"
)

func main() {
	fmt.Println(handlers.Greet)
	common.Show()
}
