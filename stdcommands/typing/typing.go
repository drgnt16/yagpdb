package typing

import (
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/dstate"
	"strconv"
	"time"
)

var Command = &commands.YAGCommand{
	CmdCategory:  commands.CategoryFun,
	Name:         "typing",
	Description:  "Types in chat for given time",
	RequiredArgs: 1,
	Arguments: []*dcmd.ArgDef{
		{Name: "Channel", Type: dcmd.ChannelOrThread, Help: "Channel to type in"},
		{Name: "Seconds", Type: dcmd.String, Help: "Seconds to type for"},
	},
	DefaultEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		channel := data.Args[0].Value.(*dstate.ChannelState)
		seconds, _ := strconv.Atoi(data.Args[1].Value.(string))
		if seconds > 10 {
			seconds = seconds - 10
		}
		//type in channel for given seconds
		for i := 0; i < seconds; i++ {
			common.BotSession.ChannelTyping(channel.ID)
			time.Sleep(time.Second)
		}

		return nil, nil
	},
}
