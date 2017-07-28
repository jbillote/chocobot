package models

type FFLogsEncounter struct {
    EncounterID int     `json:"encounter"`
    Role        int     `json:"class"`
    Class       int     `json:"spec"`
    Guild       string  `json:"guild"`
    Rank        int     `json:"rank"`
    OutOf       int      `json:"outOf"`
    Duration    int     `json:"duration"`
    StartTime   int64   `json:"startTime"`
    ReportID    string  `json:"reportID"`
    FightID     int     `json:"fightID"`
    Difficulty  int     `json:"difficulty"`
    PartySize   int     `json:"size"`
    Patch       int     `json:"itemLevel"`
    DPS         float32 `json:"total"`
    Estimated   bool    `json:"estimated"`
}