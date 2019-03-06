package util

import (
	"fmt"
	"math"
	"strings"
)

var decimalStrMaps map[uint32]string = map[uint32]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

func (u *Util) DecimalTo62(num uint32) string {
	new_num_str := ""
	var remainder uint32
	var remainder_string string
	for num != 0 {
		remainder = num % 62
		if remainder > 9 {
			remainder_string = decimalStrMaps[remainder]
		} else {
			remainder_string = fmt.Sprint(remainder)
		}
		new_num_str = new_num_str + remainder_string
		num = num / 62
	}
	return new_num_str
}

func findKeyByValule(str string) uint32 {
	var sS uint32
	sS = 99
	for k, v := range decimalStrMaps {
		if str == v {
			sS = k
		}
	}
	return sS
}

func (u *Util) ReverseToDecimal(str string) uint32 {
	var new_num uint32
	for idx, value := range strings.Split(str, "") {
		tmp := float64(findKeyByValule(value))
		if tmp != 99 {
			new_num = new_num + uint32(tmp*math.Pow(float64(62), float64(idx)))
		} else {
			break
		}

	}
	return new_num
}
