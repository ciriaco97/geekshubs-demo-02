package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	expectedResult := 4

	result := Sum(2, 2)

	if expectedResult != result {
		t.Errorf("The sum of 2 + 2 should be %d, got %d instead", expectedResult, result)
	}

}
