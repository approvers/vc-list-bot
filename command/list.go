package command

import (
	"github.com/bwmarrin/discordgo"

	"vcListBot/command/assets"
)

type VoiceStates struct {
	voiceBotNumber  int
	voiceMuteNumber int
}

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
	utterance += " ***"+ guild.Name +" のVC ***"
	utterance += assets.RandomSelectEmoji(guild.Emojis)
	if voiceJoinNumber != 0 {
		utterance += "```asciidoc\n= 現在の状況 =\n"
		utterance += assets.InVoiceMembers(voiceJoinNumber, states.voiceBotNumber)
		utterance += assets.MemberNum(states.voiceMuteNumber, "ミュート人数")
		utterance += assets.VoiceMemberRate(memberCount, voiceJoinNumber)
		utterance += assets.MuteRate(voiceJoinNumber, states.voiceMuteNumber)
		utterance += "```"
	} else {
		utterance += "\n今は誰もいないよ :pleading_face::sweat_drops: \n"
	}

	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetVoiceStates(guild *discordgo.Guild, session *discordgo.Session) (states VoiceStates, err error) {
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
