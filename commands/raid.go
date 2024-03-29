package commands

import (
	"log"

	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func RaidModeCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		utils.RaidMode = i.ApplicationCommandData().Options[0].BoolValue()
		log.Println("Raid mode :", utils.RaidMode)
		if utils.RaidMode {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: "<a:Alert:1057105839812530287>Raid Mode<a:Alert:1057105839812530287>",
							Description: "Raid mode has been enabled,\n" +
								"new members will be kicked until raid mode is disabled",
							Color: utils.RED,
							// Timestamp: "<t:" + utils.GetTimestamp() + ">",
						},
					},
				},
			})
		} else {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title:       "Raid Mode",
							Description: "Raid mode has been disabled",
							Color:       utils.GREEN,
							// Timestamp:   "<t:" + utils.GetTimestamp() + ">",
						},
					},
				},
			})
		}
		utils.Log(i.Member.User.Username, i.Member.User.ID, "raid mode command called")
	}
}
