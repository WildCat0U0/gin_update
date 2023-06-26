package main

import (
	"fmt"
	"time"
)

func main() {
	//unix := time.Now().Unix()
	//fmt.Println(unix)
	t := time.Unix(1687748303, 0)
	tt := time.Unix(1687748212, 0)
	fmt.Println(t, tt)
}
