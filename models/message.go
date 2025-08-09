// Package models defines the structure for chat messages in the application.
package models

import "time"

type Message struct {
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
