package models

import (
	"time"
)

type Partner struct {
	Id      int       `json:"Id"`
	Name    string    `json:"Name"`
	Age     int       `json:"Age"`
	Dob     time.Time `json:"Dob"`
	Balance float32   `json:"Balance"`
	Access  bool      `json:"Access"`
}
