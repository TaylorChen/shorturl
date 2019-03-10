package main

import (
	"fmt"
	"shorturl/util"
)

var (
	ut *util.Util
)

func main() {
	ret := ut.Rate("access_rate", 2, 10)
	fmt.Println("rate ret:", ret)
}
