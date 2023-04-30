package alg

func AergeAlternately(word1 string, word2 string) string {
	res := make([]byte, len(word1)+len(word2))
	for i := 0; i < len(word1); i++ {
		if i >= len(word2) {
			res = append(res, word1[i:]...)
			break
		}
		if i < len(word2) {
			res = append(res, word1[i], word2[i])
		}
	}

	if len(word1) < len(word2) {
		res = append(res, word2[len(word1):]...)
	}
	return string(res)
}
