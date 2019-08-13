package main

import "testing"


func TestInitialize(t *testing.T) {
	got := Initialize()
	if got != 3 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}


}
