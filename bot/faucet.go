package bot

import (
	"discordbot/ethereumsupport"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Index = 7
)

func (b *Bot) Faucet(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "!faucet") && m.ChannelID == b.ChannelId["faucet"] {
		receiver := strings.TrimSpace(m.Content[Index:])
		result, err := b.Client.SendFaucet(receiver)
		if err != nil {
			fmt.Println(err)
			return
		}
		var data ethereumsupport.Faucet
		json.Unmarshal(result, &data)

		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Transaction:%s,   Status: Send faucet successed!", data.Tx))
	}

}
