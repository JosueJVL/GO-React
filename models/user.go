package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	//Names     string             `gorm:"type:varchar(20);not null;"`
	UserName  string    `bson:"userName" json:"userName,omitempty"`
	Name      string    `bson:"name" json:"name,omitempty"`
	LastName  string    `bson:"lastName" json:"lastName,omitempty"`
	BirthDay  time.Time `bson:"birthday" json:"birthday,omitempty"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password,omitempty"`
	Avatar    string    `bson:"avatar" json:"avatar,omitempty"`
	Banner    string    `bson:"banner" json:"banner,omitempty"`
	Biography string    `bson:"biography" json:"biography,omitempty"`
	Location  string    `bson:"location" json:"location,omitempty"`
	WebSite   string    `bson:"webSite" json:"webSite,omitempty"`
}
