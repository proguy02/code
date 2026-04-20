package processors

import (
	"strconv"
	"strings"
)

func Hex(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(hex)" && r > 0 {
			conv, err := strconv.ParseInt(word[r-1], 16, 64)
			if err != nil {
				return err.Error()
			}
			word[r-1] = strconv.FormatInt(conv, 10)
			word = append(word[:r], word[r+1:]...)
		}
	}
	return strings.Join(word, " ")
}

func Bin(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(bin)" && r > 0 {
			conv, err := strconv.ParseInt(word[r-1], 2, 64)
			if err != nil {
				return err.Error()
			}
			word[r-1] = strconv.FormatInt(conv, 10)
			word = append(word[:r], word[r+1:]...)
		}
	}
	return strings.Join(word, " ")
}
