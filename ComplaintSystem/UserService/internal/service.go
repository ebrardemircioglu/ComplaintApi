package internal

import (
	"ComplaintSystem/UserService/internal/types"
	"context"
	"github.com/google/uuid"
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

// Return the generated ID if the insertion is successful

func (s *Service) Create(ctx context.Context, userRequestModel *types.UserRequestModel) (string, error) {
	//err := AdminService.repository.db.Create(&types.AdminRequestModel).Error
	/*_, err := AdminService.repo.Create(ctx, adminRequestModel)
	if err != nil {
		return "", err
	}*/
	userID := uuid.New().String()
	now := time.Now().Local()
	userRequestModel.CreatedAt = now

	user := &types.User{
		Name:      userRequestModel.Name,
		Phone:     userRequestModel.Phone,
		Address:   userRequestModel.Address,
		Surname:   userRequestModel.Surname,
		Email:     userRequestModel.Email,
		CreatedAt: userRequestModel.CreatedAt,
		Id:        userID,
		Username:  userRequestModel.Username,
		Password:  userRequestModel.Password,
	}
	_, err := s.repo.Create(ctx, user)

	if err != nil {
		return "", err
	}
	return userID, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*types.User, error) {
	result, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	//challenge (everything should be observable somehow in the response or console (print)):
	// 1) do something with using for loop by using customer model and manipulate it (you can add an additional field for it)
	// 2) do something with switch-case
	// 3) do something with goroutines (you should give us an example for both scenarios of not using goroutines and using)
	// 3.1) calculate the elapsed time for both scenarios and show us the gained time
	// 4) add an additional field and use maps
	// 5) add an additional field and use arrays
	// 6) manipulate an existing data to see how pointers and values work
	return result, nil
}
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)

}
func (s *Service) GetAll(ctx context.Context, name string, surname string, address string) ([]types.User, error) {
	result, err := s.repo.GetAll(ctx, name, surname, address)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *Service) Update(ctx context.Context, id string, userUpdateModel types.UserUpdateModel) error {
	user, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	now := time.Now().Local()
	userUpdateModel.UpdatedAt = now
	user.Phone = userUpdateModel.Phone
	user.Password = userUpdateModel.Password
	user.Address = userUpdateModel.Address
	user.Email = userUpdateModel.Email
	user.UpdatedAt = userUpdateModel.UpdatedAt
	return s.repo.Update(ctx, id, user)
}

// Communication secimleri
func (s *Service) PartialUpdate(ctx context.Context, id string, userUpdateModel types.UserUpdateModel) error {
	user, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	now := time.Now().Local()
	userUpdateModel.UpdatedAt = now
	user.Phone = userUpdateModel.Phone
	user.Email = userUpdateModel.Email
	user.UpdatedAt = userUpdateModel.UpdatedAt
	return s.repo.PartialUpdate(ctx, id, user)
}

/*func (AdminService *Service) Delete(ctx context.Context, id string) error {
	return AdminService.repo.Delete(ctx, id)
}*/
