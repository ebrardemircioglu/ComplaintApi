package internal

import (
	"ComplaintSystem/AdminService/internal/types"
	_ "ComplaintSystem/AdminService/internal/types"
)

func ToAdminResponse(admin *types.Admin) *types.AdminResponseModel {
	return &types.AdminResponseModel{
		Id:           admin.Id,
		CompanyName:  admin.CompanyName,
		Category:     admin.Category,
		Email:        admin.Email,
		Phone:        admin.Phone,
		Address:      admin.Address,
		EmployeeRole: admin.EmployeeRole,
		EmployeeId:   admin.EmployeeId,
	}
}
