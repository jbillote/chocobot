package models

import "github.com/bwmarrin/discordgo"

const (
    EmbedLimitTitle       = 256
    EmbedLimitDescription = 2048
    EmbedLimitFieldName   = 256
    EmbedLimitFieldValue  = 1024
    EmbedLimitField       = 25
    EmbedLimitFooter      = 2048
    EmbedLimit            = 4000
)

type Embed struct {
    *discordgo.MessageEmbed
}

func NewEmbed() *Embed {
    return &Embed{ &discordgo.MessageEmbed{} }
}

func (e *Embed) SetTitle(title string) *Embed {
    if len(title) > EmbedLimitTitle {
        title = title[:EmbedLimitTitle]
    }

    e.Title = title
    return e
}

func (e *Embed) SetDescription(description string) *Embed {
    if len(description) > EmbedLimitDescription {
        description = description[:EmbedLimitDescription]
    }

    e.Description = description
    return e
}

func (e *Embed) AddField(name string, value string) *Embed {
    if len(e.Fields) > EmbedLimitField {
        return e
    }

    if len(name) > EmbedLimitFieldName {
        name = name[:EmbedLimitFieldName]
    }

    if len(value) > EmbedLimitFieldValue {
        value = value[:EmbedLimitFieldValue]
    }

    e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
        Name:  name,
        Value: value,
    })

    return e
}

func (e *Embed) SetFooter(args ...string) *Embed {
    iconURL := ""
    text := ""
    proxyURL := ""

    if len(args) == 0 {
        return e
    }

    if len(args) > 0 {
        if len(args[0]) > EmbedLimitFooter {
            args[0] = args[0][:EmbedLimitFooter]
        }
        text = args[0]
    }

    if len(args) > 1 {
        iconURL = args[1]
    }

    if len(args) > 2 {
        proxyURL = args[2]
    }

    e.Footer = &discordgo.MessageEmbedFooter{
        IconURL:      iconURL,
        Text:         text,
        ProxyIconURL: proxyURL,
    }

    return e
}

func (e *Embed) SetImage(args ...string) *Embed {
    url := ""
    proxyURL := ""

    if len(args) == 0 {
        return e
    }

    if len(args) > 0 {
        url = args[0]
    }

    if len(args) > 1 {
        proxyURL = args[1]
    }

    e.Image = &discordgo.MessageEmbedImage{
        URL:      url,
        ProxyURL: proxyURL,
    }

    return e
}

func (e *Embed) SetThumbnail(args ...string) *Embed {
    url := ""
    proxyURL := ""

    if len(args) == 0 {
        return e
    }

    if len(args) > 0 {
        url = args[0]
    }

    if len(args) > 1 {
        proxyURL = args[1]
    }

    e.Thumbnail = &discordgo.MessageEmbedThumbnail{
        URL:      url,
        ProxyURL: proxyURL,
    }

    return e
}

func (e *Embed) SetAuthor(args ...string) *Embed {
    name := ""
    iconURL := ""
    url := ""
    proxyURL := ""

    if len(args) == 0 {
        return e
    }

    if len(args) > 0 {
        name = args[0]
    }

    if len(args) > 1 {
        iconURL = args[1]
    }

    if len(args) > 2 {
        url = args[2]
    }

    if len(args) > 3 {
        proxyURL = args[3]
    }

    e.Author = &discordgo.MessageEmbedAuthor{
        Name:         name,
        IconURL:      iconURL,
        URL:          url,
        ProxyIconURL: proxyURL,
    }

    return e
}

func (e *Embed) SetURL(url string) *Embed {
    e.URL = url
    return e
}

func (e *Embed) SetColor(color int) *Embed {
    e.Color = color
    return e
}

func (e *Embed) InlineFields() *Embed {
    for _, v := range e.Fields {
        v.Inline = true
    }

    return e
}