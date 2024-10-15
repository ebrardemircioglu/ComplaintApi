package internal

import (
	"ComplaintSystem/UserService/internal/types"
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
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
func (r *Repository) Create(ctx context.Context, user *types.User) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, user)
	return res, err
}
func (r *Repository) GetById(ctx context.Context, id string) (*types.User, error) {
	var result *types.User
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, echo.NewHTTPError(http.StatusNotFound, "No user found with the provided ID")
		}
		return nil, err
	}
	return result, nil
}
func (r *Repository) Delete(ctx context.Context, id string) error {

	filter := bson.M{"_id": id}
	result := r.collection.FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "No user found with the provided ID")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "An error occurred while deleting USER: "+err.Error())
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, id string, user *types.User) error {
	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": user}
	_, result := r.collection.UpdateOne(ctx, filter, update)
	return result
}
func (r *Repository) PartialUpdate(ctx context.Context, id string, user *types.User) error {
	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": user}
	_, result := r.collection.UpdateOne(ctx, filter, update)
	return result
}

func (r *Repository) GetAll(ctx context.Context, name string, surname string, address string) ([]types.User, error) {
	findOptions := options.Find().SetAllowPartialResults(true)
	var filter bson.A
	if name != "" {
		//filter = append(filter, bson.E{"name", name})
		filter = append(filter, bson.D{{"name", bson.D{{"$regex", name}, {"$options", "i"}}}})
	}
	if surname != "" {
		//filter = append(filter, bson.E{"surname", surname})
		filter = append(filter, bson.D{{"surname", bson.D{{"$regex", surname}, {"$options", "i"}}}})
	}
	if address != "" {
		//filter = append(filter, bson.E{"address", address})
		filter = append(filter, bson.D{{"address", bson.D{{"$regex", address}, {"$options", "i"}}}})
	}
	if len(filter) == 0 {
		filter = append(filter, bson.D{})

	}
	query := bson.D{{"$or", filter}}
	cursor, err := r.collection.Find(ctx, query, findOptions)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []types.User

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
