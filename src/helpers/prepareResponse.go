package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, httpCode int, data interface{}) interface{} {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	resp := make(map[string]interface{})
	resp["message"] = "successfully return"
	resp["data"] = data
	userJson, _ := json.Marshal(resp)
	w.Write(userJson)
	return userJson
}
