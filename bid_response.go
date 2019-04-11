package main

import "net/url"

// BidResponse represents a bidder's response for a bid request
type BidResponse struct {
	BidderName Bidder `json:"bidder_name"`
	Bids       []Bid  `json:"bids"`
}

// Bid represents a bid for a single ad placement
type Bid struct {
	PlacementID int          `json:"placement_id"`
	Price       float64      `json:"price"`
	AdCreatives []AdCreative `json:"ad_creatives"`
}

// Size represents the size of an ad
type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// AdCreative represents an ad creative information
type AdCreative struct {
	Size Size    `json:"size"`
	URL  url.URL `json:"url"`
}
