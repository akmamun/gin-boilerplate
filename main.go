package main

import (
	"context"
	"encoding/json"
	"go-pg/src/database"
	"log"
	"net/http"
	"time"
)

type Datas struct {
	Data      *string   `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//var d *Data
func testRR(w http.ResponseWriter, req *http.Request) {
	userData := Datas{}
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	err = database.Save(ctx, &userData)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]string)
	resp["message"] = "successfully return"

	userJson, err := json.Marshal(resp)
	w.Write(userJson)
	return

}

func main() {

	http.HandleFunc("/test", testRR)
	err := http.ListenAndServe(":8090", nil)
	log.Fatal(err)
}
