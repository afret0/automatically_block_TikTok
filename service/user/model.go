package user

import "go.mongodb.org/mongo-driver/bson/primitive"

var role *RoleList

func init() {
	role = new(RoleList)
	role.DM = "DM"
	role.Customer = "customer"
	role.Boss = "boss"
}

type User struct {
	ObjID        primitive.ObjectID `bson:"_id,omitempty"`
	ID           string             `json:"id"`
	Name         string             `bson:"name" json:"name"`
	Phone        string             `bson:"phone" json:"phone"`
	Password     string             `bson:"password" json:"password"`
	WXName       string             `bson:"WXName" json:"WXName"`
	Role         string             `bson:"role" json:"role"`
	Token        string             `bson:"token" json:"token"`
	Store        string             `bson:"store" json:"store"`
	RegisterTime float64            `bson:"registerTime" json:"registerTime"`
	UpdateTime   float64            `bson:"updateTime" json:"updateTime"`
}

type RoleList struct {
	Boss     string
	DM       string
	Customer string
}
