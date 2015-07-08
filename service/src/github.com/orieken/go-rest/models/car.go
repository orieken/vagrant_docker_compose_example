package models

import "gopkg.in/mgo.v2/bson"

type (
	Car struct {
		Id bson.ObjectId `json:"id" bson:"_id"`
		Make string `json:"make" bson:"make"`
		Model string `json:"model" bson:"model"`
		Lap string `json:"lap" bson:"lap"`
		Conditions string `json:"conditions" bson:"conditions"`
		Series int `json:"series" bson:"series"`
		Episode int `json:"episode" bson:"episode"`
		Details string `json:"details" bson:"details"`
	}
)