package routes

import (
	"github.com/gorilla/mux"
	carController "github.com/jeffthedeveloper/go-car-shop/api/controllers/car"
	connect "github.com/jeffthedeveloper/go-car-shop/api/database"
)

func CarRoutes() *mux.Router {
	carCollection := connect.NewCollection("CarShop", "cars")
	carCtrl := carController.New(carCollection)

	router := mux.NewRouter()
	router.HandleFunc("/cars", carCtrl.GetAll).Methods("GET")
	router.HandleFunc("/cars", carCtrl.CreateOne).Methods("POST")
	router.HandleFunc("/cars/{id}", carCtrl.GetOne).Methods("GET")
	router.HandleFunc("/cars/{id}", carCtrl.DeleteOne).Methods("DELETE")
	router.HandleFunc("/cars/{id}", carCtrl.Update).Methods("PUT")

	return router
}
