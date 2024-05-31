package main

import "time"

type Quote struct {
	Quote     string    `json:"quote"`
	Author    string    `json:"author"`
	DateAdded time.Time `json:"date"`
}
