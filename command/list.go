package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"vcListBot/command/assets"
)

type VoiceStates struct {
	voiceBotNumber int
	voiceMuteNumber int
}

const textLength = 8

func List(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, guild, err := assets.GetGuildData(session, message)
	if err != nil {
		errMessage := "**ERR: **getting channel or guild```" + err.Error() + "```"
		session.ChannelMessageSend(message.ChannelID, errMessage)
		return
	}
	memberCount := guild.MemberCount
	voiceJoinNumber := len(guild.VoiceStates)
	states, err := GetVoiceStates(guild, session)
	if err != nil {
		errMessage := "**ERR: **getting the user details```" + err.Error() + "```"
		session.ChannelMessageSend(message.ChannelID, errMessage)
		return
	}

	utterance := assets.RandomSelectEmoji(guild.Emojis)
	utterance += " ***限界リスト***"
	utterance += assets.RandomSelectEmoji(guild.Emojis)
	utterance += "```asciidoc\n= 現在の状況 =\n"
	utterance += AllMember(memberCount)
	utterance += InVoiceMembers(voiceJoinNumber, states.voiceBotNumber)
	if voiceJoinNumber != 0 {
		utterance += MuteMembers(states.voiceMuteNumber)
		utterance += VoiceMemberRate(memberCount, voiceJoinNumber)
		utterance += MuteRate(voiceJoinNumber, states.voiceMuteNumber)
	} else {
		utterance = "今は誰もいないよ :pleading_face::sweat_drops: \n" + utterance
	}
	utterance += "```"

	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetVoiceStates(guild *discordgo.Guild, session *discordgo.Session) ( states VoiceStates, err error) {
	var user *discordgo.User
	for _, vs := range guild.VoiceStates {
		user, err = session.User(vs.UserID)
		if err != nil {
			return
		}
		if user.Bot {
			states.voiceBotNumber++
		}
		if vs.SelfMute {
			states.voiceMuteNumber++
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
