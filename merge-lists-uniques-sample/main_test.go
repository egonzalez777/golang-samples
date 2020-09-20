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

	anotherUniqueList := uniqueNames([]string{}, []string{"Matt", "Henry"})

	if reflect.DeepEqual(expectedList, anotherUniqueList) == false {
		t.Errorf("Failed to retrieve correct list, expected %v, and got %v", expectedList, anotherUniqueList)
	}

	expectedDiverseList := []string{"Mark", "John", "Mikey", "Henry", "Jacob", "Mikey3"}
	lastUniqueList := uniqueNames([]string{"Mark", "John", "Mikey"}, []string{"Henry", "Jacob", "Mark", "Mikey", "Mikey3"})

	if reflect.DeepEqual(expectedDiverseList, lastUniqueList) == false {
		t.Errorf("Failed to retrieve correct list, expected %v, and got %v", expectedDiverseList, lastUniqueList)
	}
}
