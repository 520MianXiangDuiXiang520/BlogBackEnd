package utils_test

import (
	"JuneBlog/internal/utils"
	"reflect"
	"testing"
)

func TestConversionSlice(t *testing.T) {
	testCase1 := []int{1, 2, 3, 4, 5}
	expectedResult1 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	result1 := utils.SliceConversion[int, float64](testCase1)
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("ConversionSlice failed for int to float64. Expected %v, got %v", expectedResult1, result1)
	}

	testCase2 := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	expectedResult2 := []int{1, 2, 3, 4, 5}

	result2 := utils.SliceConversion[float64, int](testCase2)

	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("ConversionSlice failed for float64 to int. Expected %v, got %v", expectedResult2, result2)
	}
}
