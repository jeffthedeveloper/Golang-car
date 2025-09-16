package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffthedeveloper/go-car-shop/api/helpers"
)

func (c *carController) GetOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	car, customErr := c.CarService.GetOne(context.Background(), param["id"])
	if customErr != nil {
		helpers.SendError(customErr, res)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(car)
}
