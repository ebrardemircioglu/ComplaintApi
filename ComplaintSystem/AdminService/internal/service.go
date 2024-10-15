package internal

import (
	"ComplaintSystem/AdminService/internal/types"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, adminRequestModel *types.AdminRequestModel) (string, error) {

	existedAdmin, err := s.repo.GetByCompanyName(ctx, adminRequestModel.CompanyName)
	/*if err != nil {
		return "", err
	}

	if existedAdmin != nil {
		return "", fmt.Errorf("şirket adı mevcut kullanamzsınız ")
	}*/
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Admin bulunamadı, bu hata değil, devam edebilirsin
			existedAdmin = nil
		} else {
			return "", err
		}
	}

	if existedAdmin != nil {
		return "", fmt.Errorf("şirket adı mevcut, kullanamazsınız")
	}

	adminID := uuid.New().String()
	now := time.Now().Local()
	adminRequestModel.CreatedAt = now

	admin := &types.Admin{
		CompanyName: adminRequestModel.CompanyName,
		Phone:       adminRequestModel.Phone,
		Address:     adminRequestModel.Address,
		Email:       adminRequestModel.Email,
		CreatedAt:   adminRequestModel.CreatedAt,
		Id:          adminID,
		Password:    adminRequestModel.Password,
		//Employee:    adminRequestModel.Employee,
		//SocialMedia: adminRequestModel.SocialMedia,
		EmployeeId:   adminRequestModel.EmployeeId,
		EmployeeRole: adminRequestModel.EmployeeRole,
	}
	_, err = s.repo.Create(ctx, admin)

	if err != nil {
		return "", err
	}
	return adminID, nil
}
func (s *Service) GetAll(ctx context.Context) ([]types.AdminResponseModel, error) {
	admin, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *Service) GetByCompanyName(ctx context.Context, companyName string) (*types.Admin, error) {
	admin, err := s.repo.GetByCompanyName(ctx, companyName)
	if admin.AdditionalInfo == nil {
		admin.AdditionalInfo = make(map[string]string)
		switch admin.EmployeeRole {
		case "ceo":
			admin.AdditionalInfo["decision_making_authority"] = "100"
			admin.AdditionalInfo["budget_approval"] = "Unlimited"
		case "human_resources":
			admin.AdditionalInfo["employee_management"] = "Active"
			admin.AdditionalInfo["training_programs"] = "Available"
			admin.AdditionalInfo["recruitment"] = "Ongoing"
		case "community_manager":
			admin.AdditionalInfo["customer_engagement"] = "High"
			admin.AdditionalInfo["event_management"] = "Active"
		case "legal_advisor":
			admin.AdditionalInfo["membership_type"] = "Premium"
		case "socialmedia_expert":
			admin.AdditionalInfo["platforms_managed"] = "Facebook, Instagram, Twitter,Linkedin"
			admin.AdditionalInfo["campaigns"] = "Active"
		case "product_manager":
			admin.AdditionalInfo["product_launches"] = "Scheduled"
		case "it_support":
			admin.AdditionalInfo["response_time"] = "24 hours" // Yanıt süresi
			admin.AdditionalInfo["technical_support"] = "Available"
		default:
			admin.AdditionalInfo["employeeRole"] = "None"
		}
		fmt.Println("Admin Company Name:", companyName)
		fmt.Println("Admin Employee Role:", admin.EmployeeRole)
		fmt.Println("Additional Info:", admin.AdditionalInfo)
	}
	if err != nil {
		return nil, err
	}
	return admin, nil
}
func (s *Service) Update(ctx context.Context, companyName string, adminUpdateModel types.AdminUpdateModel) error {
	admin, err := s.GetByCompanyName(ctx, companyName)
	if err != nil {
		return err
	}
	now := time.Now().Local()
	adminUpdateModel.UpdatedAt = now
	admin.Phone = adminUpdateModel.Phone
	admin.Password = adminUpdateModel.Password
	admin.Address = adminUpdateModel.Address
	admin.Email = adminUpdateModel.Email
	admin.UpdatedAt = adminUpdateModel.UpdatedAt
	return s.repo.Update(ctx, companyName, admin)
}

/*
func (s *Service) PartialUpdate(ctx context.Context, companyName string, employee types.Employee) error {

		admin, err := s.GetByCompanyName(ctx, companyName)
		if err != nil {
			return fmt.Errorf("company not found: %w", err)
		}

		if employee.Name != "" {
			admin.Employee.Name = employee.Name
		}
		if employee.Surname != "" {
			admin.Employee.Surname = employee.Surname
		}
		if employee.Role != "" {
			admin.Employee.Role = employee.Role
		}
		admin.Employee.UpdatedAt = time.Now()

		err = s.repo.PartialUpdate(ctx, companyName, &admin.Employee)
		if err != nil {
			return fmt.Errorf("database update failed: %w", err)
		}

		return nil
	}
*/

func (s *Service) PartialUpdate(ctx context.Context, companyName string, adminUpdateModel types.AdminUpdateModel) error {
	admin, err := s.GetByCompanyName(ctx, companyName)
	if err != nil {
		return err
	}
	now := time.Now().Local()
	admin.UpdatedAt = now
	admin.EmployeeRole = adminUpdateModel.EmployeeRole
	admin.EmployeeId = adminUpdateModel.EmployeeId
	admin.UpdatedAt = adminUpdateModel.UpdatedAt
	return s.repo.PartialUpdate(ctx, companyName, admin)
}
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)

}
