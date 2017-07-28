package util

import (
    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

func SafeSendEmbed(session *discordgo.Session, channelID string, embed *discordgo.MessageEmbed) *discordgo.Message {
    val, err := session.ChannelMessageSendEmbed(channelID, embed)

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "channel": channelID,
            "embed": embed,
            "err": err,
        }).Error("Unable to send embed")
    }

    return val
}