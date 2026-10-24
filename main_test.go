package main
import "testing"

func testApplication(t *testing.T){
	message := "Hello From Go"
	if message != "Hello From Go" {
		t.Errorf("Expected greeting, got %q", message)
	}
}
