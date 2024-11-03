package iteration

import "strings"

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

func RepeatCount(character string, count int) string {
	var builder strings.Builder
	builder.Grow(len(character) * count)
	for i := 0; i < count; i++ {
		builder.WriteString(character)
	}
	return builder.String()
}
