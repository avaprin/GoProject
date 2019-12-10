//Задание 1. Программа располагает цифры из заданной последовательности в порядке убывания их частоты.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findMax(numMap map[string]int) (int, string) {
	max := 0
	ind := ""
	for i, v := range numMap {
		if max < v {
			max = v
			ind = i
		}
	}
	return max, ind
}

func main() {
	var input string
	fmt.Println("Введите последовательность (только цифры без пробелов):")
	fmt.Scan(&input)
	var nums = strings.Split(input, "")

	var err error
	for _, v := range nums {
		if _, err = strconv.Atoi(v); err != nil {
			fmt.Printf("Символ: %q не является цифрой\n", v)
		}
	}

	if err != nil {
		return
	}

numMap := make(map[string]int)
for _, v := range nums {
numMap[v]++
}

var max, ind, str = 1, "", ""
for max != 0 {
max, ind = findMax(numMap)
for i := 0; i < max; i++ {
str += ind
}
numMap[ind] = 0
}

fmt.Println(str)
}
