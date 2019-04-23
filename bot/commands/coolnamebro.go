package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/intrntsrfr/meido/bot/service"

	"github.com/bwmarrin/discordgo"
)

var CoolNameBro = Command{
	Name:          "Cool name bro",
	Description:   "Renames attentionseeking nick- or usernames.",
	Triggers:      []string{"m?coolnamebro", "m?cnb"},
	Usage:         "m?coolnamebro my name is shit",
	Category:      Moderation,
	RequiredPerms: discordgo.PermissionManageNicknames,
	Execute: func(args []string, ctx *service.Context) {

		if len(args) < 2 {
			ctx.Send("Please choose a proper name.")
			return
		}

		newName := strings.Join(args[1:], " ")

		memberList := []string{}

		f, err := os.Open("./bot/misc/ranges.json")
		if err != nil {
			return
		}
		defer f.Close()
		ich := charRanges{}

		json.NewDecoder(f).Decode(&ich)

		for _, val := range ctx.Guild.Members {
			if badName(val, &ich) {
				memberList = append(memberList, val.User.ID)
			}
		}

		if len(memberList) < 1 {
			ctx.Send("There is no one rename.")
			return
		} else {
			ctx.Send(fmt.Sprintf("Starting rename of %v user(s).", len(memberList)))
		}

		var successfulRenames, failedRenames int

		for _, val := range memberList {
			err := ctx.Session.GuildMemberNickname(ctx.Guild.ID, val, newName)
			if err != nil {
				failedRenames++
			} else {
				successfulRenames++
			}
		}

		ctx.Send(fmt.Sprintf("Rename finished. Successful: %v. Failed: %v.", successfulRenames, failedRenames))
	},
}

var NiceNameBro = Command{
	Name:          "Nice name bro",
	Description:   "Removes the nickname from users with a provided name.",
	Triggers:      []string{"m?nicenamebro", "m?nnb"},
	Usage:         "m?nicenamebro my name is shit",
	Category:      Moderation,
	RequiredPerms: discordgo.PermissionManageNicknames,
	Execute: func(args []string, ctx *service.Context) {

		if len(args) < 2 {
			ctx.Send("Please choose a name.")
			return
		}

		newName := strings.Join(args[1:], " ")

		memberList := []string{}

		f, err := os.Open("./bot/misc/ranges.json")
		if err != nil {
			return
		}
		defer f.Close()
		ich := charRanges{}

		json.NewDecoder(f).Decode(&ich)

		for _, val := range ctx.Guild.Members {
			if val.Nick != "" {
				if val.Nick == newName {
					if !isRenamed(val, &ich) {
						memberList = append(memberList, val.User.ID)
					}
				}
			}
		}

		if len(memberList) < 1 {
			ctx.Send("There is no one rename.")
			return
		} else {
			ctx.Send(fmt.Sprintf("Starting rename of %v user(s).", len(memberList)))
		}

		var successfulRenames, failedRenames int

		for _, val := range memberList {
			err := ctx.Session.GuildMemberNickname(ctx.Guild.ID, val, "")
			if err != nil {
				failedRenames++
			} else {
				successfulRenames++
			}
		}

		ctx.Send(fmt.Sprintf("Rename finished. Successful: %v. Failed: %v.", successfulRenames, failedRenames))
	},
}

func badName(u *discordgo.Member, ich *charRanges) bool {
	isIllegal := false

	if u.Nick != "" {
		r, _ := utf8.DecodeRuneInString(u.Nick)
		for _, rng := range ich.Ranges {
			isIllegal = rng.Start <= int(r) && int(r) <= rng.Stop
			if isIllegal {
				break
			}
		}
	} else {
		r, _ := utf8.DecodeRuneInString(u.User.Username)
		for _, rng := range ich.Ranges {
			isIllegal = rng.Start <= int(r) && int(r) <= rng.Stop
			if isIllegal {
				break
			}
		}
	}
	return isIllegal
}

func isRenamed(u *discordgo.Member, ich *charRanges) bool {
	isIllegal := false

	r, _ := utf8.DecodeRuneInString(u.User.Username)
	for _, rng := range ich.Ranges {
		isIllegal = rng.Start <= int(r) && int(r) <= rng.Stop
		if isIllegal {
			break
		}
	}
	return isIllegal
}

type charRanges struct {
	Ranges []struct {
		Start int `json:"start"`
		Stop  int `json:"stop"`
	} `json:"ranges"`
}