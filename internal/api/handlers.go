package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"observedb/internal/storage"
)

var StoreInstance *storage.Store

type PutRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PutRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	StoreInstance.Put(req.Key, req.Value)

	json.NewEncoder(w).Encode(map[string]string{"status": "stored"})
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	fmt.Println("Requested key: ", key)

	value, ok := StoreInstance.Get(key)

	if !ok {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"value": value})
}
