package assets

import(
	"github.com/bwmarrin/discordgo"
)

func GetGuildData(session *discordgo.Session, message *discordgo.MessageCreate)(channel *discordgo.Channel, guild *discordgo.Guild){
	channel, err := session.Channel(message.ChannelID);
	if err != nil {
		return
	}
	guild, err = session.State.Guild(channel.GuildID)
	if err != nil {
		return
	}
	return channel, guild
}