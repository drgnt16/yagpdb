package timestamp

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/discordgo"
	"strconv"
	"time"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "timestamp",
	Description:  "Adds a ticker to the screen",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "time", Type: dcmd.String, Help: "The time to display"},
	},
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		inputTime := data.Args[0].Value.(string)
		tm, _ := time.Parse("02 Jan 06 15:04 MST", inputTime)
		if tm.IsZero() || tm.Unix() < 0 {
			return "Invalid time, Please use format: " + time.Now().Format("02 Jan 06 15:04 MST"), nil
		}
		loc, err := time.LoadLocation("UTC")
		if tm.Location().String() == "PST" || tm.Location().String() == "PDT" {
			loc, err = time.LoadLocation("America/Los_Angeles")
		} else if tm.Location().String() == "CEST" || tm.Location().String() == "CET" {
			loc, err = time.LoadLocation("Europe/Brussels")
		} else if tm.Location().String() == "GMT" || tm.Location().String() == "BST" {
			loc, err = time.LoadLocation("Europe/London")
		} else if tm.Location().String() == "EST" || tm.Location().String() == "EDT" {
			loc, err = time.LoadLocation("America/New_York")
		} else {
			loc, err = time.LoadLocation(tm.Location().String())
		}
		if err != nil {
			return "Invalid Timezone, Please ask Quack to fix it", nil
		}
		zone, offset := tm.In(loc).Zone()
		epoch := tm.UTC().Unix() - int64(offset)
		strTime := strconv.FormatInt(epoch, 10)
		result := "Timezone Used: " + zone + "\n" +
			"Stream Ping: `<t:" + strTime + ":t> <t:" + strTime + ":R>` : <t:" + strTime + ":t> <t:" + strTime + ":R>\n" +
			"Relative Time: `<t:" + strTime + ":R>` : <t:" + strTime + ":R>\n" +
			"Absolute Time: `<t:" + strTime + ":F>` : <t:" + strTime + ":F>\n" +
			"Short Date: `<t:" + strTime + ":f>` : <t:" + strTime + ":f>\n" +
			"Long TIme: `<t:" + strTime + ":T>` : <t:" + strTime + ":T>\n" +
			"Short Time: `<t:" + strTime + ":t>` : <t:" + strTime + ":t>\n"
		embed := &discordgo.MessageEmbed{
			Title:       "Timestamp",
			Description: result,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "If you experience issues please try https://r.3v.fi/discord-timestamps/ until resolved",
			},
		}
		messageSend := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{embed},
		}
		common.BotSession.ChannelMessageSendComplex(data.ChannelID, messageSend)
		return nil, nil
	},
}
