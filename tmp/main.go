package main

import (
	"flag"
	"fmt"
)

// D/

func main() {
	i := flag.Int64("i", -1, "no usage")
	i2 := flag.Int64("i2", -1, "the i2 usage text")
	flag.Usage()
	flag.Parse()

	fmt.Println("i:", *i, *i2)
}
