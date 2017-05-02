package requirements

import "testing"

type TestStruct struct {
	lengths []int
	res     float64
}

var toTest = []TestStruct{
	{
		lengths: []int{3, 4, 5},
		res:     6.0,
	},
}

var panicTest1 = []int{3, 4, 7}
var panicTest2 = []int{-3, 4, 5}

func TestGetPositiveDevisors(t *testing.T) {
	var area float64
	for i := range toTest {
		area = getTriangleArea(toTest[i].lengths[0],
			toTest[i].lengths[1],
			toTest[i].lengths[2],
		)
		if area != toTest[i].res {
			t.Errorf("Calculated area %f didn't match expected area %f", area, toTest[i].res)
		}
	}
}

func TestPanic1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	getTriangleArea(panicTest1[0],
		panicTest1[1],
		panicTest1[2],
	)
}

func TestPanic2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	getTriangleArea(panicTest2[0],
		panicTest2[1],
		panicTest2[2],
	)
}
