package types

import "time"

type User struct {
	Id        string    `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Surname   string    `bson:"surname" json:"surname"`
	Email     string    `bson:"email" json:"email"`
	Phone     string    `bson:"phone" json:"phone"`
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password" json:"password"`
	Address   string    `bson:"address" json:"address"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
