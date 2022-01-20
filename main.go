package main

import (
	"context"
	"encoding/json"
	"go-pg/src/database"
	"go-pg/src/helpers"
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
	if req.Method == http.MethodPost {
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	resp := make(map[string]string)
	resp["message"] = "method not allowed"
	userJson, _ := json.Marshal(resp)
	w.Write(userJson)
	return

}

func fetchDATA(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	userDatas := make([]Datas, 0)

	err := database.Get(ctx, &userDatas, "data", 10, 10)
	if err != nil {
		return
	}
	helpers.ResponseSuccess(w, 200, &userDatas)

	//resp := make(map[string]interface{})
	//resp["message"] = "successfully return"
	//resp["data"] = userDatas
	//userJson, _ := json.Marshal(resp)
	//w.Write(userJson)
	return

}

func main() {

	http.HandleFunc("/test", testRR)
	http.HandleFunc("/fetch", fetchDATA)
	err := http.ListenAndServe(":8090", nil)
	log.Fatal(err)
}
