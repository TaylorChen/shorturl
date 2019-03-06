package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"shorturl/util"
)

var (
	ut *util.Util
)

func main() {
	hasher := murmur3.New32()
	hasher.Write([]byte("https://www.google.com/search?newwindow=1&safe=active&biw=1920&bih=965&ei=LyV_XN-8J_bF0PEPh7eHgA8&q=golang+murmurhash3+demo&oq=golang+murmurhash3+demo&gs_l=psy-ab.3...218674.222141..222658...1.0..0.229.1293.2-6......0....1..gws-wiz.......0i71j35i39j0i22i30j33i160.XXghBoIj_6E"))
	num := hasher.Sum32()
	fmt.Println(num)
	str := ut.DecimalTo62(num)
	fmt.Println(str)
	fmt.Println(ut.ReverseToDecimal(str))
}
