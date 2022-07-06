package main

import (
	"fmt"
	"time"
)

func printTime() {
	dt := time.Now()
	fmt.Println("Time: ", dt.String())
}

func main() {
	printTime()
}
