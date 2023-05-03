package alg

import (
	"strconv"
)

func Compress(chars []byte) int {
	//双指针法
	if len(chars) <= 0 {
		return 0
	}
	s := make([]byte, 0)
	lows, fast := 0, 0
	for fast <= len(chars) {
		if fast == len(chars) {
			pTotal := fast - lows
			if pTotal == 1 {
				s = append(s, chars[lows])
				lows = fast
				break
			}
			if pTotal >= 10 {
				s = append(s, chars[lows])
				s = append(s, []byte(strconv.Itoa(pTotal))...)
				lows = fast
				break
			}
			s = append(s, chars[lows])
			s = append(s, []byte(strconv.Itoa(pTotal))...)
			lows = fast
			break
		}
		if chars[lows] != chars[fast] {
			pTotal := fast - lows
			if pTotal == 1 {
				s = append(s, chars[lows])
				lows = fast
				continue
			}
			if pTotal >= 10 {
				s = append(s, chars[lows])
				s = append(s, []byte(strconv.Itoa(pTotal))...)
				lows = fast
				continue
			}
			s = append(s, chars[lows])
			s = append(s, []byte(strconv.Itoa(pTotal))...)
			lows = fast
			continue
		}
		fast++
	}
	for i := 0; i < len(s); i++ {
		chars[i] = s[i]
	}
	chars = chars[:len(s)]
	return len(chars)
}
