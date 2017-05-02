package requirements

import (
	"sort"
	"testing"
)

type TestStruct struct {
	arr []int
	res []int
}

var toTest = []TestStruct{
	{
		arr: []int{5, 4, 3, 2, 4, 5, 1, 6, 1, 2, 5, 4},
		res: []int{5, 4},
	},
	{
		arr: []int{1, 2, 3, 4, 5, 1, 6, 7},
		res: []int{1},
	},
	{
		arr: []int{1, 2, 3, 4, 5, 6, 7},
		res: []int{1, 2, 3, 4, 5, 6, 7},
	},
}

func TestGetPositiveDevisors(t *testing.T) {
	for _, el := range toTest {
		common := getCommonIntegers(el.arr...)

		// sanity test
		if len(common) != len(el.res) {
			t.Errorf("Wrong number of common numbers returned")
			t.Errorf("Expected %v, Received %v",
				el.res,
				common,
			)
			continue
		}

		resArr := el.res
		// Sort both expected result and returned result to make sure the order
		// is not causing a problem in the test
		sort.Ints(common)
		sort.Ints(resArr)

		// make sure that the actual values returned are corret
		for j, val := range common {
			if val != resArr[j] {
				t.Errorf("Wrong number common number returned.")
				t.Errorf("Expected %d, received %d\n%v != %v",
					resArr[j],
					common[j],
					common,
					el.res,
				)
				break
			}
		}
	}
}
