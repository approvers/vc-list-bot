package assets

import (
	"math"
	"strconv"
	"unicode/utf8"
)

const textLength = 8

func MemberNum(members int, text string) string {
	return PaddingRight(text, textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func InVoiceMembers(members int, bot int) string {
	return PaddingRight("通話人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人（bot " + strconv.Itoa(bot) + "人）\n"
}

func VoiceMemberRate(memberCount int, voiceJoinNumber int) string {
	rate := float64(voiceJoinNumber) / float64(memberCount)
	return PaddingRight("通話率", textLength, "　") + ":: " + FormatRateNum(rate) + " %\n"
}

func MuteRate(voiceJoinNumber int, voiceMuteNumber int) string {
	rate := float64(voiceMuteNumber) / float64(voiceJoinNumber)
	return PaddingRight("ミュート率", textLength, "　") + ":: " + FormatRateNum(rate) + " %\n"
}

func PaddingRight(val string, length int, text string) string {
	for utf8.RuneCountInString(val+text) <= length {
		val += text
	}
	t := []rune(text)
	return val + string(t[0:length-utf8.RuneCountInString(val)])
}

func FormatRateNum(rate float64) string {
	return strconv.FormatFloat(math.Round(rate*1000)/10, 'g', 4, 64)
}
