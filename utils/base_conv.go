package utils

import (
	"strconv"
	"strings"
)

func Base(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(hex)" && r > 0 {
			conv, err := strconv.ParseInt(word[r-1], 16, 64)
			if err != nil {
				return err.Error()
			}
			word[r-1] = strconv.FormatInt(conv, 10)
		}
		if word[r] == "(bin)" && r > 0 {
			conv, err := strconv.ParseInt(word[r-1], 2, 64)
			if err != nil {
				return err.Error()
			}
			word[r-1] = strconv.FormatInt(conv, 10)
		}
	}
	return strings.Join(word, " ")
}
