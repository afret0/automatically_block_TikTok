package store

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	ObjID        primitive.ObjectID `bson:"_id"`
	ID           string             `json:"id"`
	Name         string             `bson:"name" json:"name"`
	Address      string             `bson:"address" json:"address"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	Boss         string             `bson:"boss" json:"boss"`
	DM           []string           `bson:"dm" json:"dm"`
	RegisterTime float64            `bson:"registerTime" json:"registerTime"`
	UpdateTime   float64            `bson:"updateTime" json:"updateTime"`
}

//db.store.createIndex({name:1},{unique:true})
