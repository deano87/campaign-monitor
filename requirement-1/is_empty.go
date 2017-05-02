package requirements

/**
isEmpty receives a string s and returns true if it is empty
As there are no nulls in go, only the length of the string was checked
*/
func isEmpty(s string) bool {
	return len(s) == 0
}
