package models

type FFLogsClassesResponse struct {
    ID      int           `json:"id"`
    Name    string        `json:"name"`
    Classes []FFLogsClass `json:"specs"`
}