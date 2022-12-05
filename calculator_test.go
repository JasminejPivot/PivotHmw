package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if addition(5, 5) != 10 {
		t.Error("Expected 5 + 5 to equal 10")
	}

}
func TestSubtract(t *testing.T) {
	if subtract(5, 3) != 2 {
		t.Error("Expected 5 - 3 to equal 2")
	}
}
func TestMultiply(t *testing.T) {
	if multiply(5, 8) != 40 {
		t.Error("Expected 5 * 8 to equal 40")
	}
}
func TestDivide(t *testing.T) {
	if divide(20, 5) != 4 {
		t.Error("Expected 20 / 5 to equal 4")
		return
	}
}
