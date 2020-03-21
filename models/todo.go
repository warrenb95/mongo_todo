package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Desc        string             `json:"desc,omitempty" bson:"desc,omitempty"`
	TimeCreated time.Time          `json:"timecreated,omitempty" bson:"timecreated,omitempty"`
	// Deadline       time.Time          `json:"deadline,omitempty" bson:"deadline,omitempty"`
	Estimate       int64       `json:"estimate,omitempty" bson:"estimate,omitempty"`
	TotalTimeSpent int64       `json:"totaltimespent,omitempty" bson:"totaltimespent,omitempty"`
	TimeSpent      []Timespent `json:"timespent,omitempty" bson:"timespent,omitempty"`
}
