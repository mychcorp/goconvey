package convey

import "fmt"

const success = ""
const needOneValue = "This expectation requires exactly one comparison value (none provided)."
const onlyAcceptsOneValue = "This expectation only accepts 1 value to be compared (and %v were provided)."
const noValuesAccepted = "This expectation does not allow for user-supplied comparison values."
const shouldHaveBeenEqual = "'%v' should equal '%v' (but it doesn't)!"
const shouldNotHaveBeenEqual = "'%v' should NOT equal '%v' (but it does)!"
const shouldHaveResembled = "'%v' should resemble '%v' (but it doesn't)!"
const shouldNotHaveResembled = "'%v' should NOT resemble '%v' (but it does)!"
const shouldHaveBeenNil = "'%v' should have been nil (but it wasn't)!"
const shouldNotHaveBeenNil = "'%v' should NOT have been nil (but it was)!"
const shouldHaveBeenTrue = "Should have been 'true', not '%v'!"
const shouldHaveBeenFalse = "Should have been 'false', not '%v'!"

func onlyOne(expected []interface{}) string {
	switch {
	case len(expected) == 0:
		return needOneValue
	case len(expected) > 1:
		return fmt.Sprintf(onlyAcceptsOneValue, len(expected))
	}
	return success
}

func none(expected []interface{}) string {
	if len(expected) > 0 {
		return noValuesAccepted
	}
	return success
}
