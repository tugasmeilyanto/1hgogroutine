package main

import (
	"fmt"
	"time"
)

type Notification struct {
	UserID  int
	Message string
}

type Result struct {
	Index int
	Text  string
}

func sendEmailAsync(userID int, message string, index int, results chan<- Result) {
	time.Sleep(2 * time.Second)
	result := Result{
		Index: index,
		Text:  fmt.Sprintf("Email notification sent to user %d: %s", userID, message),
	}
	results <- result
}

func main() {
	notifications := []Notification{
		{UserID: 101, Message: "Your Order has been confirmed"},
		{UserID: 202, Message: "Your Account has been created"},
		{UserID: 303, Message: "Your Payment was successful"},
	}

	results := make(chan Result, len(notifications))

	for i, notification := range notifications {
		go sendEmailAsync(notification.UserID, notification.Message, i, results)
	}

	// Collect results in order
	orderedResults := make([]string, len(notifications))
	for i := 0; i < len(notifications); i++ {
		result := <-results
		orderedResults[result.Index] = result.Text
	}

	// Print results in order
	for _, res := range orderedResults {
		fmt.Println(res)
	}

	fmt.Println("Main application continues...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main application finished")
}
