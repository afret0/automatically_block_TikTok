package screenplay

import "go.mongodb.org/mongo-driver/bson/primitive"

type Screenplay struct {
	ObjId        primitive.ObjectID `bson:"_id"`
	Id           string             `bson:"id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Cover        string             `bson:"cover" json:"cover"`
	Instruction  string             `bson:"instruction" json:"instruction"`
	Store        string             `bson:"store" json:"store"`
	Note         string             `bson:"note" json:"note"`
	Noun         *Noun              `bson:"noun" json:"noun"`
	Tag          []string           `bson:"tag" json:"tag"`
	Style        string             `bson:"style" json:"style"`
	RegisterTime float64            `bson:"registerTime" json:"registerTime"`
	UpdateTime   float64            `bson:"updateTime" json:"updateTime"`
}

type Noun struct {
	Num      int  `bson:"num" json:"num"`
	Boy      int  `bson:"boy" json:"boy"`
	Girl     int  `bson:"girl" json:"girl"`
	CrossSex bool `bson:"crossSex" json:"crossSex"`
}

/*
db.screenplay.createIndex({name:1},{unique:true})
db.screenplay.createIndex({store:1,tag:1})
db.screenplay.createIndex({store:1,style:1})
db.screenplay.createIndex({store:1,noun.num:1})
*/