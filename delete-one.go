package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonyaltoe/go-car-shop/api/helpers"
)

func (c carController) DeleteOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	_, customErr := c.CarService.DeleteOne(context.Background(), param["id"])
	if customErr != nil {
		helpers.SendError(customErr, res)
		return
	}

	var message = map[string]string{"message": ""}

	message["message"] = "deleted"
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(message)
}
