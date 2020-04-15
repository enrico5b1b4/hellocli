package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "Name to greet.")
	flag.Parse()

	fmt.Printf("Hello: %s!\n", *name)
}
