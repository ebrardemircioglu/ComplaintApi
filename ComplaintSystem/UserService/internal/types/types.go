package types

import "time"

type UserRequestModel struct {
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password" json:"password"`
	Name      string    `bson:"name" json:"name"`
	Surname   string    `bson:"surname" json:"surname"`
	Email     string    `bson:"email" json:"email"`
	Phone     string    `bson:"phone" json:"phone"`
	Address   string    `bson:"address" json:"address"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
type UserResponseModel struct {
	//ID---Id
	Id       string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
}
type UserUpdateModel struct {
	Password  string    `bson:"password" json:"password"`
	Phone     string    `bson:"phone" json:"phone"`
	Address   string    `bson:"address" json:"address"`
	Email     string    `bson:"email" json:"email"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}
