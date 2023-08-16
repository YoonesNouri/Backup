package main

import "fmt"

// Recursive function to generate all subsets
func generateSubsets(set []string, subset []string, index int) {
	// Print the current subset
	fmt.Println(subset)

	// Iterate through the remaining elements
	for i := index; i < len(set); i++ {
		// Include the current element in the subset
		subset = append(subset, set[i])

		// Recursively generate subsets with the remaining elements
		generateSubsets(set, subset, i+1)

		// Exclude the current element from the subset
		subset = subset[:len(subset)-1]
	}
}

func main() {
	var n int

	// Prompt the user to enter the number of members in the set
	fmt.Print("Enter the number of members in the set: ")
	fmt.Scanln(&n)

	// Create a slice to store the members of the set
	set := make([]string, n)

	// Prompt the user to enter the members of the set
	fmt.Println("Enter the members of the set:")
	for i := 0; i < n; i++ {
		fmt.Printf("Member %d: ", i+1)
		fmt.Scanln(&set[i])
	}

	// Call the function to generate subsets
	generateSubsets(set, []string{}, 0)
}
