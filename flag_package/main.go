package main

import (
	"flag"
	"fmt"
)

func main() {
	textPtr := flag.String("text", "", "Text to parse")
	intPtr := flag.Int("integer", 0, "this int is not used")

	flag.Parse()

	fmt.Println(*textPtr)
	fmt.Println(*intPtr)

	for i, flag := range flag.Args() {
		fmt.Println(i, "-", flag)
	}
}
