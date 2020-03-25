package utils

import (
	"fmt"
	"testing"
)

func TestAdjustPrecision(t *testing.T) {
	num := float64(4)
	guess := float64(1)

	res := adjustPrecision(num, guess)
	returnType := fmt.Sprintf("%T", res)

	if returnType != "float64" {
		t.Errorf("Expected return type to be float64, got %v", returnType)
	}
}

func TestRoundFloat(t *testing.T) {

}
