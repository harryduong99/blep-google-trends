package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trend struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	KeywordID primitive.ObjectID `json:"keyword_id" bson:"keyword_id"`      // tag golang
	Time      string             `json:"time" bson:"time"`
	Value     int                `json:"value" bson:"value"`
}

type Keyword struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Keyword string             `json:"keyword" bson:"keyword"`
}
