package entity

import (
	"math/big"
	"time"
)

type Room struct {
	ID        big.Int
	UserA     string
	UserB     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
