package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// Formatting time into a readable string
	formattedTime := currentTime.Format("2006-01-02 15:04:05 MST")
	fmt.Println("Formatted Time:", formattedTime)

	// Parsing a date string
	layout := "2006-01-02 15:04:05"
	dateStr := "2025-02-28 14:30:00"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}

	// Getting individual components
	year, month, day := currentTime.Date()
	hour, minute, second := currentTime.Clock()
	fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, minute, second)

	// Adding time duration
	futureTime := currentTime.Add(48 * time.Hour)
	fmt.Println("Future Time (after 48 hours):", futureTime)

	// Subtracting two dates
	duration := futureTime.Sub(currentTime)
	fmt.Println("Duration between two dates:", duration)

	// Creating a specific date
	specificDate := time.Date(2025, time.March, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific Date:", specificDate)
}
