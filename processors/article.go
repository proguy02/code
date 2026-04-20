package processors

import "strings"

func Qoute(sen string) string {
	words := strings.Split(sen, "'")

	for x := 1; x < len(words); x += 2 {
		words[x] = strings.TrimSpace(words[x])
	}
	return strings.Join(words, "'")
}

func Qoute2(sen string) string {
	words := strings.Split(sen, `"`)

	for x := 1; x < len(words); x += 2 {
		words[x] = strings.TrimSpace(words[x])
	}
	return strings.Join(words, `"`)
}

func AnA(sen string) string {
	words := strings.Fields(sen)

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
			char = char[1:]
		}
		if char != " " {
			result = append(result, char)
		}
	}
	return strings.Join(result, " ")
}
