package main

import "errors"

//Bag is a holder of multiple Item
type Bag struct {
	Items []Item
	Size  int
}

//NewBag creates a new Bag
func NewBag(size int) Bag {
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
