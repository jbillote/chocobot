package util

import (
    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

func SafeSendMessage(session *discordgo.Session, channelID string, message string) *discordgo.Message {
    val, err := session.ChannelMessageSend(channelID, message)

    if err != nil {
        logrus.WithFields(logrus.Fields{
            "channel": channelID,
            "message": message,
            "err": err,
        }).Error("Unable to send message")
    }

    return val
}