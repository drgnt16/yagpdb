package ticker

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"github.com/botlabs-gg/yagpdb/v2/lib/dstate"
	"strings"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "ticker",
	Description:  "Adds a ticker to the screen",
	RequiredArgs: 1,
	ArgSwitches: []*dcmd.ArgDef{
		{Name: "Action", Type: dcmd.String, Help: "add, update or remove"},
		{Name: "Name", Type: dcmd.String},
		{Name: "Channel", Type: dcmd.Channel},
	},
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		guildID := data.GuildData.GS.ID
		//Check what action we are doing
		switch strings.ToLower(data.Switch("Action").Value.(string)) {
		case "add":
			c, _ := common.BotSession.GuildChannelCreate(guildID, data.Switch("Name").Value.(string), 2)
			common.BotSession.ChannelPermissionSet(c.ID, guildID, 0, 0, discordgo.PermissionVoiceConnect)
			return c.Mention() + " Added", nil
		case "update":
			channel := data.Switch("Channel").Value.(*dstate.ChannelState)
			if channel.Type == 2 && channel.ParentID == 0 {
				o, _ := common.BotSession.Channel(channel.ID)
				c, _ := common.BotSession.ChannelEdit(channel.ID, data.Switch("Name").Value.(string))
				return o.Name + " Renamed to " + c.Mention(), nil
			}
			return "Invalid channel, You may only edit channels created by the ticker command", nil
		case "remove":
			channel := data.Switch("Channel").Value.(*dstate.ChannelState)
			if channel.Type == 2 && channel.ParentID == 0 {
				c, _ := common.BotSession.ChannelDelete(data.Switch("Channel").Value.(*dstate.ChannelState).ID)
				return c.Name + " Removed", nil
			}
			return "Invalid channel, You may only remove channels created by the ticker command", nil
		default:
			return "Invalid Command: Please use Add, Update or Remove as the action", nil
		}
		return nil, nil
	},
}
