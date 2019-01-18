package main

import "github.com/google/uuid"

//Item is a consumable
type Item struct {
	ID   uuid.UUID
	Name string
}

//NewItem creates a new Item with random ID
func NewItem(name string) Item {
	return Item{uuid.New(), name}
}
