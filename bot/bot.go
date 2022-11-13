package bot

import (
	"discordbot/ethereumsupport"
	"errors"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Token = os.Getenv("token")
)

var (
	ErrorConnectDiscord = errors.New("connect discord bot failed")
)

type Bot struct {
	Discord   *discordgo.Session
	ChannelId map[string]string
	Client    *ethereumsupport.Provide
}

func NewDiscordBot() (*Bot, error) {
	if Token == "" {
		return nil, ErrorConnectDiscord
	}

	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}

	var bot = new(Bot)
	bot.Discord = discord
	bot.registerChannelId()

	return bot, nil
}

func (b *Bot) registerChannelId() {
	b.ChannelId = make(map[string]string, 10)
	b.ChannelId["faucet"] = "1040150522935640124"
	b.ChannelId["whale-track"] = "1040184323359383562"
	b.ChannelId["defi-track"] = "1040519418985332817"
	b.ChannelId["originization-token"] = "1040185040656662568"
}

func (b *Bot) Ping() {
	b.Discord.AddHandler(func(s *discordgo.Session, r *discordgo.MessageCreate) {
		fmt.Println("Stevel is find!")
	})
}
