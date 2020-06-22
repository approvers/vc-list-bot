package command

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"vcListBot/command/assets"
)

const textLength = 8

func List(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, guild, err := assets.GetGuildData(session, message)
	if err != nil {
		fmt.Println("error getting channel or guild,", err)
		return
	}
	memberCount := guild.MemberCount
	voiceJoinNumber := len(guild.VoiceStates)
	voiceBotNumber, voiceMuteNumber := GetVoiceStates(guild, session)

	utterance := assets.RandomSelectEmoji(guild.Emojis)
	utterance += " ***限界リスト***"
	utterance += assets.RandomSelectEmoji(guild.Emojis)
	utterance += "```asciidoc\n= 現在の状況 =\n"
	utterance += AllMember(memberCount)
	utterance += InVoiceMembers(voiceJoinNumber, voiceBotNumber)
	if voiceJoinNumber != 0 {
		utterance += MuteMembers(voiceMuteNumber)
		utterance += VoiceMemberRate(memberCount, voiceJoinNumber)
		utterance += MuteRate(voiceJoinNumber, voiceMuteNumber)
	} else {
		utterance = "今は誰もいないよ :pleading_face::sweat_drops: \n" + utterance
	}
	utterance += "```"

	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetVoiceStates(guild *discordgo.Guild, session *discordgo.Session) (voiceBotNumber int, voiceMuteNumber int) {
	for _, vs := range guild.VoiceStates {
		user,err := session.User(vs.UserID)
		if err != nil {
			return
		}
		if user.Bot {
			voiceBotNumber ++
		}
		if vs.SelfMute {
			voiceMuteNumber++
		}
	}
	return
}

func AllMember(members int) string {
	return assets.PaddingRight("鯖人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人\n"
}

func InVoiceMembers(members int, bot int) string {
	return assets.PaddingRight("通話人数", textLength, "　") + ":: " + strconv.Itoa(members) + " 人（bot " + strconv.Itoa(bot) + "人）\n"
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
