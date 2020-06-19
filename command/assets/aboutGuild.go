package assets

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetGuildData(session *discordgo.Session, message *discordgo.MessageCreate) (channel *discordgo.Channel, guild *discordgo.Guild, err error) {
	channel, err = session.Channel(message.ChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}
	guild, err = session.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
