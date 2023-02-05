package calc

import "errors"

func Sum(no1, no2 int) int {
	return no1 + no2
}

func Divide(no1, no2 int) (error, int) {
	if no2 == 0 {
		return errors.New("Can't divide by zero"), 0
	} else {
		return nil, no1 / no2
	}

}
