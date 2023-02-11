package compression

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Compression struct{}

func (c *Compression) Encode(str string) string {
	n := len(str)

	result := strings.Builder{}

	for i := 0; i < n; i++ {

		count := 1
		for i < n-1 && str[i] == str[i+1] {
			count++
			i++
		}

		result.WriteString(fmt.Sprint(count, string(str[i])))
	}

	return result.String()
}

func (c *Compression) Decode(str string) string {
	runes := []rune(str)

	decode := ""
	count := ""

	for _, r := range runes {
		if unicode.IsLetter(r) {
			n, _ := strconv.Atoi(count)
			for i := 0; i < n; i++ {
				decode += string(r)
			}

			count = ""
		}

		if unicode.IsDigit(r) {
			count += string(r)
		}
	}

	return decode
}

func NewCompression() *Compression {
	return &Compression{}
}
