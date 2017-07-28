package modules

import (
    "github.com/jbillote/chocobot/constants"
    "github.com/jbillote/chocobot/models"
    "github.com/jbillote/chocobot/util"

    "encoding/json"
    "net/http"
    "strconv"
    "strings"

    "github.com/bwmarrin/discordgo"

    "github.com/Sirupsen/logrus"
)

var classes map[int]string

func FFLogs(session *discordgo.Session, message *discordgo.MessageCreate, config *models.Config) error {
    // Init classes map if necessary
    if len(classes) == 0 {
        err := initClasses(config)
        if err != nil {
            logrus.WithFields(logrus.Fields{
                "err": err,
            }).Error(err)
        }
    }
    p := strings.Split(message.Content, " ")

    if len(p) < 5 {
        logrus.Error("Not enough arguments")
        return nil
    }

    firstName := p[1]
    lastName := p[2]
    server := p[3]
    region := p[4]

    requestURL := constants.FFLOGS_API_ENDPOINT + constants.FFLOGS_API_CHARACTER + "/" + firstName + "%20" + lastName + "/" + server + "/" + region + "?api_key=" + config.FFLogsApiKey

    r, err := http.Get(requestURL)
    if err != nil {
        // TODO: Make more verbose
        logrus.Error(err)
        return err
    }

    var encounters []models.FFLogsEncounter
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&encounters)
    if err != nil {
        // TODO: Make more verbose
        logrus.Error(err)
        return err
    }

    // Parse encounters
    var o1sEncounters []models.FFLogsEncounter
    var o2sEncounters []models.FFLogsEncounter
    var o3sEncounters []models.FFLogsEncounter
    var exdeathEncounters []models.FFLogsEncounter
    var neoExdeathEncounters []models.FFLogsEncounter
    for _, e := range encounters {
        // Use a switch to make future revisions look cleaner
        switch e.EncounterID {
        case 42:
            // O1S
            o1sEncounters = append(o1sEncounters, e)
            break
        case 43:
            // O2S
            o2sEncounters = append(o2sEncounters, e)
            break
        case 44:
            o3sEncounters = append(o3sEncounters, e)
            break
        case 45:
            exdeathEncounters = append(exdeathEncounters, e)
            break
        case 46:
            neoExdeathEncounters = append(neoExdeathEncounters, e)
            break
        default:
            break
        }
    }

    if len(o1sEncounters) == 0 {
        util.SafeSendMessage(session, message.ChannelID, "No current statistics available for " + strings.Title(firstName + " " + lastName) + ".")
        return nil
    }

    var embeds []*models.Embed

    if len(o1sEncounters) > 0 {
        o1sEmbed := models.NewEmbed().
            SetTitle("O1S Statistics").
            SetDescription("For " + strings.Title(firstName + " " + lastName)).
            SetColor(0x0000ff)
        for _, e := range o1sEncounters {
            o1sEmbed.AddField(classes[e.Class], "**" + strconv.Itoa(int(e.DPS)) + " DPS**, **" + strconv.Itoa(encounterPercentile(e)) + "th** percentile")
        }

        embeds = append(embeds, o1sEmbed)
    }

    if len(o2sEncounters) > 0 {
        o2sEmbed := models.NewEmbed().
            SetTitle("O2S Statistics").
            SetDescription("For " + strings.Title(firstName + " " + lastName)).
            SetColor(0x7d7c7b)
        for _, e := range o2sEncounters {
            o2sEmbed.AddField(classes[e.Class], "**" + strconv.Itoa(int(e.DPS)) + " DPS**, **" + strconv.Itoa(encounterPercentile(e)) + "th** percentile")
        }

        embeds = append(embeds, o2sEmbed)
    }

    if len(o3sEncounters) > 0 {
        o3sEmbed := models.NewEmbed().
            SetTitle("O3S Statistics").
            SetDescription("For " + strings.Title(firstName + " " + lastName)).
            SetColor(0xff0000)
        for _, e := range o3sEncounters {
            o3sEmbed.AddField(classes[e.Class], "**" + strconv.Itoa(int(e.DPS)) + " DPS**, **" + strconv.Itoa(encounterPercentile(e)) + "th** percentile")
        }

        embeds = append(embeds, o3sEmbed)
    }

    if len(exdeathEncounters) > 0 {
        exdeathEmbed := models.NewEmbed().
            SetTitle("Exdeath Statistics").
            SetDescription("For " + strings.Title(firstName + " " + lastName)).
            SetColor(0x00007f)
        for _, e := range exdeathEncounters {
            exdeathEmbed.AddField(classes[e.Class], "**" + strconv.Itoa(int(e.DPS)) + " DPS**, **" + strconv.Itoa(encounterPercentile(e)) + "th** percentile")
        }

        embeds = append(embeds, exdeathEmbed)
    }

    if len(neoExdeathEncounters) > 0 {
        // Neo Exdeath Embed
        neoExdeathEmbed := models.NewEmbed().
            SetTitle("Neo Exdeath Statistics").
            SetDescription("For " + strings.Title(firstName + " " + lastName)).
            SetColor(0x00ff00)
        for _, e := range neoExdeathEncounters {
            neoExdeathEmbed.AddField(classes[e.Class], "**" + strconv.Itoa(int(e.DPS)) + " DPS**, **" + strconv.Itoa(encounterPercentile(e)) + "th** percentile")
        }

        embeds = append(embeds, neoExdeathEmbed)
    }

    for _, e := range embeds {
        util.SafeSendEmbed(session, message.ChannelID, e.MessageEmbed)
    }

    return nil
}

func initClasses(config *models.Config) error {
    classes = make(map[int]string)

    r, err := http.Get(constants.FFLOGS_API_ENDPOINT + constants.FFLOGS_API_CLASSES + "?api_key=" + config.FFLogsApiKey)
    if err != nil {
        return err
    }

    var respBody []models.FFLogsClassesResponse
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&respBody)
    if err != nil {
        return err
    }

    for _, e := range respBody[0].Classes {
        classes[e.ID] = e.Name
    }

    return nil
}

func encounterPercentile(encounter models.FFLogsEncounter) int {
    return int(float64(encounter.OutOf - encounter.Rank) / float64(encounter.OutOf) * 100)
}