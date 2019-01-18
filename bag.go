package main

import "errors"

//Item is a consumable
type Item struct {
	Name string
}

//Bag is a holder of multiple Item
type Bag struct {
	Items []Item
	Size  int
}

//New creates a new Bag
func New(size int) Bag {
	return Bag{[]Item{}, size}
}

//Add a new item to the bag, if possible
func (b Bag) Add(i Item) (Bag, error) {
	if len(b.Items) >= b.Size {
		return b, errors.New("Bag is full")
	}
	b.Items = append(b.Items, i)
	return b, nil
}
