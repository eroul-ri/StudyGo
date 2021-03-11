package main

/*
	1. 산술 연산자
	2. 관계 연산자
	3. 논리 연산자
	4. Bitwise 연산자
	5. 할당   연산자
*/
func separator() {
	var a int = 0
	var b int = 10
	var c = (a + b) / 5
	println(c)
	c++
	var k int = 10
	var p = &k  // k의 주소
	println(*p) // p가 가르키는 주소에 있는 내용
}
