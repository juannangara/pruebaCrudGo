package models

import (
	"time"
)

type Course struct {
	Course_id     int       `json:"course_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Creation_date time.Time `json:"created"`
}
