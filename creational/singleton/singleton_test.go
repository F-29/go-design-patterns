package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	count1 := GetInstance()
	if count1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance() got: nil")
	}
	expectedCount := count1

	currentCount := count1.AddOne()
	if currentCount != 1 {
		t.Errorf(`after calling AddOne() for the first time, Expected currentCount to be: 1 got: %+v instead.`, currentCount)
	}
	count2 := GetInstance()
	if count2 != expectedCount {
		t.Errorf(`Expected %+v same instance in count2 but got: %+v instead(count2).`, reflect.TypeOf(expectedCount), reflect.TypeOf(count2))
	}

	currentCount = count1.AddOne()
	if currentCount != 2 {
		t.Errorf(`after calling AddOne() for the second time, Expected currentCount to be: 2 got: %+v instead.`, currentCount)
	}
}
