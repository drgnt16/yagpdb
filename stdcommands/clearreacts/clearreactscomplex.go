package clearreactscomplex

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"strings"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "ClearReactsComplex",
	Description:  "Clear reacts from a given message not contained in a given list of emojis",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "Message", Type: dcmd.Int, Help: "Message to work on"},
		{Name: "Channel", Type: dcmd.Int, Help: "Message to work on"},
		{Name: "Whitelist", Type: dcmd.String, Help: "Reactions to keep"},
	},
	DefaultEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		//Get Message and Slice
		message, _ := common.BotSession.ChannelMessage(data.Args[1].Value.(int64), data.Args[0].Value.(int64))
		whitelist := strings.Split(data.Args[2].Value.(string), ";")
		//Loop reactions in message
		for _, reaction := range message.Reactions {
			//Check if reaction is in list
			if !contains(whitelist, reaction.Emoji.Name) {
				//Remove reaction
				common.BotSession.MessageReactionRemoveEmoji(message.ChannelID, message.ID, reaction.Emoji.APIName())
			}
		}
		return nil, nil
	},
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
