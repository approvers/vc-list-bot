package assets

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func PaddingRight(val string, length int, text string) string {
	for ; utf8.RuneCountInString(val + text) <= length; {
		val += text
	}
	t := []rune(text)
	return val + string(t[0:length - utf8.RuneCountInString(val)])
}

func FormatRateNum(rate float64) string {
	return strconv.FormatFloat(math.Round(rate*1000) / 10, 'g', 4, 64)
}
