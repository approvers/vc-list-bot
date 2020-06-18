package command

import(
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func List(session *discordgo.Session, message *discordgo.MessageCreate){
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content == "!list" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}
