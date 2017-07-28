package main

import (
    "github.com/jbillote/chocobot/models"

    "encoding/json"
    "flag"
    "os"
    "os/signal"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

var (
    session *discordgo.Session
    config  models.Config
    botUID  string
)

func onReady(session *discordgo.Session, e *discordgo.Ready) {
    logrus.Info("Choco Bot started")
    session.UpdateStatus(0, "Kweh!")
    botUID = e.User.ID
}

func onMessageCreate(session *discordgo.Session,
    message *discordgo.MessageCreate) {

    // Parse messages in their own thread to prevent stalling due to API calls
    go ParseMessage(session, message, &config)
}

func main() {
    args := flag.String("config", "./config.json", "Configuration file")
    flag.Parse()

    // Try to open config file
    configFile, err := os.Open(*args)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "path": *args,
            "err": err,
        }).Fatal("Unable to open config")
        os.Exit(-1)
    }

    decoder := json.NewDecoder(configFile)
    err = decoder.Decode(&config)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Unable to parse config file")
        os.Exit(-1)
    }

    // Create discord session
    logrus.Info("Starting Discord session...")
    session, err = discordgo.New(config.BotToken)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Failed to create Discord session")
        os.Exit(-1)
    }

    // Add ready and message create handlers
    session.AddHandler(onReady)
    session.AddHandler(onMessageCreate)

    // Open Discord websocket connection
    err = session.Open()
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Fatal("Failed to create Discord websocket connection")
        os.Exit(-1)
    }

    // Wait for a signal to quit
    s := make(chan os.Signal, 1)
    signal.Notify(s, os.Interrupt, os.Kill)
    <-s

    return
}