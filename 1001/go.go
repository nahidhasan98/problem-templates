package main

import "fmt"

/**
 * Returns the number of problems in the first computer and second computer
 *
 * @param n denotes the total number of problems in both computer
**/
func process(n int) (int, int) {
	// Implement this method

}

/**
 * Takes care of the problem input and output.
**/
func main() {
	var test int
	fmt.Scan(&test)

	for t := 1; t <= test; t++ {
		var n int
		fmt.Scan(&n)

		a, b := process(n)
		fmt.Printf("%d %d\n", a, b)
	}
}
