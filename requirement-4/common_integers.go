package requirements

/**
The function receives arguments of type int and returns
a slice of the most common ints found
The slice returned is not sorted - in order to keep the solution at O(n)
Logic:
Using a Hash Map, we count the number of appearances a number has,
and eventually loop trough the map, recording only the numbers that match the
max appearances
*/
func getCommonIntegers(args ...int) []int {
	var countMap map[int]int
	var finalSlice []int
	var max int

	// initializing map
	countMap = make(map[int]int)

	for _, el := range args {
		countMap[el]++
		if countMap[el] > max {
			max = countMap[el]
		}
	}

	for i, el := range countMap {
		if el == max {
			finalSlice = append(finalSlice, i)
		}
	}

	return finalSlice
}
