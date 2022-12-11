package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getFilterAttribute() int {
	var choice string
	errors := 0
	for errors < 5 {
		fmt.Printf("Enter the attribute you wish to filter by: \n")
		fmt.Printf("\n1. First Name\n2. Last Name\n3. Birth Date\n")
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == "1" {
			return 0
		} else if choice == "2" {
			return 1
		} else if choice == "3" {
			return 2
		} else {
			errors++
			fmt.Printf("Invalid Input. Try again")
		}
	}
	fmt.Printf("You made too many mistakes. I'm choosing first name for you\n")
	return 0
}

func getTarget(attributeType int) string {

	if attributeType == 0 {
		return getFirstName()
	} else if attributeType == 1 {
		return getLastName()
	} else {
		return getDOB()
	}

}

func getDOB() string {
	errors := 0
	for errors < 5 {
		var name string
		fmt.Printf("Enter the date of birth you wish to search for in format YYYYMMDD\n")
		fmt.Scanln(&name)
		if _, err := strconv.Atoi(name); err == nil {
			return name
		} else {
			fmt.Printf("Error. Please enter a integer in the given format for DOB.\n")
			errors++
		}
	}
	fmt.Printf("You made too many mistakes. I'm gonna say you chose 20000320 because that's my DOB")
	return "20000320"
}

func getLastName() string {
	errors := 0
	for errors < 5 {
		var name string
		fmt.Printf("Enter the last name you wish to search for\n")
		fmt.Scanln(&name)
		if _, err := strconv.Atoi(name); err == nil {
			fmt.Printf("Error. Please enter a name, not a number.\n")
			errors++
		} else {
			return name
		}
	}
	fmt.Printf("You made too many mistakes. I'm gonna say you chose Svetec because that's my last name")
	return "Svetec"
}

func getFirstName() string {
	errors := 0
	for errors < 5 {
		var name string
		fmt.Printf("Enter the first name you wish to search for\n")
		fmt.Scanln(&name)
		if _, err := strconv.Atoi(name); err == nil {
			fmt.Printf("Error. Please enter a name, not a number.\n")
			errors++
		} else {
			return name
		}
	}
	fmt.Printf("You made too many mistakes. I'm gonna say you chose Jack because that's my name")
	return "Jack"
}

func findTargetIndex(csv [][]string, attributeType int, target string) int {
	var middle int
	left := 0
	right := len(csv) - 1
	for left <= right {
		middle = (left + right) / 2
		if strings.ToLower(csv[middle][attributeType]) == target {
			return middle
		} else if strings.ToLower(csv[middle][attributeType]) > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func returnMatchingAttributes(csv [][]string, attributeType int, targetIndex int, target string) [][]string {
	// find first instance index by going down
	//then loop up list till no longer found, and return list
	var answer [][]string
	for targetIndex >= 0 && strings.ToLower(csv[targetIndex][attributeType]) == target {
		targetIndex--
	}

	targetIndex++

	for targetIndex < len(csv) && strings.ToLower(csv[targetIndex][attributeType]) == target {
		answer = append(answer, csv[targetIndex])
		targetIndex++
	}

	return answer
}

func main() {
	// initialize CSV
	csv := [][]string{
		{"Jack", "Svetec", "20000320"},
		{"Sarah", "Beatrice", "19980305"},
		{"Jack", "Reacher", "19990405"},
		{"Jack", "Machino", "19420418"},
		{"John", "Joseph", "19980305"},
		{"Sarah", "Devlin", "19880521"},
		{"Katy", "Machino", "20031103"}}
	// get User choice
	attributeType := getFilterAttribute()
	// sort CSV based on choice
	sort.Slice(csv, func(a int, b int) bool {
		return csv[a][attributeType] < csv[b][attributeType]
	})

	// get value they are searching for
	target := strings.ToLower(getTarget(attributeType))

	// search for match in CSV
	targetIndex := findTargetIndex(csv, attributeType, target)

	// return list of rows containing appropriate value
	if targetIndex == -1 {
		fmt.Printf("No rows exist with that attribute\n")
	} else {
		fmt.Printf("Here are the rows that contain this attribute: \n")
		answer := returnMatchingAttributes(csv, attributeType, targetIndex, target)
		fmt.Println(answer)
	}
	// fmt.Println(csv)

}
