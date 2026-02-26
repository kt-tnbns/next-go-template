package entity

import "time"

// Timestamp holds createdAt and updatedAt fields for entities.
// Embed this in any entity that needs audit timestamps.
type Timestamp struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
