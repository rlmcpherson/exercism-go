package leap

const TestVersion = 1

func IsLeapYear(year int) bool {

	if year%4 != 0 {
		return false

	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true

}
