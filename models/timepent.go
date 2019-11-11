package models

import "time"

type Timespent struct {
	Duration int64     `json:"timespent,omitempty" bson:"timespent,omitempty"`
	Date     time.Time `json:"timecreated,omitempty" bson:"timecreated,omitempty"`
	Desc     string    `json:"desc,omitempty" bson:"desc,omitempty"`
}
