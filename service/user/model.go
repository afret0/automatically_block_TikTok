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
	Avatar       string             `bson:"avatar" json:"avatar"`
	Phone        string             `bson:"phone" json:"phone"`
	Sex          int                `bson:"sex" json:"sex"`
	WXName       string             `bson:"WXName" json:"WXName"`
	DM           bool               `bson:"dm" json:"dm"`
	Owner        string             `bson:"owner" json:"owner"`
	Token        string             `bson:"tokenManager" json:"tokenManager"`
	Store        string             `bson:"store" json:"store"`
	RegisterTime float64            `bson:"registerTime" json:"registerTime"`
	UpdateTime   float64            `bson:"updateTime" json:"updateTime"`
}

//db.user.createIndex({phone:1},{unique:true})

type RoleList struct {
	Boss     string
	DM       string
	Customer string
}

func GetRole() *RoleList {
	return role
}
