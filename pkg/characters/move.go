package characters

import "time"

type MovementRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// CharacterMovementDataSchema struct for CharacterMovementDataSchema
type Movement struct {
	// Cooldown details
	Cooldown CooldownSchema `json:"cooldown"`
	// Destination details.
	Destination MapSchema `json:"destination"`
	// Character details.
	Character Character `json:"character"`
}

type CooldownSchema struct {
	// The total seconds of the cooldown.
	TotalSeconds int32 `json:"total_seconds"`
	// The remaining seconds of the cooldown.
	RemainingSeconds int32 `json:"remaining_seconds"`
	// The start of the cooldown.
	StartedAt time.Time `json:"started_at"`
	// The expiration of the cooldown.
	Expiration time.Time `json:"expiration"`
	// The reason of the cooldown.
	Reason string `json:"reason"`
}

type MapSchema struct {
	// Name of the map.
	Name string `json:"name"`
	// Skin of the map.
	Skin string `json:"skin"`
	// Position X of the map.
	X int32 `json:"x"`
	// Position Y of the map.
	Y       int32            `json:"y"`
	Content MapContentSchema `json:"content"`
}

type MapContentSchema struct {
	// Type of the content.
	Type string `json:"type"`
	// Code of the content.
	Code string `json:"code"`
}
