package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	assets "vcListBot/command/assets"

	"github.com/bwmarrin/discordgo"
)

type HelpData struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

func Help(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Content != "!help" {
		return
	}
	data, err := ioutil.ReadFile("./command/assets/help.json")
	if err != nil {
		fmt.Println("error loading file,", err)
	}
	var helpData []HelpData
	err = json.Unmarshal(data, &helpData)
	if err != nil {
		fmt.Println(err)
	}
	helpList := ReturnCommandList(helpData)
	session.ChannelMessageSend(message.ChannelID, helpList)
}

func ReturnCommandList(helpData []HelpData) string {
	helpList := "**= Commands List =**\n```asciidoc\n= Bot Info =\n"
	for _, data := range helpData {
		helpList += assets.PaddingRight(data.Title, 9, " ") + ":: " + data.Explanation + "\n"
	}
	helpList += "```"
	return helpList
}
