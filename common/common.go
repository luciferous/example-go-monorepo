package common

import (
	"fmt"
	"golang.org/x/net/trace"
)

func Show() {
	fmt.Println(trace.New("family", "title"))
}
