package models

import "encoding/json"

type PayPalRequest struct {
	// Headers
	AuthAlgo         string `json:"auth_algo"`
	CertURL          string `json:"cert_url"`
	TransmissionID   string `json:"transmission_id"`
	TransmissionSig  string `json:"transmission_sig"`
	TransmissionTime string `json:"transmission_time"`

	// Body
	ID           string          `json:"-"`
	ResourceID   string          `json:"-"`
	Status       string          `json:"-"`
	Custom       string          `json:"-"`
	Price        string          `json:"-"`
	WebhookID    string          `json:"webhook_id"`
	WebhookEvent json.RawMessage `json:"webhook_event"`
}
