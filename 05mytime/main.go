package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("welcome to time class")
	currentTime := time.Now()

	fmt.Println(currentTime)

	reformat := currentTime.Format("01-02-2023 Monday 15:03:03")
	fmt.Println(reformat)

	fmt.Println(time.Now().UTC())

	dt := time.Date(2024, 05, 12, 07, 20, 11, 11, time.UTC)

	milli := dt.UnixMilli()
	fmt.Println(milli)

	fmt.Println("no of CPU processors available", runtime.NumCPU())
	//runtime.GC()
	fmt.Println("go Arch.", runtime.GOARCH)
}
