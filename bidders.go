package main

// Bidder represents a string name for a particular bidder
type Bidder string

const (
	// Nike represents Nike as a bidder/buyer on the ad exchange
	Nike Bidder = "nike"
	// Amazon represents Amazon as a bidder/buyer on the ad exchange
	Amazon Bidder = "amazon"
	// Apple represents Apple as a bidder/buyer on the ad exchange
	Apple Bidder = "apple"
)

var bidders = []Bidder{Nike, Apple, Amazon}

// GetAllBidders return the list of supported bidders
func GetAllBidders() []Bidder {
	return bidders
}
