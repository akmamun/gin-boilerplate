package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteSuccess(w http.ResponseWriter, code int, data interface{}) error {
	fmt.Println(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	//item := Item{Item: "success"}
	return json.NewEncoder(w).Encode(data)
}
