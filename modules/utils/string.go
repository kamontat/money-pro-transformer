package utils

import "strconv"

// FloatToString will convert float64 to string
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'G', -1, 64)
}
