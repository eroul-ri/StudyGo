package main

const CT = 10
const (
	Apple = iota // iota : identifier iota가 지정된 변수에는 0이 할당 이후 순차적으로 1씩 증가된 값을 할당
	Grape
	Orange
)

var g int = 6

/*
	1. 변수선언 및 초기화
	- 선언된 변수가 사용되지 않으면 에러
	- 선언 후 초기화를 하지 않은 경우 "Zero Value" 할당 numeric : 0, bool : false, string : ""
*/
func main() {

	var i int           // set 0
	var f float32 = 11. // set
	var a, b, c int
	var j, k int = 1, 2

	println("i = ", i)
	println("f = ", f)
	println("a = ", a)
	println("b = ", b)
	println("c = ", c)
	println("j = ", j)
	println("k = ", k)

	d := 12

	println("d = ", d)
	println("g = ", g)
	println("CT = ", CT)
	println("fruit = ", Apple, Grape, Orange)
}
