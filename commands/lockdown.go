package commands

import (
	"github.com/intrntsrfr/meido/service"

	"github.com/bwmarrin/discordgo"
)

var Lockdown = Command{
	Name:          "lockdown",
	Description:   "Locks down the current channel, denying the everyonerole send message perms.",
	Triggers:      []string{"m?lockdown"},
	Usage:         "m?lockdown",
	RequiredPerms: discordgo.PermissionManageRoles,
	Execute: func(args []string, ctx *service.Context) {
		var erole *discordgo.Role

		for _, val := range ctx.Guild.Roles {
			if val.ID == ctx.Guild.ID {
				erole = val
			}
		}

		var eperms *discordgo.PermissionOverwrite

		for _, val := range ctx.Channel.PermissionOverwrites {
			if val.ID == erole.ID {
				eperms = val
			}
		}

		if erole == nil || eperms == nil {
			return
		}

		if eperms.Allow&discordgo.PermissionSendMessages == 0 && eperms.Deny&discordgo.PermissionSendMessages == 0 {
			// DEFAULT
			err := ctx.Session.ChannelPermissionSet(
				ctx.Channel.ID,
				erole.ID,
				"role",
				eperms.Allow,
				eperms.Deny+discordgo.PermissionSendMessages,
			)
			if err != nil {
				ctx.Send("Could not lock channel.")
				return
			}
			ctx.Send("Channel locked.")
		} else if eperms.Allow&discordgo.PermissionSendMessages != 0 && eperms.Deny&discordgo.PermissionSendMessages == 0 {
			// IS ALLOWED
			err := ctx.Session.ChannelPermissionSet(
				ctx.Channel.ID,
				erole.ID,
				"role",
				eperms.Allow-discordgo.PermissionSendMessages,
				eperms.Deny+discordgo.PermissionSendMessages,
			)
			if err != nil {
				ctx.Send("Could not lock channel.")
				return
			}
			ctx.Send("Channel locked")
		} else if eperms.Allow&discordgo.PermissionSendMessages == 0 && eperms.Deny&discordgo.PermissionSendMessages != 0 {
			// IS DENIED
			ctx.Send("Channel already locked")
		}
	},
}
