package car_service

import (
	"context"

	"github.com/jeffthedeveloper/go-car-shop/api/database/entities"
	carModel "github.com/jeffthedeveloper/go-car-shop/api/database/models/car"
	"github.com/jeffthedeveloper/go-car-shop/api/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarService interface {
	GetAll(ctx context.Context) (*[]entities.TCar, *helpers.CustomError)
	CreateOne(ctx context.Context, car *entities.TCar) (*entities.TCar, *helpers.CustomError)
	GetOne(ctx context.Context, id string) (*entities.TCar, *helpers.CustomError)
	DeleteOne(ctx context.Context, id string) (*mongo.DeleteResult, *helpers.CustomError)
	Update(ctx context.Context, id string, car *entities.TCar) (*entities.TCar, *helpers.CustomError)
}

type carService struct {
	carModel carModel.TCarModel
}

func New(collection *mongo.Collection) TCarService {
	return &carService{
		carModel.New(collection),
	}
}
