package main

func conditional() {
	var k = 1
	if k == 1 {
		println(1)
	}
	// if 사용법
	if k == 1 {

	} else if k == 2 {

	} else {

	}

	if val := 1 * 2; val < k {
		val++
	}
	// go는 switch에서 break를 사용하지 않아도 default로 가거나 다음 case로 가지 않음 : fallthrough를 사용하면 다음 case로 이동
	switch k := 1; k + 1 {
	case 1:
		k++
	case 2, 3: // ,를 기준으로 나열가능
		k--
		fallthrough // 다음 케이스로 이동
	default:
	}

	switch { // if else if 를 단순화 하여 표현할 수 있음
	case k > 1:
		k++
	case k < 1: // ,를 기준으로 나열가능
		k--
		fallthrough // 다음 케이스로 이동
	default:
	}

}

// switch type 체크
func typeSwitch(tst interface{}) {
	switch v := tst.(type) {
	case int:
		v++
	default:
	}
}
