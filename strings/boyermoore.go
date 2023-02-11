package strs

var NO_OF_CHARS = 256

type BoyerMoore struct{}

func (this *BoyerMoore) max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (this *BoyerMoore) badCharHeuristic(str []rune, size int, badchar []int) {
	var i int

	for i = 0; i < NO_OF_CHARS; i++ {
		badchar[i] = -1
	}

	for i = 0; i < size; i++ {
		badchar[str[i]] = i
	}
}

func (this *BoyerMoore) search(text string, part string) (bool, []int) {
	textRunes := []rune(text)
	partRunes := []rune(part)

	n := len(textRunes)
	m := len(partRunes)

	badchar := make([]int, NO_OF_CHARS)

	this.badCharHeuristic(partRunes, m, badchar)

	resultStartIndexes := []int{}

	s := 0

	for s <= (n - m) {
		j := m - 1

		for j >= 0 && partRunes[j] == textRunes[s+j] {
			j--
		}

		if j < 0 {
			resultStartIndexes = append(resultStartIndexes, s)

			if s+m < n {
				s += m - badchar[textRunes[s+m]]
			} else {
				s += 1
			}
		} else {
			s += this.max(1, j-badchar[textRunes[s+j]])
		}
	}

	if len(resultStartIndexes) > 0 {
		return true, resultStartIndexes
	}

	return false, resultStartIndexes
}

func NewBM() *BoyerMoore {
	return &BoyerMoore{}
}

type BoyerMooreUpdated struct{}

func (this *BoyerMooreUpdated) Find(text, pattern string) int {
	t := 0
	last := len(pattern) - 1

	for t < len(text)-last {
		p := 0

		for p <= last && text[t+p] == pattern[p] {
			p++
		}

		if p == len(pattern) {
			return t
		}

		t++
	}

	return -1
}

func (this *BoyerMooreUpdated) Find2(text, pattern string) int {
	t := 0
	last := len(pattern) - 1

	for t < len(text)-last {
		p := last

		for p >= 0 && text[t+p] == pattern[p] {
			p--
		}

		if p == -1 {
			return t
		}

		t++
	}

	return -1
}

func (this *BoyerMooreUpdated) Find3(text, pattern string) int {
	shift := this.CreateShift(pattern)
	t := 0
	last := len(pattern) - 1

	for t < len(text)-last {
		p := last
		for p >= 0 && text[t+p] == pattern[p] {
			p--
		}

		if p == -1 {
			return t
		}

		t += shift[text[t+last]]
	}

	return -1
}

func (this *BoyerMooreUpdated) CreateShift(pattern string) []int {
	shift := make([]int, 128)

	for j := 0; j < len(shift); j++ {
		shift[j] = len(pattern)
	}

	for p := 0; p < len(pattern)-1; p++ {
		shift[pattern[p]] = len(pattern) - p - 1
	}

	return shift
}

func NewBMUpdated() *BoyerMooreUpdated {
	return &BoyerMooreUpdated{}
}
