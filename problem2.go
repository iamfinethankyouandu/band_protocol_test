package main

func SupermanChickenRescue(n, k int, chickens []int) int {
	res := 0

	// The actual position of the chickens is needed because in the process of finding the amount of chickens that can protect, the position needs to be changed continuously.
	chickensClone := chickens

	for i := 0; i < n; i++ {
		help := 0
		// Create roof
		position := k + chickens[i]

		// Place the roof and find out how many chickens can help
		for j := 0; j < len(chickensClone); j++ {
			if position > chickensClone[j] {
				help++
			} else {
				break
			}
		}

		// Amount that can help the most
		if res < help {
			res = help
		}

		// Change frist positon for next roof
		chickensClone = chickensClone[1:]
	}

	return res
}
