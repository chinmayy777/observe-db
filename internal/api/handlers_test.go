package api

import (
	"net/http"
	"net/http/httptest"
	"observedb/internal/storage"
	"strings"
	"testing"
)

func TestGetMissingKey(t *testing.T) {
	StoreInstance = storage.NewStore()

	req := httptest.NewRequest(http.MethodGet, "/get?key=missing", nil)

	rr := httptest.NewRecorder()

	GetHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rr.Code)
	}

}

func TestValidGet(t *testing.T) {
	StoreInstance = storage.NewStore()

	StoreInstance.Put("name", "chinmay")

	req := httptest.NewRequest(http.MethodGet, "/get?key=name", nil)

	rr := httptest.NewRecorder()

	GetHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
		t.Log(rr.Body.String())
	}
	if !strings.Contains(rr.Body.String(), "chinmay") {
		t.Errorf("expected response to contain chinmay")
	}
}

func TestValidPut(t *testing.T) {
	StoreInstance = storage.NewStore()

	body := strings.NewReader(
		`{
			"key" : "name",
			"value" : "chinmay"
		}`,
	)

	req := httptest.NewRequest(http.MethodPost, "/put", body)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	PutHandler(rr, req)

	value, ok := StoreInstance.Get("name")

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	if !ok {
		t.Fatal("expected key to exist")
	}

	if value != "chinmay" {
		t.Errorf("expected chinmay, got %s", value)
	}
}

func TestBadJSON(t *testing.T) {
	StoreInstance := storage.NewStore()

	
}
