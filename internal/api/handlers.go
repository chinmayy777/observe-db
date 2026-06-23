package api

import (
	"encoding/json"
	"net/http"
	"observedb/internal/storage"
)

var StoreInstance *storage.Store

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	var req PutRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
}
