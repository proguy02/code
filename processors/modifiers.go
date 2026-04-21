package processors

import (
	"strconv"
	"strings"
)

func Up(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(up)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1])
			word = append(word[:r], word[r+1:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}

func Low(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(low)" && r > 0 {
			word[r-1] = strings.ToLower(word[r-1])
			word = append(word[:r], word[r+1:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}

func Cap(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(cap)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1][:1]) + strings.ToLower(word[r-1][1:])
			word = append(word[:r], word[r+1:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}

func UpN(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(up," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j])
			}
			word = append(word[:r], word[r+2:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}

func LowN(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(low," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToLower(word[j])
			}
			word = append(word[:r], word[r+2:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}

func CapN(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(cap," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j][:1]) + strings.ToLower(word[j][1:])
			}
			word = append(word[:r], word[r+2:]...)
			r--
		}
	}

	return strings.Join(word, " ")
}
