package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"vcListBot/command/assets"

	"github.com/bwmarrin/discordgo"
)

type HelpData struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

var helpData []HelpData

func init() {
	data, err := ioutil.ReadFile("./command/assets/help.json")
	if err != nil {
		fmt.Println("error loading file,", err)
		os.Exit(1)
	}
	err = json.Unmarshal(data, &helpData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Help(session *discordgo.Session, message *discordgo.MessageCreate) {
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
