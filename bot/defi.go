package bot

import (
	"discordbot/ethereumsupport/defi"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) Defi(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Listen(b, s)

	if m.Content == "!change1h" {
		var tokens []string
		for _, token := range defi.DefiTrack().Change1d {
			tokens = append(tokens, token.Name)

		}
		fmt.Println(strings.Join(tokens, ","))
		_, _ = s.ChannelMessageSend(b.ChannelId["defi-track"], fmt.Sprintf("您好，主人。1小时10个百分比涨幅币种有：%s", strings.Join(tokens, ",")))
	}

	if m.Content == "!change1d" {
		var tokens []string
		for _, token := range defi.DefiTrack().Change1d {
			tokens = append(tokens, token.Name)

		}
		fmt.Println(strings.Join(tokens, ","))
		_, _ = s.ChannelMessageSend(b.ChannelId["defi-track"], fmt.Sprintf("您好，主人。1天10个百分比涨幅币种有：%s", strings.Join(tokens, ",")))
	}

	if m.Content == "!change7d" {
		var tokens []string
		for _, token := range defi.DefiTrack().Change7d {
			tokens = append(tokens, token.Name)

		}
		fmt.Println(strings.Join(tokens, ","))
		_, _ = s.ChannelMessageSend(b.ChannelId["defi-track"], fmt.Sprintf("您好，主人。7天10个百分比币种有：%s", strings.Join(tokens, ",")))
	}
}

func Listen(b *Bot, s *discordgo.Session) {
	c := NewTimerTask()
	c.AddTask(Second)
	c.AddJob(Second, func() {
		var tokens []string
		for _, token := range defi.DefiTrack().Change1d {
			tokens = append(tokens, token.Name)
		}
		_, _ = s.ChannelMessageSend(b.ChannelId["defi-track"], fmt.Sprintf("自动警报！！！7天10个百分比币种有：%s", strings.Join(tokens, ",")))
	})

	c.Start()

}
