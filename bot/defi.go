package bot

import (
	"discordbot/ethereumsupport/defi"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) Defi(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "!Change1h" {

	}

	if m.Content == "!Change1d" {
		b.DefiToken = defi.NewDefi()
		b.DefiToken.DefiTrack()
		for {
			select {
			case change1d := <-b.DefiToken.Change1d:
				fmt.Println(change1d)
				_, _ = s.ChannelMessageSend(b.ChannelId["defi-track"], change1d.Name)
			}
		}

	}
}
