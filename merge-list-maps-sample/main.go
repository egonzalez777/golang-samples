package main

import "fmt"

func createUniqueSet(a, b []string) []string {
	combinedList := append(a, b...)
	mapped := map[string]string{}
	keysList := []string{}
	for i := 0; i < len(combinedList); i++ {
		fmt.Println(combinedList[i], "mmoooo: ", mapped[combinedList[i]])
		if _, found := mapped[combinedList[i]]; found {
			continue
		} else {
			mapped[combinedList[i]] = combinedList[i]
		}
	}
	for _, k := range mapped {
		keysList = append(keysList, k)
	}
	return keysList
}

func main() {
	uniques := createUniqueSet([]string{"Michael", "Graham", "Junior"}, []string{"Roger", "Spencer", "Graham", "Michael"})
	fmt.Println(uniques)
}
