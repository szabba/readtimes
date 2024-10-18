package main

import (
	"flag"
	"fmt"
	"time"
)

var format string = "15:04"

func main() {

	flag.StringVar(&format, "fmt", "15 04", "Time format to use. Specified as a Go time package layout.")

	flag.Parse()

	fmt.Printf("It is now %s.", time.Now().Format(format))
}
