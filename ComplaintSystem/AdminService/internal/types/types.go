package types

import "time"

type AdminRequestModel struct {
	CompanyName string `bson:"companyName" json:"companyName" validate:"required,min=4,max=30,companyNameFormat"`
	Email       string `bson:"email" json:"email"`
	Phone       string `bson:"phone" json:"phone" validate:"required,phoneNumber"`
	Password    string `bson:"password" json:"password"`
	Address     string `bson:"address" json:"address"`
	Category    string `bson:"category" json:"category"`
	//SocialMedia SocialMedia `bson:"socialMedia" json:"socialMedia"`
	//Employee Employee `bson:"employee" json:"employee"`
	EmployeeId   string    `bson:"employeeId" json:"employeeId"`
	EmployeeRole string    `bson:"employeeRole" json:"employeeRole"`
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`
}

type AdminResponseModel struct {
	Id           string `bson:"_id" json:"id"`
	CompanyName  string `bson:"companyName" json:"companyName" validate:"required,min=4,max=30,companyNameFormat"`
	Category     string `bson:"category" json:"category"`
	Email        string `bson:"email" json:"email"`
	Phone        string `bson:"phone" json:"phone" validate:"required,phoneNumber"`
	Address      string `bson:"address" json:"address"`
	EmployeeId   string `bson:"employeeId" json:"employeeId"`
	EmployeeRole string `bson:"employeeRole" json:"employeeRole"`
	//Employee    Employee `bson:"employee" json:"employee"`
}
type AdminUpdateModel struct {
	Email    string `bson:"email" json:"email"`
	Phone    string `bson:"phone" json:"phone"`
	Password string `bson:"password" json:"password"`
	//Employee Employee `bson:"employee" json:"employee"`
	Address      string    `bson:"address" json:"address"`
	EmployeeId   string    `bson:"employeeId" json:"employeeId"`
	EmployeeRole string    `bson:"employeeRole" json:"employeeRole"`
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`
}
