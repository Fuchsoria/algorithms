package strs

type KMP struct{}

func (this *KMP) Search(text, pattern string) (bool, []int) {
	M := len(pattern)
	N := len(text)

	lps := make([]int, M)
	j := 0

	resultStartIndexes := []int{}

	this.computeLPSArray(pattern, M, lps)

	i := 0
	for (N - i) >= (M - j) {
		if pattern[j] == text[i] {
			j++
			i++
		}

		if j == M {
			resultStartIndexes = append(resultStartIndexes, i-j)

			j = lps[j-1]
		} else if i < N && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}

	if len(resultStartIndexes) > 0 {
		return true, resultStartIndexes
	}

	return false, resultStartIndexes
}

func (this *KMP) computeLPSArray(pattern string, M int, lps []int) {
	len := 0
	i := 1
	lps[0] = 0

	for i < M {
		if pattern[i] == pattern[len] {
			len++
			lps[i] = len
			i++
		} else {
			if len != 0 {
				len = lps[len-1]
			} else {
				lps[i] = len
				i++
			}
		}
	}
}

func NewKMP() *KMP {
	return &KMP{}
}
