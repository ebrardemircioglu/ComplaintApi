package types

import (
	"time"
)

type Admin struct {
	CompanyName string `bson:"companyName" json:"companyName" validate:"required,min=4,max=30,companyNameFormat"`
	Id          string `bson:"_id" json:"id"`
	Email       string `bson:"email" json:"email" `
	Phone       string `bson:"phone" json:"phone" validate:"required,phoneNumber"`
	Password    string `bson:"password" json:"password"`
	Address     string `bson:"address" json:"address"`
	Category    string `bson:"category" json:"category"`
	//SocialMedia SocialMedia `bson:"socialMedia" json:"socialMedia"`
	//Employee     Employee  `bson:"employee" json:"employee"`
	EmployeeId     string            `bson:"employeeId" json:"employeeId" `
	EmployeeRole   string            `bson:"employeeRole" json:"employeeRole"`
	CreatedAt      time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time         `bson:"updatedAt" json:"updatedAt"`
	AdditionalInfo map[string]string `bson:"additionalInfo" json:"additionalInfo"`
	/*CustomValidator struct {
		validator *validator.Validate
	}*/
}

/*type SocialMedia struct {
	Facebook string `bson:"facebook" json:"facebook"`
	Twitter  string `bson:"twitter" json:"twitter"`
	LinkedIn string `bson:"linkedin" json:"linkedin"`
	Website  string `bson:"website"   json:"website"`
}
type Employee struct {
	EmployeeId string    `bson:"employeeId" json:"employeeId"`
	Name       string    `bson:"name" json:"name"`
	Surname    string    `bson:"surname" json:"surname"`
	Role       string    `bson:"role" json:"role"`
	UpdatedAt  time.Time `bson:"updatedAt" json:"updatedAt"`
}*/
