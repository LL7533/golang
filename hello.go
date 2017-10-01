package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "word "
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], "  ")
	}
	fmt.Println("hello ", who)
}

//func main() {
//fmt.Printf("hello, world\n")
//}
