package meeting

import (
	"time"
)

type Meeting struct {
	Date   time.Time `json:"date"`
	Orders []Order   `json:"orders"`
}
