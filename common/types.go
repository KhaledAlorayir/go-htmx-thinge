package common

import "time"

type JWT struct {
	Jwt       string
	ExpiresAt time.Time
}
