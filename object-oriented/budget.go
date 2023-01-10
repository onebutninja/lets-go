// Initialisers for New Functions

package main

import (
	"fmt"
	"time"
)

// Budget is a budgtest for campaign
type Budget struct {
	campaignID string
	Balance    float64 // USD
	Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty CAMPAIGN ID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("BALANCE MUST BE MORE THAN 0")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("BAD EXPIRATION DATE")
	}

	b := Budget{
		campaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}

	return &b, nil
}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)
	b1, err := NewBudget("puppies", 3.5, expires)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("%#v\n", b1)
	}
}
