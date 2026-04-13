package utils

import (
	"strconv"
	"strings"
)

func String(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(up)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1])
		}
		if word[r] == "(low)" && r > 0 {
			word[r-1] = strings.ToLower(word[r-1])
		}
		if word[r] == "(cap)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1][:1]) + strings.ToLower(word[r-1][1:])
		}
		if word[r] == "(up," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j])
			}
		}
		if word[r] == "(low," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToLower(word[j])
			}
		}
		if word[r] == "(cap," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j][:1]) + strings.ToUpper(word[j][1:])
			}
		}
	}
	return strings.Join(word, " ")
}
