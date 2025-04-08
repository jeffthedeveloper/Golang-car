package car_controller

import (
	"net/http"

	carService "github.com/jhonyaltoe/go-car-shop/api/services/car"
	"go.mongodb.org/mongo-driver/mongo"
)

type TCarController interface {
	GetAll(res http.ResponseWriter, req *http.Request)
	CreateOne(res http.ResponseWriter, req *http.Request)
	GetOne(res http.ResponseWriter, req *http.Request)
	DeleteOne(res http.ResponseWriter, req *http.Request)
	Update(res http.ResponseWriter, req *http.Request)
}

type carController struct {
	CarService carService.TCarService
}

func New(collection *mongo.Collection) TCarController {
	return &carController{
		CarService: carService.New(collection),
	}
}
