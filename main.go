package main

import(
	"fmt"
	"observedb/internal/storage"
)

func main(){
	store := storage.NewStore()

	store.Put("name", "Chinmay")
	store.Put("city", "Pune")

	value, ok := store.Get("name")
	fmt.Println("name: ", value, ok)

	deleted := store.Delete("name")
	fmt.Println("deleted name: ", deleted)

	value, ok = store.Get("name")
	fmt.Println("name after delete: ", value, ok)
}