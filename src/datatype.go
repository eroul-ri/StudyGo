package main

func main() {
	var real bool = true
	var number int = 0
	var duble complex64 = 0
	/*
		byte == unit : type mismatched error
		byte == uint8 : 바이트 코드에 사용
		Go의 경우 묵시적 형 변환은 발생 X 명시적 O
		명시적 형변환 : type(value)
	*/
	var byt byte = 0
	var uni uint8 = 0
	println(real)
	println(number)
	println(duble)
	println(byt == uni)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	rawLtr := `이거슨
			  자바스크립트의 백틱과 동일하게
			  동작한다.`
	itrLtr := "\n" + "기존 문자열 연산처럼 사용할 수 있다."
	println(rawLtr)
	println(itrLtr)
}
