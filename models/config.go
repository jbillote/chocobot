package models

type Config struct {
    BotToken      string `json:"bot_token"`
    FFLogsApiKey  string `json:"fflogs_api_key"`
    CommandPrefix string `json:"command_prefix"`
}