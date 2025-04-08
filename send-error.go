package helpers

import (
	"fmt"
	"net/http"
)

func SendError(c *CustomError, res http.ResponseWriter) {
	fmt.Printf("\nMessage: %v\nMethod: the error ocurred in %v\n%v", c.Message, c.Method, c.Err)
	res.WriteHeader(c.StatusCode)
	res.Write([]byte(`{ "message": "` + c.Message + `" }`))
}
