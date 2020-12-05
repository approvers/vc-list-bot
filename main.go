package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	command "vcListBot/command"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("error loading .env file,", err)
	}
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		fmt.Println("please set ENV: DISCORD_BOT_TOKEN")
		return
	}
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(Apportion)
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	discord.Close()
}

func Apportion(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	if message.Content == "!help" {
		command.Help(session, message)
		return
	}
	if message.Content == "!list" {
		command.List(session, message)
		return
	}
	if message.Content == "!member_num" {
		command.Member(session, message)
		return
	}
	return
}
