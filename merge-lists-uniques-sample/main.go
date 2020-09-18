package main

import "fmt"

func uniqueNames(a, b []string) []string {
	var result []string
	updatedList := append(a, b...)

	for i := 0; i < len(updatedList); i++ {
		result = appendIfNotFound(result, updatedList[i])
	}
	return result
}

func appendIfNotFound(result []string, name string) []string {
	for _, value := range result {
		if value == name {
			return result
		}
	}
	return append(result, name)
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
