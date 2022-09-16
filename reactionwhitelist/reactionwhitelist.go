package reactionwhitelist

import (
	"context"
	"github.com/botlabs-gg/yagpdb/v2/bot/eventsystem"
	"github.com/botlabs-gg/yagpdb/v2/common/scheduledevents2"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"strings"
	"time"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)

var (
	logger = common.GetPluginLogger(&Plugin{})
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Reaction Whitelist",
		SysName:  "reactionwhitelist",
		Category: common.PluginCategoryMisc,
	}
}

func RegisterPlugin() {
	p := &Plugin{}
	common.RegisterPlugin(p)
}

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLastLegacy(p, handleReactionAdd, eventsystem.EventMessageReactionAdd)
}

func handleReactionAdd(evt *eventsystem.EventData) {
	whitelist := "👍;👎;👌;❤️;🖤;🧡;💙;💚;💜;💛;🤎;🤍;😃;😂;☺️;😍;🥰;♥️;💖;😁;😎;🎂;🥳;🫂;👏;💕;💞;🙂;😀;💀;💯;👀;😻;😊;🤔;😰;😥;🙁;😦;😭;🇫;🖕;🎉;👆;🆗;☝️;💗;🙏;🐰;🔼;⬇️"
	emoji, cID, gID, uID, mID, _ := getReactionDetails(evt)
	//Allow numbers in Bulletin-Board
	if cID == 880127379119415306 {
		whitelist = whitelist + ";1⃣;2⃣;3⃣;4⃣;5⃣;6⃣;7⃣;8⃣;9⃣;0⃣⃣"
	}
	//If in ping channel
	if uID != common.BotUser.ID && (cID == 731407385624838197 || cID == 567144073857859609 || cID == 880127379119415306) {
		//And not in whitelist or guild emotes
		if !stringcontains(strings.Split(whitelist, ";"), emoji.Name) && !emojicontains(evt.GS.Emojis, emoji) {
			//Remove reaction
			logger.Info("Whitelist: " + whitelist)
			logger.Info("Emote: " + emoji.Name)
			logger.Info("API: " + emoji.APIName())
			common.BotSession.MessageReactionRemove(cID, mID, emoji.APIName(), uID)
			//add NoReaction group to user who reacted
			common.BotSession.GuildMemberRoleAdd(gID, uID, 988901586669551687)
			scheduledevents2.ScheduleRemoveRole(context.Background(), gID, uID, 988901586669551687, time.Now().Add(time.Hour*72))
		}
	}

}
func getReactionDetails(evt *eventsystem.EventData) (emoji *discordgo.Emoji, cID, gID, uID, mID int64, add bool) {
	if evt.Type == eventsystem.EventMessageReactionAdd {
		ra := evt.MessageReactionAdd()
		cID = ra.ChannelID
		uID = ra.UserID
		gID = ra.GuildID
		mID = ra.MessageID
		emoji = &ra.Emoji
		add = true
	} else {
		rr := evt.MessageReactionRemove()
		cID = rr.ChannelID
		uID = rr.UserID
		gID = rr.GuildID
		mID = rr.MessageID
		emoji = &rr.Emoji
	}

	return
}

var _ commands.CommandProvider = (*Plugin)(nil)

func (p *Plugin) AddCommands() {
	commands.AddRootCommands(p, &commands.YAGCommand{
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
				if !stringcontains(whitelist, reaction.Emoji.Name) {
					//Remove reaction
					common.BotSession.MessageReactionRemoveEmoji(message.ChannelID, message.ID, reaction.Emoji.APIName())
				}
			}
			return nil, nil
		},
	})
}
func stringcontains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func emojicontains(s []discordgo.Emoji, e *discordgo.Emoji) bool {
	for _, a := range s {
		if a.APIName() == e.APIName() {
			return true
		}
	}
	return false
}
