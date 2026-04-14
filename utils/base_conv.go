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
			word = append(word[:r], word[r+1:]...)
		}
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

func String(str string) string {
	word := strings.Fields(str)

	for r := 0; r < len(word); r++ {
		if word[r] == "(up)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1])
			word = append(word[:r], word[r+1:]...)
		}
		if word[r] == "(low)" && r > 0 {
			word[r-1] = strings.ToLower(word[r-1])
			word = append(word[:r], word[r+1:]...)
		}
		if word[r] == "(cap)" && r > 0 {
			word[r-1] = strings.ToUpper(word[r-1][:1]) + strings.ToLower(word[r-1][1:])
			word = append(word[:r], word[r+1:]...)
		}
		if word[r] == "(up," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j])
			}
			word = append(word[:r], word[r+2:]...)
		}
		if word[r] == "(low," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToLower(word[j])
			}
			word = append(word[:r], word[r+2:]...)
		}
		if word[r] == "(cap," && r > 0 {
			next := word[r+1]
			next = strings.Trim(next, ")")
			conv, _ := strconv.Atoi(next)
			for j := r - conv; j < r; j++ {
				word[j] = strings.ToUpper(word[j][:1]) + strings.ToUpper(word[j][1:])
			}
			word = append(word[:r], word[r+2:]...)
		}
	}
	return strings.Join(word, " ")
}

func Qoute(sen string) string {
	words := strings.Split(sen, "'")

	for x := 1; x < len(words); x += 2 {
		words[x] = strings.TrimSpace(words[x])
	}
	return strings.Join(words, "'")
}

func AnA(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		vowels := strings.ContainsAny("aeiouhAEIOUH", string(words[i+1][:1]))
		if words[i] == "a" && vowels {
			words[i] = "an"
		} else if words[i] == "A" && vowels {
			words[i] = "An"
		} else if words[i] == "an" && !vowels {
			words[i] = "a"
		} else if words[i] == "An" && !vowels {
			words[i] = "A"
		}
	}
	return strings.Join(words, " ")
}

func Puncts(s string) string {
	words := strings.Fields(s)
	result := []string{}
	for _, char := range words {
		for len(char) > 0 && strings.ContainsAny(char[:1], ".,!?,") {
			if len(result) > 0 {
				result[len(result)-1] += char[:1]
			}
			char = char[:1]
		}
		if char != " " {
			result = append(result, char)
		}
	}
	return strings.Join(result, " ")
}
