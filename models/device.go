package models

type Device struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Type  string `json:"type"`
    Owner string `json:"owner"`
}
