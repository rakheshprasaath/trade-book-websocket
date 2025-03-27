package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rakheshprasaath/trade-book-websocket/database"
	"github.com/rakheshprasaath/trade-book-websocket/models"
)

func ProcessData(msg string) {
	// Create a map to hold JSON key-value pairs
	var data map[string]interface{}
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Parse date strings to time.Time
	timeEntryParsed, err := time.Parse("2006-01-02 15:04:05", data["time_entry"].(string))
	if err != nil {
		log.Fatalf("Error parsing time_entry: %v", err)
	}
	timeParsed, err := time.Parse("2006-01-02 15:04:05", data["time"].(string))
	if err != nil {
		log.Fatalf("Error parsing time: %v", err)
	}

	// Manually assign values to the struct fields
	historyOrder := models.HistoryOrder{
		AccountID: int(data["account_id"].(float64)), // Convert float64 to int
		Type:      data["type"].(string),
		TimeEntry: timeEntryParsed, // Assign parsed time.Time
		Symbol:    data["symbol"].(string),
		Ticket:    int(data["ticket"].(float64)),
		Volume:    data["volume"].(float64),
		Price:     data["price"].(float64),
		Value:     data["value"].(float64),
		SL:        data["sl"].(float64),
		TP:        data["tp"].(float64),
		Time:      timeParsed, // Assign parsed time.Time
		State:     data["state"].(string),
		Magic:     int(data["magic"].(float64)),
		Comment:   data["comment"].(string),
	}

	// Print the manually assigned struct
	fmt.Printf("Manually Assigned HistoryOrder Struct: %+v\n", historyOrder)

	// Save historyOrder to the database
	result := database.DB.Create(&historyOrder)
	if result.Error != nil {
		log.Fatalf("Failed to insert record: %v", result.Error)
	}


	

}
