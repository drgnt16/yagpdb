package redirect

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/botlabs-gg/yagpdb/v2/lib/dstate"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "Redirect",
	Description:  "Redirects a message to a specified channel",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "channel", Type: dcmd.Channel},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		//Load channel and message
		channel := data.Args[0].Value.(*dstate.ChannelState)
		msgIN := data.TraditionalTriggerData.Message.ReferencedMessage
		//Re-send message to specified channel
		msgOUT, err := common.BotSession.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
			Content: msgIN.Content,
			Embeds:  msgIN.Embeds,
			AllowedMentions: discordgo.AllowedMentions{
				Parse: []discordgo.AllowedMentionType{discordgo.AllowedMentionTypeUsers},
			},
		})
		//check for errors
		if err != nil {
			return nil, err
		}
		//Notify of success
		msgChannel, _ := common.BotSession.Channel(msgOUT.ChannelID)
		return "Sent to " + msgChannel.Mention(), nil
	},
}
