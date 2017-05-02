package requirements

import "errors"

/**
Get a slice of all the integers that devide a given positive integer num
Go doesn't support an "inArray" functionality which is why this solution
is fairly basic, but it also makes it very readable
*/
func getPositiveDevisors(num int) ([]int, error) {
	var devisors []int

	if num <= 0 {
		return nil, errors.New("Provided value is not a positive integer")
	}

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			devisors = append(devisors, i)
		}
	}

	return devisors, nil
}
