package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("== Strava → Google Sheets Sync ==")
	now := time.Now()
	weekAgo := now.AddDate(0, 0, -7)
	fmt.Printf("Today:      %s\n", now.Format(time.RFC3339))
	fmt.Printf("Week ago:   %s\n", weekAgo.Format(time.RFC3339))

	fmt.Println("Step 1: here is where we will call Strava API.")
	fmt.Println("Step 2: here is where we will send data to Google Sheets.")
}
