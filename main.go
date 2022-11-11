package main

import (
	"discordbot/bot"
	"discordbot/ethereumsupport"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartDiscordServer() {
	// ethereum client
	p, err := ethereumsupport.Provider()
	if err != nil {
		fmt.Println(err)
	}

	dc, err := bot.NewDiscordBot()
	if err != nil {
		fmt.Println(err)
	}

	dc.Client = p

	dc.Discord.Identify.Intents = discordgo.IntentsGuildMessages

	dc.Discord.AddHandler(dc.Faucet)
	dc.Discord.AddHandler(dc.Defi)
	err = dc.Discord.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
		os.Interrupt,
	)
	<-sc
	dc.Discord.Close()
}

func main() {
	StartDiscordServer()

	//bot.Table()
}
