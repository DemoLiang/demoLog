package demoLog

import "time"

type DemoLogS struct {
	Prefix string    `json:"prefix"`
	Module string    `json:"module"`
	Level  string    `json:"level"`
	Time   time.Time `json:"time"`
	File   string    `json:"file"`
	Line   int       `json:"line"`
	Cnt    string    `json:"cnt"`
}
