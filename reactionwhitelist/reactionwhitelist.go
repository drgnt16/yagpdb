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
	whitelist := "dougwut;spoonsup;D1;D2;D3;D4;DougDoug;Barn;Dy;Dx;DougShirt;DougShirt2;Pword;alphabetcrew;zcrew;Minecraft;DrinkWater;RawEggRoulette;DougFacingMayo;DougOnTime;Dz;Dw;DoubleOrNothing;Confidence;dougMayoHead;dougFull;dougHappy;dougDoubt;dougBarf;DouglasWredenLivestream;Panic;dougDisbelief;dougPain;dougNom;REAL_DooGdoOg;mayo_on_head_of_Doug;Gun_down;Gun_up;Gun_right;Gun_left;PlushPlush;DeadDead;dougHappyfuntimes;dougThinking;Ban_hammer;dougCute;SpoonGoon;DrugDrug;BananaUp;BananaRight;BananaLeft;BananaDown;Not;Real;Pickle;Bandit;BarnFindingArt;BarnFindingOG;DOOBIE1;DOOBIE2;DougTired;DOOBIE3;DougReallyTired;Angel_Frog;Gun_Screen;BanSoftly;BanHarder;BanHardest;Pepper_hands;Pepper_sleep;Food;DetectiveDoug;catJAM;edfoot;EDDIEWHAT;PixelPepper;Rosa;Birthday_Rosa;RosaMVP;DestroyFrogBadge;FrogBadge;superfrog;doug404;POWPOW;dotheMARIO;PetDoug;dougCry;waluigiSpecialist;EddiePray;trueDoug;dougButter;parkzerSoap;marioPound;plush1;plush2;plush3;plush4;parkzerGun;monkez;eddieCrisps;waluigiNut;henry;simon;parkzerShoot;dougPeek;eddiePing;eddieT;dougWow;dougThong;warOREO;brocLUIGI;DOUGGERS;dougShank;checkThePins;dougdougWave;gamezAngy;pepperUwu;strawberryCouch;frogCouch;dougJAM;dougBAKA;pepperStare;LuL;dougNotBad;waluigiDAB;dougINSIDEOUT;dougHeadache;dougRIVATIVE;dougINTEGRAL;dougScared;WONTON;dougPog;THEDARGEBARGE;THECOOLDARGEBARGE;dougFU;pepperLUL;swooper;swooperBack;wooperBack;wooper;dougLeprechaun;DougBriish;destroypeach;dougSONIC;dougConfidence;pepperGun;monsterLUL;pepperCute;DougNotBad2;BarnHub;OldDougLogo;Goudas;Gouda;DougShocked;EddiePls;dougShrug;parkzerGunRight;dougTroll;dunsparceCapture;dougPOWER;HenryCute;SimonCute;toe4;toe5;toe3;toe1;toe2;dougSimoned;dougButtsimon;dougKonoDa;ryjo6;eddoug;dougIntimidate;dougBeans;ddarkwalnut;ddarkwalnutd1;ddarkwalnutd2;PridePepper;PepperPride;dougTeaching;dougLICK;dougdougBackpack;DougConfused;dougnotagain;dougLuL;dougMEGALUL;pepperarrive;pepperleave;dougWOW;waPepper;petGamez;dougKEKW;e_;PauseDoug;dougSMOHUD;dougAHHHHH;PetQs;peach;ASU;bun;dougBunnyHat;bunluigi;halt;charredpepper;dougLaugh;D_;petRosa;dougdougF;pet_henry;pet_simon;ClaudeS;GrandDoug;newcamera;goose;toad;dougdougToad;dougChamp;dougdougA2U;dougdougClaude;dougdougPump;dougdougBunny;dougdougAnime;dougdougHeart;dougdougShank;dougdougRigged;dougdougPain;dougdougPog;dougdougDoubt;dougdougUwU;dougdoug404;dougdougPANIC;dougdougLOUDER;dougdougPPF;dougdougNerf;dougdougWut;dougdougChamp;dougdougWuv;dougdougACrew;dougdougZCrew;dougdougLUL;dougdougMagicHat;PLANNED;typing;dougdougCheer;dougdougConfused;dougdougTears;dougdougHmm;dougdougMonka;dougdougNotes;dougdougSalute;dougdougShocked;dougcostumechange;A2U;DougBonk;BonkHardest;SheriffDougGun;acrew;DUGGERS;boop;dougdougMurica;dougdougYell;dougdougPause;dougdougShock;dougdougConfidence;dougdougHuh;dougdougThink;dougdougASU;dougdougBrush;dougdougConfidence;dougdougFU;dougdougGasm;dougdougHappy;dougdougKEKW;dougdougNotLikeThis;dougdougPain;dougdougPause;dougdougPog;dougdougScared;dougdougSimoned;dougdougThink;dougdougWut;👍;👎;👌;❤️;🖤;🧡;💙;💚;💜;💛;🤎;🤍;😃;😂;☺️;😍;🥰;♥️;💖;😁;😎;🎂;🥳;🫂;👏;💕;💞;🙂;😀;💀;💯;👀;😻;😊;🤔;😰;😥;🙁;😦;😭;🇫;🖕;🎉;👆;🆗;☝️;💗;🙏;🐰"
	emoji, cID, gID, uID, mID, _ := getReactionDetails(evt)
	//If in ping channel
	if uID != common.BotUser.ID && (cID == 731407385624838197 || cID == 567144073857859609 || cID == 880127379119415306) {
		//And not in whitelist
		if !contains(strings.Split(whitelist, ";"), emoji.Name) {
			//Remove reaction
			common.BotSession.MessageReactionRemoveEmoji(cID, mID, emoji.APIName())
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
				if !contains(whitelist, reaction.Emoji.Name) {
					//Remove reaction
					common.BotSession.MessageReactionRemoveEmoji(message.ChannelID, message.ID, reaction.Emoji.APIName())
				}
			}
			return nil, nil
		},
	})
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
