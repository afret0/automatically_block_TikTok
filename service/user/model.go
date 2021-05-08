package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjID        primitive.ObjectID `bson:"_id,omitempty"`
	ID           string             `json:"id"`
	Name         string             `bson:"name" json:"name"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	Phone        string             `bson:"phone" json:"phone"`
	Sex          int                `bson:"sex" json:"sex"`
	WXName       string             `bson:"WXName" json:"WXName"`
	DM           bool               `bson:"dm" json:"dm"`
	Boss         bool               `bson:"boss" json:"boss"`
	Token        string             `bson:"tokenManager" json:"tokenManager"`
	Store        string             `bson:"store" json:"store"`
	RegisterTime float64            `bson:"registerTime" json:"registerTime"`
	UpdateTime   float64            `bson:"updateTime" json:"updateTime"`
}

//db.user.createIndex({phone:1},{unique:true})
