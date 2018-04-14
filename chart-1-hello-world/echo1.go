package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

func main() {
	var s, sep string
	startLow := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("spnd time:%d\n", time.Since(startLow))

	startLow = time.Now()
	result := strings.Join(os.Args, " ")

	fmt.Printf("hight mode spend time:%d\n", time.Since(startLow))
	fmt.Printf("paramter:%s", result)

	for index, arg := range os.Args[1:] {
		fmt.Printf("索引:%d,值:%s", index, arg)
	}

	//fmt.Println(s)
}
