package requirements

import "testing"

type TestStruct struct {
	s   string
	res bool
}

var toTest = []TestStruct{
	{
		s:   "",
		res: true,
	},
	{
		s:   "a",
		res: false,
	},
	{
		s:   "\"\"",
		res: false,
	},
	{
		s:   "null",
		res: false,
	},
}

func TestIsEmpty(t *testing.T) {
	for i := range toTest {
		if isEmpty(toTest[i].s) != toTest[i].res {
			t.Errorf("Expected %t, %t was given for string \"%s\"", toTest[i].res, isEmpty(toTest[i].s), toTest[i].s)
		}
	}
}
