package command

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	assets "vcListBot/command/assets"
)

const textLength = 8

func List(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content != "!list" {
		return
	}
	_, guild, err := assets.GetGuildData(session, message)
	if err != nil {
		fmt.Println("error getting channel or guild,", err)
		return
	}
	memberCount := guild.MemberCount
	voiceJoinNumber, voiceMuteNumber := GetVoiceStates(guild)
	utterance := assets.RandomSelectEmoji(guild.Emojis) + " ***限界リスト***" + assets.RandomSelectEmoji(guild.Emojis)
	utterance += "```asciidoc\n= 現在の状況 =\n"
	utterance += AllMember(memberCount) + InVoiceMembers(voiceJoinNumber)
	if voiceJoinNumber != 0 {
		utterance += MuteMembers(voiceMuteNumber) + VoiceMemberRate(memberCount, voiceJoinNumber) + MuteRate(voiceJoinNumber, voiceMuteNumber)
	} else {
		utterance = "今は誰もいないよ :pleading_face::sweat_drops: \n" + utterance
	}
	utterance += "```"
	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetVoiceStates(guild *discordgo.Guild) (voiceJoinNumber int, voiceMuteNumber int) {
	voiceJoinNumber = len(guild.VoiceStates)
	for _, vs := range guild.VoiceStates {
		if vs.SelfMute {
			voiceMuteNumber++
		}
	}
	return
}

func AllMember(members int) string {
	return assets.PaddingRight("鯖人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func InVoiceMembers(members int) string {
	return assets.PaddingRight("通話人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func MuteMembers(members int) string {
	return assets.PaddingRight("ミュート人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func VoiceMemberRate(memberCount int, voiceJoinNumber int) string {
	rate := float64(voiceJoinNumber) / float64(memberCount)
	return assets.PaddingRight("通話率", textLength, "　") + ":: " + assets.FormatRateNum(rate) + " %\n"
}

func MuteRate(voiceJoinNumber int, voiceMuteNumber int) string {
	rate := float64(voiceMuteNumber) / float64(voiceJoinNumber)
	return assets.PaddingRight("ミュート率", textLength, "　") + ":: " + assets.FormatRateNum(rate) + " %\n"
}
