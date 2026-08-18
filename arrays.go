package main

import "fmt"

func main() {
	// [] - means can have dynamic length
	x := []int{1, 2: 11, 3, 10: 12}
	x = append(x, 69)

	// [...] - means it has fixed length, array itself is immutable
	y := [...]int{1, 2, 3}
	// y = append(y, 70) displays error

	fmt.Println(x)
	fmt.Println(y)

	// Matrix
	multiple := [10][10]int{}

	for i := 1; i <= len(multiple); i++ {
		for j := 1; j <= len(multiple[i]); j++ {
			ctr := i * j
			fmt.Printf("%4v ", ctr)
		}
		fmt.Println()
	}

	// capacity & make - capacity works something like an autoscaler, if len(slice) == 9 - go automatically allocates
	// new +1 capacity in that slice
	cap_sample := make([]int, 9, 10)

	cap_sample = append(cap_sample, 6)
	cap_sample = append(cap_sample, 7) // reached capacity here
	cap_sample = append(cap_sample, 8) // capacity is now 11
	fmt.Println(cap_sample)

	// slicing a slice
	slice_dice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice_dice[0:])
	fmt.Println(slice_dice[:8])
	fmt.Println(slice_dice[:2:3]) // outputs [1,2]; third parameter is the capacity. as general rule 0 <= len(slice) <= cap(slice)
}
