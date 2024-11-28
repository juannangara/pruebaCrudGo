package models

import "time"

type Grade struct {
	Grade_id      int       `json:"course_id"`
	Course_id     int       `json:"name"`
	Grade         int       `json:"description"`
	Creation_date time.Time `json:"created"`
}
