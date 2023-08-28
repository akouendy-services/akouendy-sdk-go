package otp

import (
	"fmt"

	"github.com/google/uuid"
)

type Status string

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Namespace          string `json:"namespace,omitempty"`
		ID                 string `json:"id,omitempty"`
		To                 string `json:"to,omitempty"`
		ChannelDescription string `json:"channel_description,omitempty"`
		AddressDescription string `json:"address_description,omitempty"`
		Provider           string `json:"provider,omitempty"`
		Otp                string `json:"otp,omitempty"`
		MaxAttempts        int    `json:"max_attempts,omitempty"`
		Attempts           int    `json:"attempts,omitempty"`
		Closed             bool   `json:"closed,omitempty"`
		TTL                int    `json:"ttl,omitempty"`
		URL                string `json:"url,omitempty"`
	} `json:"data,omitempty"`
}

type Error struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		TTLSeconds  int `json:"ttl_seconds,omitempty"`
		Attempts    int `json:"attempts,omitempty"`
		MaxAttempts int `json:"max_attempts,omitempty"`
	} `json:"data,omitempty"`
}

// Error implements go error interface.
func (msg *Error) Error() string {
	return fmt.Sprintf("API Error: %s", msg.Message)
}

type InitRequest struct {
	ID       uuid.UUID
	Provider string
	Receiver string
}

type ValidateRequest struct {
	ID         uuid.UUID
	Code       string
	SkipDelete bool
}
