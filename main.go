package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	args := os.Args
	fmt.Println(dbengine(args))
	runtime.GC()
}
