package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"id":        1202,
		"full_name": "Nguyen Duc Thuan",
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
}
