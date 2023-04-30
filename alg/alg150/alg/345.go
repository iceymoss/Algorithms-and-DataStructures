package alg

import "strings"

// 双指针法

func ReverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}
	str := []byte(s)

	m := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	first, last := 0, len(s)-1
	for first < last {
		if m[str[first]] && m[str[last]] {
			str[first], str[last] = str[last], str[first]
			first++
			last--
		} else {
			if !m[str[first]] {
				first++
			}
			if !m[str[last]] {
				last--
			}
		}
	}
	return string(str)
}

func ReverseVowels1(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}

	str := []byte(s)
	first, last := 0, n-1
	for first < last {
		for first < n && !strings.Contains("aeiouAEIOU", string(str[first])) {
			first++
		}
		for last > 0 && !strings.Contains("aeiouAEIOU", string(str[last])) {
			last++
		}
		if first < last {
			str[first], str[last] = str[last], str[first]
			first++
			last--
		}
	}
	return string(str)
}
