package faq

import "time"

type FAQ struct {
	Question    string    `json:"question" bson:"question"`
	Link        string    `json:"link" bson:"link"`
	Description string    `json:"description" bson:"description"`
	Active      bool      `json:"active" bson:"active"`
	UpdateTime  time.Time `json:"updateTime" bson:"update_time"`
	CreateTime  time.Time `json:"createTime" bson:"create_time"`
}
