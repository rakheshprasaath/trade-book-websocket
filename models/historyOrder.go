package models

import(
	"time"
)


type HistoryOrder struct{
	AccountID int       `json:"account_id"`  // int to match the integer value in the JSON
	TimeEntry time.Time `json:"time_entry"`  // using time.Time for datetime parsing
	Symbol    string    `json:"symbol"`      // string for currency symbol
	Ticket    int       `json:"ticket"`      // int to store ticket ID
	Type      string    `json:"type"`        // string for the trade type
	Volume    float64   `json:"volume"`      // float64 for fractional volume
	Price     float64   `json:"price"`       // float64 for the trade price
	Value     float64   `json:"value"`       // float64 for the value in trade
	SL        float64   `json:"sl"`          // Stop Loss as float64
	TP        float64   `json:"tp"`          // Take Profit as float64
	Time      time.Time `json:"time"`        // time.Time for datetime parsing
	State     string    `json:"state"`       // string for the state of the trade
	Magic     int       `json:"magic"`       // int for the magic number
	Comment   string    `json:"comment"`     // string for the trade comment
}

