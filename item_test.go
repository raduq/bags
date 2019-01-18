package main

import "testing"

func TestNewItem(t *testing.T) {
	i := NewItem("Ball")
	if "Ball" != i.Name {
		t.Fail()
	}
}
