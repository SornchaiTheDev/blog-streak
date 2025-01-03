package models

import "time"

type Streak struct {
	StartedDate time.Time `json:"started_date"`
	LatestDate  time.Time `json:"latest_date"`
}
