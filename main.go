package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		config := NewConfig(os.Args[1])
		fmt.Printf("Process_grep using %s\n", os.Args[1])
		ProcessGrep(*config)
	}

}
