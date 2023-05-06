package ArrayAndString

//本题的核心：1.将字符串去冗余空格；2.反正字符串；3.反转单词

func ReverseWords(s string) string {
	str := []byte(s)
	//将字符串去冗余空格
	outEmptyFormat(&str)

	//反转字符串
	reversed(&str, 0, len(str)-1)

	//反转单词
	for i, j := 0, 0; i < len(str); i++ {
		if i > 0 && str[i] == ' ' {
			reversed(&str, j, i-1)
			j = i
			j++
		}
		if i == len(str)-1 {
			reversed(&str, j, i)
		}
	}

	return string(str)
}

func outEmptyFormat(str *[]byte) {
	//快慢指针，去除冗余字符
	//slow用来记录去除冗余后的字符
	//fast用于判断冗余空格
	slow, fast := 0, 0

	//去除头部冗余：将fast指向不为空格的地方
	for fast < len(*str) && (*str)[fast] == ' ' {
		fast++
	}

	//去除单词间冗余
	for ; fast < len(*str); fast++ {
		if fast-1 > 0 && (*str)[fast-1] == (*str)[fast] && (*str)[fast] == ' ' {
			continue
		}
		//当fast遇到不为空的字符后，需要将
		(*str)[slow] = (*str)[fast]
		slow++
	}

	//去除尾部冗余
	if slow-1 > 0 && (*str)[slow-1] == ' ' {
		*str = (*str)[:slow-1]
	} else {
		*str = (*str)[:slow]
	}
}

func reversed(str *[]byte, first, last int) {
	//反转字符串
	//双指针法
	for first <= last {
		(*str)[first], (*str)[last] = (*str)[last], (*str)[first]
		first++
		last--
	}
}
