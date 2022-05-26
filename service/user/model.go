package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjID        primitive.ObjectID `bson:"_id,omitempty"`
	ID           string             `json:"id"`
	Name         string             `bson:"name" json:"name"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Avatar       string             `bson:"avatar" json:"avatar"`
	Token        string             `bson:"tokenManager" json:"tokenManager"`
	RegisterTime int64              `bson:"registerTime" json:"registerTime"`
	UpdateTime   int64              `bson:"updateTime" json:"updateTime"`
}

type RegisterInformation struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	VerificationCode int    `json:"verificationCode"`
}

//db.user.createIndex({phone:1},{unique:true})
