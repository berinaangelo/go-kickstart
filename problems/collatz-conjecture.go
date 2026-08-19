// For all 2.36×10^21 > n > 0:
// write a method that proves the collatz conjecture given n as an input
// while n do:
//	if n is even; n = n / 2
//  else; n = 3n + 1
// Output: all the possible results of that number, until it n == 1 (this might cause a repeated infinite loop of 1, 2, 4, 2, 1 ...)

package main

import "fmt"

// Loop approach
func threeN(n int) {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}

		fmt.Println(n)
	}
}

// recursive approach
func threeNRecursive(n int) {
	// TODO
}
