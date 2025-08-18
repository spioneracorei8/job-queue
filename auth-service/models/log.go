package models

import "time"

type LogForm struct {
	Time     time.Time      `json:"time"`
	Service  string         `json:"service"`
	Level    string         `json:"level"`
	Message  string         `json:"message"`
	Metadata map[string]any `json:"metadata"`
	Method   any            `json:"method"`
	Path     any            `json:"path"`
	Ip       any            `json:"ip"`
}
