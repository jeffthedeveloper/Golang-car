package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jeffthedeveloper/Golang-car/api/database/entities"
	"github.com/jeffthedeveloper/Golang-car/api/helpers"
)

func (c *carController) CreateOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	var car *entities.TCar
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		customErr := helpers.CustomErrBuilder(
			http.StatusBadRequest,
			"Wrong json body",
			"Controller.CreateOne",
			err,
		)
		helpers.SendError(customErr, res)
		return
	}

	createdCar, custErr := c.CarService.CreateOne(context.Background(), car)
	if custErr != nil {
		helpers.SendError(custErr, res)
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(*createdCar)
}
