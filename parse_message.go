package main

import (
    "github.com/jbillote/chocobot/models"
    "github.com/jbillote/chocobot/modules"

    "strings"

    "github.com/bwmarrin/discordgo"

    // "github.com/Sirupsen/logrus"
)

func ParseMessage(session *discordgo.Session, message *discordgo.MessageCreate,
    config *models.Config) {

    // Make sure bot doesn't attempt to parse its own messages
    if message.Author.ID == botUID {
        return
    }

    // Stop parsing if the message wasn't a command
    if len(message.Content) < 1 || message.Content[0:1] !=
        config.CommandPrefix {
        
        return
    }

    command := strings.Split(message.Content, " ")[0]

    if command[1:len(command)] == "character" {
        modules.FFLogs(session, message, config)
    }
}