package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("qqf: milli <number of milliseconds>")
		os.Exit(1)
	}
	str := os.Args[1]
	number, err := strconv.Atoi(str)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(time.Duration(number) * time.Millisecond)
}
