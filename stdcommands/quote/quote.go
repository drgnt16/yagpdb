package quote

import (
	"strconv"
	"strings"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "quote",
	Description:  "Quotes a message",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "URL", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		//Load message data
		s := strings.Split(data.Args[0].Value.(string), "/")
		if s[2] == "discord.com" {
			channelID, err := strconv.ParseInt(s[5], 10, 64)
			messageID, err := strconv.ParseInt(s[6], 10, 64)
			if err != nil {
				return nil, err
			}
			//Get message
			msgIN, err := common.BotSession.ChannelMessage(channelID, messageID)
			//Get attachments
			msgAttach := ""
			for i, a := range msgIN.Attachments {
				if i != 0 {
					msgAttach += " : " + a.URL
				} else {
					msgAttach += a.URL
				}
			}
			if err != nil {
				return nil, err
			}
			//Check for embeds, if so just redirect
			print(msgIN.Embeds)
			print(len(msgIN.Embeds))
			if len(msgIN.Embeds) > 0 {
				_, err = common.BotSession.ChannelMessageSendComplex(data.ChannelID, &discordgo.MessageSend{
					Content: msgIN.Content + "\n" + msgAttach,
					Embeds:  msgIN.Embeds,
					AllowedMentions: discordgo.AllowedMentions{
						Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
					},
				})
				if err != nil {
					return nil, err
				}
				return nil, nil
			}
			//If no embed create one to house message
			msgChannel, _ := common.BotSession.Channel(msgIN.ChannelID)
			embed := &discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name:    msgIN.Author.Username,
					IconURL: msgIN.Author.AvatarURL("64"),
				},
				URL:         data.Args[0].Value.(string),
				Title:       "Message in #" + msgChannel.Name,
				Description: msgIN.Content,
				Fields: []*discordgo.MessageEmbedField{{
					Name:  "Attachments",
					Value: msgAttach,
				}},
				Color:     1,
				Timestamp: string(msgIN.Timestamp),
			}
			_, err = common.BotSession.ChannelMessageSendEmbed(data.ChannelID, embed)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		//URL is invalid
		return "Invalid URL", nil
	},
}
