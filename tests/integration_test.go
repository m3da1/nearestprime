package tests

import (
	"fmt"
	"testing"
)

func TestHelloAnother(t *testing.T) {
	expected := "3.00"
	actual := fmt.Sprintf("%.2f", 2.0+1.00)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
