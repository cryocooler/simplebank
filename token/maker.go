package token

import "time"

// provides token creation and validation methods
type Maker interface {
	// createtoken creates a new token with a duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
