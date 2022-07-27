package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func testingMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

type customer struct {
	Name    string
	City    string
	ZipCode string
}

func getAllConsumner(w http.ResponseWriter, r *http.Request) {
	customers := []customer{
		{"andi", "jakarta", "109090"},
		{"farah", "jakarta", "109090"},
	}

	json.NewEncoder(w).Encode(customers)
}
