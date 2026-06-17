package storage

import "testing"

func TestPutAndGet(t *testing.T){
	store := NewStore()

	store.Put("name", "Chinmay")

	value, ok := store.Get("name")

	if !ok{
		t.Error("Expected key to exist.")
	}

	if value != "Chinmay"{
		t.Errorf("Expected Chinmay, got %s", value)
	}
}

func TestGetMissingKey(t *testing.T){
	store := NewStore()

	_, ok := store.Get("missing")

	if ok{
		t.Error("Expected key to be missing")
	}
}

func TestDelete(t *testing.T){
	store := NewStore()

	store.Put("name", "Chinmay")

	deleted := store.Delete("name")

	if !deleted{
		t.Error("Expected delete to succeed")
	}

	_, ok := store.Get("name")

	if ok{
		t.Error("Expected key to be deleted")
	}
}

func TestDeleteMissingKey(t *testing.T){
	store := NewStore()

	deleted := store.Delete("missing")

	if deleted{
		t.Error("Expected delete to fail")
	}
}

