package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

var age int
var weight float64

func main() {
	fmt.Println("hello,world!")
	fmt.Println(math.Floor(2.35))
	fmt.Println(strings.Title("hello"))
	fmt.Println('H')
	fmt.Println(reflect.TypeOf(3.45))
	fmt.Println(age, weight)
	fmt.Println(reflect.TypeOf(float64(age)))
	fmt.Println(reflect.TypeOf(age))
}
