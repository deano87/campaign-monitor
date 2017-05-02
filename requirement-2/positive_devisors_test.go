package requirements

import "testing"

type TestStruct struct {
	num int
	res []int
}

var toTest = []TestStruct{
	{
		num: 60,
		res: []int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60},
	},
	{
		num: 42,
		res: []int{1, 2, 3, 6, 7, 14, 21, 42},
	},
	{
		num: -1,
		res: nil,
	},
}

func TestGetPositiveDevisors(t *testing.T) {
	for i := range toTest {
		divisors, err := getPositiveDevisors(toTest[i].num)

		if err != nil {
			if toTest[i].res != nil {
				t.Errorf("Problem with the positive integer")
			}
			return
		}

		if len(divisors) != len(toTest[i].res) {
			if toTest[i].res != nil {
				t.Errorf("Wrong number of divisors")
			}
			return
		}

		// assuming that getPositiveDevisors returns a sorted slice
		// and that our result slice is sorted as well
		for j := range divisors {
			if divisors[j] != toTest[i].res[j] {
				t.Errorf("Wrong devisor, expected %d, got %d", toTest[i].res[j], divisors[j])
			}
		}
	}
}
