package main
import "testing"

func TestApplication(t *testing.T){
	result := add(10,20)
	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}
