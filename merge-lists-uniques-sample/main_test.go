package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	expectedList := []string{"Matt", "Henry"}
	uniqueList := uniqueNames([]string{"Matt", "Henry"}, []string{})

	if reflect.DeepEqual(expectedList, uniqueList) == false {
		t.Errorf("Failed to retrieve correct list, expected %v, and got %v", expectedList, uniqueList)
	}

	anotherUniqueLIst := uniqueNames([]string{}, []string{"Matt", "Henry"})

	if reflect.DeepEqual(expectedList, anotherUniqueLIst) == false {
		t.Errorf("Failed to retrieve correct list, expected %v, and got %v", expectedList, anotherUniqueLIst)
	}
}
