package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
)

type TransferInfo struct {
	Receivers []common.Address `json:"receivers"`
}

func (b *Bot) Transfer(s *discordgo.Session, m *discordgo.MessageCreate) {
	//1041241584949264384
	dmChannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	fmt.Println(m.Content)
	_, _ = s.ChannelMessageSend(dmChannel.ID, "Pong!")

}
