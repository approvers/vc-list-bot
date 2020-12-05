package command

import (
	"github.com/bwmarrin/discordgo"

	"vcListBot/command/assets"
)

func Member(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, guild, err := assets.GetGuildData(session, message)
	if err != nil {
		errMessage := "**ERR: **getting channel or guild```" + err.Error() + "```"
		session.ChannelMessageSend(message.ChannelID, errMessage)
		return
	}
	memberCount := guild.MemberCount
	botNumber := GetGuildStates(guild, session)

	utterance := assets.RandomSelectEmoji(guild.Emojis)
	utterance += " ***"+ guild.Name +" の人数 ***"
	utterance += assets.RandomSelectEmoji(guild.Emojis)
	utterance += "```asciidoc\n= 現在の状況 =\n"
	utterance += assets.MemberNum(memberCount - botNumber, "人間")
	utterance += assets.MemberNum(botNumber, "ボット")
	utterance += assets.MemberNum(memberCount, "合計人数")
	utterance += "```"

	session.ChannelMessageSend(message.ChannelID, utterance)
}

func GetGuildStates(guild *discordgo.Guild, session *discordgo.Session) (botNumber int) {
	for _, member := range guild.Members {
		if member.User.Bot {
			botNumber++
		}
	}
	return
}
