package bot

import (
	"fmt"
	"golang-discord-bot/config"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	// new bot session created
	goBot, err := discordgo.New("Bot " + config.Token)

	//error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
