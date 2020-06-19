package assets

import (
	"unicode/utf8"
	"strconv"
	"math"
)

func PaddingRight(val string, length int, text string) string {
	for i:= utf8.RuneCountInString(val); i < length; i++ {
		val += text
	}
	return val
}

func FormatRateNum(rate float64) float64 {
	return math.Round(rate * 1000) / 10
}

func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'g', 4, 64)
}
