package commands

import (
	"fmt"
	"strings"

	"github.com/intrntsrfr/meido/bot/service"

	"github.com/bwmarrin/discordgo"
)

var Invite = Command{
	Name:          "Invite",
	Description:   "Sends bot invite link and support server invite.",
	Triggers:      []string{"m?invite", "m?inv"},
	Usage:         "m?invite",
	Category:      Utility,
	RequiredPerms: discordgo.PermissionSendMessages,
	Execute: func(args []string, ctx *service.Context) {
		botLink := "https://discordapp.com/oauth2/authorize?client_id=394162399348785152&scope=bot"
		serverLink := "https://discord.gg/KgMEGK3"
		ctx.Send(fmt.Sprintf("Invite me to your server: %v\nSupport server: %v", botLink, serverLink))
	},
}

var Feedback = Command{
	Name:          "Feedback",
	Description:   "Sends your very nice and helpful feedback to the Meido Café.",
	Triggers:      []string{"m?feedback", "m?fb"},
	Usage:         "m?feedback wow what a really COOL and NICE bot that works flawlessly",
	Category:      Utility,
	RequiredPerms: discordgo.PermissionSendMessages,
	Execute: func(args []string, ctx *service.Context) {
		if len(args) > 1 {
			text := fmt.Sprintf("Message from %v - %v (%v) from channel %v (%v) in server %v (%v)\n", ctx.User.Mention(), ctx.User.String(), ctx.User.ID, ctx.Channel.Name, ctx.Channel.ID, ctx.Guild.Name, ctx.Guild.ID)
			text += fmt.Sprintf("`%v`", (strings.Join(args[1:], " ")))
			_, err := ctx.Session.ChannelMessageSend("533009623188242443", text)
			if err != nil {
				fmt.Println(err)
			}
			ctx.Send("Feedback left")
		}
	},
}