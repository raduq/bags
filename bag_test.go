package main

import (
	"testing"
)

func TestNewBag(t *testing.T) {
	b := NewBag(2)
	if 2 != b.Size {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	b := NewBag(2)
	b, _ = b.Add(NewItem("Ball"))
	if "Ball" != b.Items[0].Name {
		t.Fail()
	}
}
func TestAddWhenFull(t *testing.T) {
	b := NewBag(1)
	b, err := b.Add(NewItem("Ball"))
	b, err = b.Add(NewItem("Ball"))

	if err == nil || err.Error() != "Bag is full" {
		t.Fail()
	}
}
