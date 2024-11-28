package models

import "time"

type Student struct {
	Student_id    int       `json:"course_id"`
	Name          string    `json:"name"`
	Creation_date time.Time `json:"created"`
}

type NewStudent struct {
	Name string `json:"name"`
}
