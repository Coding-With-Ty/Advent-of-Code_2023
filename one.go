package day

import (
	_ "embed"
	"unicode"
)

//go:embed input.txt
var input string

func DayOne() int {
	var tokens []int
	var result int
	scalingFactor := 10
	for _, char := range input {
		if char == '\n' {
			if len(tokens) > 0 {
				if len(tokens) == 1 {
					result = result + tokens[0]*scalingFactor + tokens[0]
				} else {
					result = result + tokens[0]*scalingFactor + tokens[len(tokens)-1]
				}
				tokens = tokens[:0]
			}
		} else if unicode.IsDigit(char) {
			digit := int(char - '0')
			tokens = append(tokens, digit)
		}
	}
	return result
	//fmt.Println(result)
}
