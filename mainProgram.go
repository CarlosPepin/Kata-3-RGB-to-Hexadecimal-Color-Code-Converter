package main

import (
	"fmt"
	"strconv"
)

type RGB struct {
	red, green, blue int64
}

type HEX struct {
	str string
}

//Funcion de prueba
func Sum(x int, y int) int {
	return x + y
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func (color RGB) Rgb2Hex() HEX {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return HEX{r + g + b}
}

func main() {
	color1 := RGB{56, 206, 255}
	fmt.Println(color1.Rgb2Hex().str)
}
