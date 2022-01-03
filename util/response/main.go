package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  int
	Data    interface{}
	Message interface{}
}

func (r Response) ResponseJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{
		Status:  r.Status,
		Data:    r.Data,
		Message: r.Message,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (r Response) ResponseNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
