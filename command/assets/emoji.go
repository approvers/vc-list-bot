package assets

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RandomSelectEmoji(emojis []*discordgo.Emoji) string {
	if len(emojis) < 1 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	selectNum := rand.Intn(len(emojis))
	return FormatEmoji(emojis[selectNum])
}

func FormatEmoji(emojiData *discordgo.Emoji) string {
	return "<:" + emojiData.Name + ":" + emojiData.ID + ">"
}
