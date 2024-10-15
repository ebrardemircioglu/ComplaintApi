package internal

import (
	"ComplaintSystem/AdminService/internal/types"
	"context"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{
		collection: col,
	}
}
func (r *Repository) Create(ctx context.Context, admin *types.Admin) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, admin)
	return res, err
}

/*
	func (r *Repository) GetByCompanyName(ctx context.Context, companyName string) (*types.Admin, error) {
		var result *types.Admin
		filter := bson.M{"companyName": companyName}
		err := r.collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, echo.NewHTTPError(http.StatusNotFound, "No admin found with the provided ID")
			}
			return nil, err
		}
		return result, nil
	}
*/
func (r *Repository) GetByCompanyName(ctx context.Context, companyName string) (*types.Admin, error) {
	var result *types.Admin
	filter := bson.M{"companyName": companyName}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
func (r *Repository) Update(ctx context.Context, companyName string, admin *types.Admin) error {
	filter := bson.D{{"_id", companyName}}
	update := bson.M{"$set": admin}
	_, result := r.collection.UpdateOne(ctx, filter, update)
	return result
}

/*func (r *Repository) PartialUpdate(ctx context.Context, companyName string, employee *types.Admin) error {
	filter := bson.D{{"companyName", companyName}}
	update := bson.M{"$set": bson.M{
		"employee.name":      employee.Name,
		"employee.surname":   employee.Surname,
		"employee.updatedAt": employee.UpdatedAt,
		"employee.role":      employee.Role,

	},
	}
	_, result := r.collection.UpdateOne(ctx, filter, update)
	return result
}*/

/*_, err := r.collection.UpdateOne(ctx, filter, update)

if err != nil {
	return err
}*/
//return nil

func (r *Repository) PartialUpdate(ctx context.Context, companyName string, admin *types.Admin) error {
	filter := bson.D{{"companyName", companyName}}
	update := bson.M{"$set": bson.M{
		"employee.employeeId":   admin.EmployeeId,
		"employee.employeeRole": admin.EmployeeRole,
	},
	}
	_, result := r.collection.UpdateOne(ctx, filter, update)
	return result
}
func (r *Repository) GetAll(ctx context.Context) ([]types.AdminResponseModel, error) {
	var admin []types.AdminResponseModel
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		//if err == mongo.ErrNoDocuments {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &admin); err != nil {
		return nil, err
	}

	return admin, nil
}
func (r *Repository) Delete(ctx context.Context, id string) error {

	filter := bson.D{{"id", id}}
	result := r.collection.FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "No admin found with the provided ID")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "An error occurred while deleting ADMIN: "+err.Error())
	}
	return nil
}
