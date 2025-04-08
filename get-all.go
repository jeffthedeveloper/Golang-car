package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (c *carController) GetAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	cars, customErr := c.CarService.GetAll(context.Background())
	if customErr != nil {
		helpers.SendError(customErr, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(cars)
}
