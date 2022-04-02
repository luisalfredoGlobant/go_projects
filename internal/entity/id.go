package entity

import "github.com/google/uuid"

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func GenerateID() string {
	return uuid.New().String()
}

func GenerateUint32ID() uint32 {
	return uint32(uuid.New().ID())
}